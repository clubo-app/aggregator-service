package partyhandler

import (
	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	"github.com/clubo-app/protobuf/party"
	"github.com/clubo-app/protobuf/profile"
	sg "github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h partyGatewayHandler) UpdateParty(c *fiber.Ctx) error {
	req := new(party.UpdatePartyRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	user := middleware.ParseUser(c)
	req.RequesterId = user.Sub
	req.PartyId = c.Params("id")

	p, err := h.pc.UpdateParty(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, err := h.prf.GetProfile(c.Context(), &profile.GetProfileRequest{Id: p.UserId})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	res := datastruct.PartyToAgg(p).AddCreator(datastruct.ProfileToAgg(profileRes))

	storyRes, err := h.sc.GetByParty(c.Context(), &sg.GetByPartyRequest{PartyId: p.Id})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	if storyRes != nil {
		res.AddStory(storyRes.Stories)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
