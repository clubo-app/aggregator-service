package partyhandler

import (
	"strconv"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/protobuf/party"
	"github.com/clubo-app/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

func (h partyGatewayHandler) GetPartyByUser(c *fiber.Ctx) error {
	uId := c.Params("id")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
	if limit > 20 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 20")
	}

	offsetStr := c.Query("offset")
	offset, _ := strconv.ParseInt(offsetStr, 10, 32)

	parties, err := h.pc.GetByUser(c.Context(), &party.GetByUserRequest{UserId: uId, Offset: int32(offset), Limit: int32(limit)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profile, _ := h.prf.GetProfile(c.Context(), &profile.GetProfileRequest{Id: uId})

	aggP := make([]datastruct.AggregatedParty, len(parties.Parties))
	for i, p := range parties.Parties {
		aggP[i] = datastruct.PartyToAgg(p).AddCreator(datastruct.ProfileToAgg(profile))
	}

	res := datastruct.PagedAggregatedParty{
		Parties: aggP,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
