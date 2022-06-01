package partyhandler

import (
	"time"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	pg "github.com/clubo-app/protobuf/party"
	"github.com/clubo-app/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

func (h partyGatewayHandler) CreateParty(c *fiber.Ctx) error {
	req := new(pg.CreatePartyRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	user := middleware.ParseUser(c)
	req.RequesterId = user.Sub

	p, err := h.pc.CreateParty(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, _ := h.prof.GetProfile(c.Context(), &profile.GetProfileRequest{Id: p.UserId})

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
		StartDate:     p.StartDate.AsTime().UTC().Format(time.RFC3339),
		CreatedAt:     p.CreatedAt.AsTime().UTC().Format(time.RFC3339),
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
