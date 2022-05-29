package profilehandler

import (
	"log"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	rg "github.com/clubo-app/protobuf/relation"
	ug "github.com/clubo-app/protobuf/user"
	"github.com/gofiber/fiber/v2"
)

func (h userGatewayHandler) GetProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	user := middleware.ParseUser(c)

	profile, err := h.uc.GetProfile(c.Context(), &ug.GetProfileRequest{Id: id})
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
		Id:          profile.Id,
		Username:    profile.Username,
		Firstname:   profile.Firstname,
		Lastname:    profile.Lastname,
		Avatar:      profile.Avatar,
		FriendCount: friendCountRes.FriendCount,
	}

	log.Println(relation)

	if relation != nil {
		st := datastruct.FriendshipStatus{}

		if relation.Accepted {
			st.IsFriend = true
			st.AcceptedAt = relation.AcceptedAt
		} else {
			st.OutgoingRequest = true
			st.RequestedAt = relation.RequestedAt
		}

		res.FriendshipStatus = st
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
