package profilehandler

import (
	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	"github.com/clubo-app/protobuf/auth"
	pg "github.com/clubo-app/protobuf/profile"
	rg "github.com/clubo-app/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h profileGatewayHandler) GetMe(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	p, err := h.pc.GetProfile(c.Context(), &pg.GetProfileRequest{Id: user.Sub})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	if p == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Profile not found")
	}

	a, err := h.ac.GetAccount(c.Context(), &auth.GetAccountRequest{Id: user.Sub})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	if a == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Account not found")
	}

	friendCountRes, _ := h.rc.GetFriendCount(c.Context(), &rg.GetFriendCountRequest{UserId: p.Id})

	res := datastruct.AggregatedAccount{
		Id: p.Id,
		Profile: datastruct.AggregatedProfile{
			Id:          p.Id,
			Username:    p.Username,
			Firstname:   p.Firstname,
			Lastname:    p.Lastname,
			Avatar:      p.Avatar,
			FriendCount: friendCountRes.FriendCount,
		},
		Email: a.Email,
		Type:  a.Type.String(),
		Role:  a.Role.String(),
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
