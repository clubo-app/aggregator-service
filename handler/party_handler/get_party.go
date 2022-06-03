package partyhandler

import (
	"time"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/protobuf/party"
	"github.com/clubo-app/protobuf/profile"
	sg "github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h partyGatewayHandler) GetParty(c *fiber.Ctx) error {
	id := c.Params("id")

	p, err := h.pc.GetParty(c.Context(), &party.GetPartyRequest{PartyId: id})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, _ := h.prof.GetProfile(c.Context(), &profile.GetProfileRequest{Id: p.UserId})

	storyRes, _ := h.sc.GetByParty(c.Context(), &sg.GetByPartyRequest{PartyId: p.Id})

	res := datastruct.AggregatedParty{
		Id:            p.Id,
		Creator:       profileRes,
		Title:         p.Title,
		IsPublic:      p.IsPublic,
		Lat:           p.Lat,
		Long:          p.Long,
		StreetAddress: p.StreetAddress,
		Stories:       []*sg.PublicStory{},
		PostalCode:    p.PostalCode,
		State:         p.State,
		Country:       p.Country,
		StartDate:     p.StartDate.AsTime().UTC().Format(time.RFC3339),
		CreatedAt:     p.CreatedAt.AsTime().UTC().Format(time.RFC3339),
	}

	if storyRes != nil {
		res.Stories = storyRes.Stories
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
