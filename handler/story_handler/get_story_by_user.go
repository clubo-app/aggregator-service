package storyhandler

import (
	"strconv"

	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h storyGatewayHandler) GetStoryByUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)
	nextPage := c.Query("nextPage")

	res, err := h.sc.GetByUser(c.Context(), &story.GetByUserRequest{
		UserId:   userId,
		NextPage: nextPage,
		Limit:    uint32(limit),
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
