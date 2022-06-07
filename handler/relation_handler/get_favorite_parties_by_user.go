package relationhandler

import (
	"strconv"
	"time"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/protobuf/party"
	rg "github.com/clubo-app/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h *relationGatewayHandler) GetFavoritePartiesByUser(c *fiber.Ctx) error {
	uId := c.Params("id")
	nextPage := c.Query("nextPage")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)

	fpRes, err := h.rc.GetFavoritePartiesByUser(c.Context(), &rg.GetFavoritePartiesByUserRequest{UserId: uId, NextPage: nextPage, Limit: limit})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	var partyIds []string
	for _, fp := range fpRes.FavoriteParties {
		partyIds = append(partyIds, fp.PartyId)
	}

	pRes, _ := h.pc.GetManyPartiesMap(c.Context(), &party.GetManyPartiesRequest{Ids: partyIds})

	aggFP := make([]datastruct.AggregatedFavoriteParty, len(fpRes.FavoriteParties))
	for i, fp := range fpRes.FavoriteParties {
		aggFP[i] = datastruct.AggregatedFavoriteParty{
			UserId:      fp.UserId,
			Party:       pRes.Parties[fp.PartyId],
			FavoritedAt: fp.FavoritedAt.AsTime().UTC().Format(time.RFC3339),
		}
	}

	res := datastruct.PagedAggregatedFavoriteParty{
		FavoriteParties: aggFP,
		NextPage:        fpRes.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
