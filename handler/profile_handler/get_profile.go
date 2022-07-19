package profilehandler

import (
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

	p, err := h.pc.GetProfile(c.Context(), &pg.GetProfileRequest{
		Id: id,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	if p == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Profile not found")
	}

	res := datastruct.ProfileToAgg(p)

	relation := new(rg.FriendRelation)
	// if somebody wants the Profile of somebody else we also return the friendship status between them two
	if id != user.Sub {
		relation, _ = h.rc.GetFriendRelation(c.Context(), &rg.GetFriendRelationRequest{UserId: user.Sub, FriendId: id})
	}
	if relation != nil {
		fs := datastruct.ParseFriendShipStatus(user.Sub, relation)
		res.AddFs(fs)
	}

	friendCountRes, _ := h.rc.GetFriendCount(c.Context(), &rg.GetFriendCountRequest{UserId: p.Id})
	if friendCountRes != nil {
		res.AddFCount(friendCountRes.FriendCount)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
