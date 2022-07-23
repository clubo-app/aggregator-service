package profilehandler

import (
	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	ag "github.com/clubo-app/protobuf/auth"
	pg "github.com/clubo-app/protobuf/profile"
	rg "github.com/clubo-app/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

type UpdateRequest struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}

func (h profileGatewayHandler) UpdateUser(c *fiber.Ctx) error {
	req := new(UpdateRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	user := middleware.ParseUser(c)

	a := new(ag.Account)
	p := new(pg.Profile)

	if req.Email != "" || req.Password != "" {
		res, err := h.ac.UpdateAccount(c.Context(), &ag.UpdateAccountRequest{
			Id:       user.Sub,
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			return utils.ToHTTPError(err)
		}
		a = res
	}

	if req.Username != "" || req.Firstname != "" || req.Lastname != "" || req.Avatar != "" {
		res, err := h.pc.UpdateProfile(c.Context(), &pg.UpdateProfileRequest{
			Id:        user.Sub,
			Username:  req.Username,
			Firstname: req.Firstname,
			Lastname:  req.Lastname,
			Avatar:    req.Avatar,
		})
		if err != nil {
			return utils.ToHTTPError(err)
		}
		p = res
	}

	friendCountRes, _ := h.rc.GetFriendCount(c.Context(), &rg.GetFriendCountRequest{UserId: p.Id})

	res := datastruct.AggregatedAccount{
		Id:      p.Id,
		Profile: datastruct.ProfileToAgg(p).AddFCount(friendCountRes.FriendCount),
		Email:   a.Email,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
