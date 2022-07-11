package authhandler

import (
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/protobuf/auth"
	"github.com/gofiber/fiber/v2"
)

func (h authGatewayHandler) RefreshAccessToken(c *fiber.Ctx) error {
	rt := c.Params("rt")

	t, err := h.ac.RefreshAccessToken(c.Context(), &auth.RefreshAccessTokenRequest{RefreshToken: rt})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(t)
}
