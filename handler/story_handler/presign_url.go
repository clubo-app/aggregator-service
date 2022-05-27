package storyhandler

import (
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h storyGatewayHandler) PresignURL(c *fiber.Ctx) error {
	key := c.Params("key")

	res, err := h.sc.PresignURL(c.Context(), &story.PresignURLRequest{Key: key})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
