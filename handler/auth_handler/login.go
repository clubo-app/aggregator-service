package authhandler

import (
	"github.com/clubo-app/packages/utils"
	ag "github.com/clubo-app/protobuf/auth"
	"github.com/gofiber/fiber/v2"
)

func (h authGatewayHandler) Login(c *fiber.Ctx) error {
	req := new(ag.LoginUserRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	res, err := h.ac.LoginUser(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
