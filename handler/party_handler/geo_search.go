package partyhandler

import (
	"strconv"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	pg "github.com/clubo-app/protobuf/party"
	"github.com/gofiber/fiber/v2"
)

func (h partyGatewayHandler) GeoSearch(c *fiber.Ctx) error {
	limitStr := c.Query("limit")
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
	offsetStr := c.Query("offset")
	offset, _ := strconv.ParseInt(offsetStr, 10, 32)

	latStr := c.Params("lat")
	lat, _ := strconv.ParseFloat(latStr, 32)
	longStr := c.Params("long")
	long, _ := strconv.ParseFloat(longStr, 32)

	radiusStr := c.Query("radius")
	radius, _ := strconv.ParseInt(radiusStr, 10, 32)

	parties, err := h.pc.GeoSearch(c.Context(), &pg.GeoSearchRequest{
		Lat:    float32(lat),
		Long:   float32(long),
		Limit:  int32(limit),
		Offset: int32(offset),
		Radius: int32(radius),
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	aggP := make([]datastruct.AggregatedParty, len(parties.Parties))
	for i, p := range parties.Parties {
		aggP[i] = datastruct.PartyToAgg(p)
	}

	res := datastruct.PagedAggregatedParty{
		Parties: aggP,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
