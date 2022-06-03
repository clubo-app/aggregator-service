package storyhandler

import (
	"strconv"

	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h storyGatewayHandler) GetStoryByParty(c *fiber.Ctx) error {
	pId := c.Params("id")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)
	offsetStr := c.Query("offset")
	offset, _ := strconv.ParseInt(offsetStr, 10, 32)

	res, err := h.sc.GetByParty(c.Context(), &story.GetByPartyRequest{PartyId: pId, Offset: int32(offset), Limit: int32(limit)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
