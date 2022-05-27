package authhandler

import (
	"github.com/clubo-app/packages/utils"
	ug "github.com/clubo-app/protobuf/user"
	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

func (h authGatewayHandler) Register(c *fiber.Ctx) error {
	req := new(ug.RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	u, err := h.uc.Register(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(u)
}
