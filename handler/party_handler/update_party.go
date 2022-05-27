package partyhandler

import (
	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	"github.com/clubo-app/protobuf/party"
	sg "github.com/clubo-app/protobuf/story"
	ug "github.com/clubo-app/protobuf/user"
	"github.com/gofiber/fiber/v2"
)

func (h partyGatewayHandler) UpdateParty(c *fiber.Ctx) error {
	req := new(party.UpdatePartyRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	user := middleware.ParseUser(c)
	req.RequesterId = user.Sub

	p, err := h.pc.UpdateParty(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, err := h.uc.GetProfile(c.Context(), &ug.GetProfileRequest{Id: p.UserId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	storyRes, err := h.sc.GetByParty(c.Context(), &sg.GetByPartyRequest{PartyId: p.Id})
	if err != nil {
		return utils.ToHTTPError(err)
	}

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
		Stories:       storyRes.Stories,
		StartDate:     p.StartDate,
		CreatedAt:     p.CreatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
