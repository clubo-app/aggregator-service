package relationhandler

import (
	"strconv"
	"time"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/protobuf/profile"
	rg "github.com/clubo-app/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h relationGatewayHandler) GetFriends(c *fiber.Ctx) error {
	uId := c.Params("id")
	nextPage := c.Query("nextPage")

	acceptedStr := c.Query("accepted")
	accepted, acceptedErr := strconv.ParseBool(acceptedStr)

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)

	var fr *rg.PagedFriendRelations
	if !accepted && acceptedErr == nil {
		var err error
		fr, err = h.rc.GetIncomingFriendRequests(c.Context(), &rg.GetIncomingFriendRequestsRequest{UserId: uId, NextPage: nextPage, Limit: limit})
		if err != nil {
			return utils.ToHTTPError(err)
		}
	} else {
		var err error
		fr, err = h.rc.GetFriends(c.Context(), &rg.GetFriendsRequest{UserId: uId, NextPage: nextPage, Limit: limit})
		if err != nil {
			return utils.ToHTTPError(err)
		}
	}

	var ids []string
	for _, fp := range fr.Relations {
		ids = append(ids, fp.FriendId)
	}

	profiles, err := h.pc.GetManyProfilesMap(c.Context(), &profile.GetManyProfilesRequest{Ids: utils.UniqueStringSlice(ids)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	aggP := make([]datastruct.AggregatedProfile, len(profiles.Profiles))
	for i, f := range fr.Relations {
		p := profiles.Profiles[f.FriendId]

		fs := datastruct.FriendshipStatus{
			IsFriend:        f.Accepted,
			OutgoingRequest: !f.Accepted,
		}

		if !f.RequestedAt.AsTime().IsZero() {
			fs.RequestedAt = f.RequestedAt.AsTime().UTC().Format(time.RFC3339)
		}
		if !f.AcceptedAt.AsTime().IsZero() {
			fs.AcceptedAt = f.AcceptedAt.AsTime().UTC().Format(time.RFC3339)
		}

		aggP[i] = datastruct.AggregatedProfile{
			Id:               p.Id,
			Username:         p.Username,
			Firstname:        p.Lastname,
			Lastname:         p.Lastname,
			Avatar:           p.Avatar,
			FriendshipStatus: fs,
		}
	}

	res := datastruct.PagedAggregatedProfile{
		AggregatedProfile: aggP,
		NextPage:          fr.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
