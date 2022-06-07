package profilehandler

import (
	"time"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	pg "github.com/clubo-app/protobuf/profile"
	rg "github.com/clubo-app/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h profileGatewayHandler) GetProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	user := middleware.ParseUser(c)

	profile, err := h.pc.GetProfile(c.Context(), &pg.GetProfileRequest{
		Id: id,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	if profile == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Profile not found")
	}

	var relation *rg.FriendRelation
	// if somebody wants the Profile of somebody else we also return the friendship status between them two
	if id != user.Sub {
		relation, _ = h.rc.GetFriendRelation(c.Context(), &rg.GetFriendRelationRequest{UserId: user.Sub, FriendId: id})
	}

	friendCountRes, _ := h.rc.GetFriendCount(c.Context(), &rg.GetFriendCountRequest{UserId: profile.Id})

	res := datastruct.AggregatedProfile{
		Id:        profile.Id,
		Username:  profile.Username,
		Firstname: profile.Firstname,
		Lastname:  profile.Lastname,
		Avatar:    profile.Avatar,
	}

	if friendCountRes != nil {
		res.FriendCount = friendCountRes.FriendCount
	}

	if relation != nil {
		fs := datastruct.FriendshipStatus{}

		if relation.Accepted {
			fs.IsFriend = true
			fs.AcceptedAt = relation.AcceptedAt.AsTime().UTC().Format(time.RFC3339)
		} else {
			fs.OutgoingRequest = true
			fs.RequestedAt = relation.RequestedAt.AsTime().UTC().Format(time.RFC3339)
		}

		res.FriendshipStatus = fs
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
