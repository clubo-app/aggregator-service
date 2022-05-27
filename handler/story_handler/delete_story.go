package storyhandler

import (
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	"github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h storyGatewayHandler) DeleteStory(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	sId := c.Params("id")

	res, err := h.sc.DeleteStory(c.Context(), &story.DeleteStoryRequest{RequesterId: user.Sub, StoryId: sId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
