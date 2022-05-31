package partyhandler

import (
	"github.com/clubo-app/protobuf/party"
	"github.com/clubo-app/protobuf/profile"
	"github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

type partyGatewayHandler struct {
	pc   party.PartyServiceClient
	prof profile.ProfileServiceClient
	sc   story.StoryServiceClient
}

type PartyGatewayHandler interface {
	CreateParty(c *fiber.Ctx) error
	UpdateParty(c *fiber.Ctx) error
	DeleteParty(c *fiber.Ctx) error
	GetParty(c *fiber.Ctx) error
	GetPartyByUser(c *fiber.Ctx) error
}

func NewPartyGatewayHandler(pc party.PartyServiceClient, prof profile.ProfileServiceClient, sc story.StoryServiceClient) PartyGatewayHandler {
	return &partyGatewayHandler{
		pc:   pc,
		prof: prof,
		sc:   sc,
	}
}
