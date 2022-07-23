package relationhandler

import (
	"strconv"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	"github.com/clubo-app/protobuf/profile"
	rg "github.com/clubo-app/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h relationGatewayHandler) GetFriends(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	uId := c.Params("id")
	nextPage := c.Query("nextPage")

	acceptedStr := c.Query("accepted")
	accepted, acceptedErr := strconv.ParseBool(acceptedStr)

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)
	if limit > 40 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 40")
	}

	// find the friends or incoming friend requests of the wanted user
	fr := new(rg.PagedFriendRelations)
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

	ids := make([]string, len(fr.Relations))
	for i, fp := range fr.Relations {
		ids[i] = fp.FriendId
	}

	// we get the profile of all the user having a relation to the wanted user
	profiles, err := h.pc.GetManyProfilesMap(c.Context(), &profile.GetManyProfilesRequest{Ids: utils.UniqueStringSlice(ids)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	aggP := make([]datastruct.AggregatedProfile, len(profiles.Profiles))
	for i, f := range fr.Relations {
		p := profiles.Profiles[f.FriendId]

		// we see if the friends are also freinds of the requester, this way we can display if the user is already friends with the friends
		fs := datastruct.ParseFriendShipStatus(user.Sub, f)

		aggP[i] = datastruct.ProfileToAgg(p).AddFs(fs)
	}

	res := datastruct.PagedAggregatedProfile{
		AggregatedProfile: aggP,
		NextPage:          fr.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
