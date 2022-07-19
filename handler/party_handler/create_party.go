package partyhandler

import (
	"time"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	"github.com/clubo-app/protobuf/party"
	"github.com/clubo-app/protobuf/profile"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreatePartyReq struct {
	Title         string    `json:"title"`
	Lat           float32   `json:"lat"`
	Long          float32   `json:"long"`
	IsPublic      bool      `json:"is_public"`
	StreetAddress string    `json:"street_address"`
	PostalCode    string    `json:"postal_code"`
	State         string    `json:"state"`
	Country       string    `json:"country"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
}

func (h partyGatewayHandler) CreateParty(c *fiber.Ctx) error {
	req := new(CreatePartyReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	user := middleware.ParseUser(c)

	p, err := h.pc.CreateParty(c.Context(), &party.CreatePartyRequest{
		RequesterId:   user.Sub,
		Title:         req.Title,
		Lat:           req.Lat,
		Long:          req.Long,
		IsPublic:      req.IsPublic,
		StreetAddress: req.StreetAddress,
		PostalCode:    req.PostalCode,
		State:         req.State,
		Country:       req.Country,
		StartDate:     timestamppb.New(req.StartDate),
		EndDate:       timestamppb.New(req.EndDate),
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, _ := h.prf.GetProfile(c.Context(), &profile.GetProfileRequest{Id: p.UserId})

	res := datastruct.PartyToAgg(p).AddCreator(profileRes)
	return c.Status(fiber.StatusCreated).JSON(res)
}
