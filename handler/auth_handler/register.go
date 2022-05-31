package authhandler

import (
	"github.com/clubo-app/packages/utils"
	ag "github.com/clubo-app/protobuf/auth"
	"github.com/clubo-app/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}

func (h authGatewayHandler) Register(c *fiber.Ctx) error {
	req := new(RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	nameTaken, err := h.pc.UsernameTaken(c.Context(), &profile.UsernameTakenRequest{
		Username: req.Username,
	})
	if err != nil || nameTaken.Taken {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Username already taken")
	}

	u, err := h.ac.RegisterUser(c.Context(), &ag.RegisterUserRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	p, err := h.pc.CreateProfile(c.Context(), &profile.CreateProfileRequest{
		Id:        u.Account.Id,
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Avatar:    req.Avatar,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusCreated).JSON(p)
}
