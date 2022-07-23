package partyhandler

import (
	"strconv"
	"time"

	"github.com/clubo-app/protobuf/profile"
	rg "github.com/clubo-app/protobuf/relation"
	"github.com/gofiber/fiber/v2"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
)

func (h partyGatewayHandler) GetFavorisingUsersByParty(c *fiber.Ctx) error {
	pId := c.Params("id")
	nextPage := c.Query("nextPage")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)
	if limit > 20 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 20")
	}

	fpRes, err := h.rc.GetFavorisingUsersByParty(c.Context(), &rg.GetFavorisingUsersByPartyRequest{PartyId: pId, NextPage: nextPage, Limit: limit})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	ids := make([]string, len(fpRes.FavoriteParties))
	for i, fp := range fpRes.FavoriteParties {
		ids[i] = fp.UserId
	}

	pRes, _ := h.prf.GetManyProfilesMap(c.Context(), &profile.GetManyProfilesRequest{Ids: ids})
	if pRes == nil {
		res := datastruct.PagedAggregatedFavorisingUsers{
			FavoriteParties: []datastruct.AggregatedFavorisingUsers{},
			NextPage:        fpRes.NextPage,
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}

	aggFP := make([]datastruct.AggregatedFavorisingUsers, len(fpRes.FavoriteParties))
	for i, fp := range fpRes.FavoriteParties {
		aggFP[i] = datastruct.AggregatedFavorisingUsers{
			User:        pRes.Profiles[fp.UserId],
			PartyId:     fp.PartyId,
			FavoritedAt: fp.FavoritedAt.AsTime().UTC().Format(time.RFC3339),
		}
	}

	res := datastruct.PagedAggregatedFavorisingUsers{
		FavoriteParties: aggFP,
		NextPage:        fpRes.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
