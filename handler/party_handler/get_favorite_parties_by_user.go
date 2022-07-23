package partyhandler

import (
	"strconv"
	"time"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/protobuf/party"
	rg "github.com/clubo-app/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h *partyGatewayHandler) GetFavoritePartiesByUser(c *fiber.Ctx) error {
	uId := c.Params("id")
	nextPage := c.Query("nextPage")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)
	if limit > 20 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 20")
	}

	favParties, err := h.rc.GetFavoritePartiesByUser(c.Context(), &rg.GetFavoritePartiesByUserRequest{UserId: uId, NextPage: nextPage, Limit: limit})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	partyIds := make([]string, len(favParties.FavoriteParties))
	for i, fp := range favParties.FavoriteParties {
		partyIds[i] = fp.PartyId
	}

	parties, _ := h.pc.GetManyPartiesMap(c.Context(), &party.GetManyPartiesRequest{Ids: partyIds})
	if parties == nil {
		res := datastruct.PagedAggregatedFavoriteParty{
			FavoriteParties: []datastruct.AggregatedFavoriteParty{},
			NextPage:        favParties.NextPage,
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}

	aggFP := make([]datastruct.AggregatedFavoriteParty, len(favParties.FavoriteParties))
	for i, fp := range favParties.FavoriteParties {
		aggFP[i] = datastruct.AggregatedFavoriteParty{
			UserId:      fp.UserId,
			Party:       datastruct.PartyToAgg(parties.Parties[fp.PartyId]),
			FavoritedAt: fp.FavoritedAt.AsTime().UTC().Format(time.RFC3339),
		}
	}

	res := datastruct.PagedAggregatedFavoriteParty{
		FavoriteParties: aggFP,
		NextPage:        favParties.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
