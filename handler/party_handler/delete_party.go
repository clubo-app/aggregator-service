package partyhandler

import (
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	"github.com/clubo-app/protobuf/party"
	"github.com/gofiber/fiber/v2"
)

func (h partyGatewayHandler) DeleteParty(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	pId := c.Params("id")

	res, err := h.pc.DeleteParty(c.Context(), &party.DeletePartyRequest{RequesterId: user.Sub, PartyId: pId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
