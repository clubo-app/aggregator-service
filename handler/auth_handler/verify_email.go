package authhandler

import (
	"github.com/clubo-app/packages/utils"
	ug "github.com/clubo-app/protobuf/user"
	"github.com/gofiber/fiber/v2"
)

func (h authGatewayHandler) VerifyEmail(c *fiber.Ctx) error {
	req := new(ug.VerifyEmailRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	res, err := h.uc.VerifyEmail(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
