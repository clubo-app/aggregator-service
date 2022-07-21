package partyhandler

import (
	pg "github.com/clubo-app/protobuf/party"
	prfg "github.com/clubo-app/protobuf/profile"
	rg "github.com/clubo-app/protobuf/relation"
	sg "github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

type partyGatewayHandler struct {
	pc  pg.PartyServiceClient
	prf prfg.ProfileServiceClient
	sc  sg.StoryServiceClient
	rc  rg.RelationServiceClient
}

type PartyGatewayHandler interface {
	CreateParty(c *fiber.Ctx) error
	UpdateParty(c *fiber.Ctx) error
	DeleteParty(c *fiber.Ctx) error
	GetParty(c *fiber.Ctx) error
	GetPartyByUser(c *fiber.Ctx) error

	GetFavoritePartiesByUser(c *fiber.Ctx) error
	GetFavorisingUsersByParty(c *fiber.Ctx) error
	GeoSearch(c *fiber.Ctx) error
}

func NewPartyGatewayHandler(pc pg.PartyServiceClient, prf prfg.ProfileServiceClient, sc sg.StoryServiceClient, rc rg.RelationServiceClient) PartyGatewayHandler {
	return &partyGatewayHandler{
		pc:  pc,
		prf: prf,
		sc:  sc,
		rc:  rc,
	}
}
