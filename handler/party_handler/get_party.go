package partyhandler

import (
	"log"

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
	log.Printf("Party: %v", p)

	profileRes, _ := h.prof.GetProfile(c.Context(), &profile.GetProfileRequest{Id: p.UserId})
	log.Printf("Profile: %v", profileRes)

	storyRes, _ := h.sc.GetByParty(c.Context(), &sg.GetByPartyRequest{PartyId: p.Id})

	res := datastruct.AggregatedParty{
		Id:            p.Id,
		Creator:       profileRes,
		Title:         p.Title,
		IsPublic:      p.IsPublic,
		Lat:           p.Lat,
		Long:          p.Long,
		StreetAddress: p.StreetAddress,
		PostalCode:    p.PostalCode,
		State:         p.State,
		Country:       p.Country,
		StartDate:     p.StartDate,
		CreatedAt:     p.CreatedAt,
	}

	if storyRes != nil {
		res.Stories = storyRes.Stories
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
