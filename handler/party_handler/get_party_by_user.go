package partyhandler

import (
	"strconv"
	"time"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/protobuf/party"
	"github.com/clubo-app/protobuf/profile"
	sg "github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h partyGatewayHandler) GetPartyByUser(c *fiber.Ctx) error {
	uId := c.Params("id")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
	offsetStr := c.Query("offset")
	offset, _ := strconv.ParseInt(offsetStr, 10, 32)

	partyRes, err := h.pc.GetByUser(c.Context(), &party.GetByUserRequest{UserId: uId, Offset: int32(offset), Limit: int32(limit)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	// get the profile of the party creator
	profilesRes, _ := h.prf.GetProfile(c.Context(), &profile.GetProfileRequest{Id: uId})

	aggP := make([]datastruct.AggregatedParty, len(partyRes.Parties))
	for i, p := range partyRes.Parties {
		aggP[i] = datastruct.AggregatedParty{
			Id:            p.Id,
			Creator:       profilesRes,
			Title:         p.Title,
			IsPublic:      p.IsPublic,
			Lat:           p.Lat,
			Long:          p.Long,
			StreetAddress: p.StreetAddress,
			PostalCode:    p.PostalCode,
			State:         p.State,
			Country:       p.Country,
			// TODO: we might want to fetch some stories of the party but would have to do this for all party returned of this user
			Stories:   []*sg.Story{},
			StartDate: p.StartDate.AsTime().UTC().Format(time.RFC3339),
			CreatedAt: p.CreatedAt.AsTime().UTC().Format(time.RFC3339),
		}
	}

	res := datastruct.PagedAggregatedParty{
		Parties: aggP,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
