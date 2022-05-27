package partyhandler

import (
	pg "github.com/clubo-app/protobuf/party"
	sg "github.com/clubo-app/protobuf/story"
	ug "github.com/clubo-app/protobuf/user"
	"github.com/gofiber/fiber/v2"
)

type partyGatewayHandler struct {
	pc pg.PartyServiceClient
	uc ug.UserServiceClient
	sc sg.StoryServiceClient
}

type PartyGatewayHandler interface {
	CreateParty(c *fiber.Ctx) error
	UpdateParty(c *fiber.Ctx) error
	DeleteParty(c *fiber.Ctx) error
	GetParty(c *fiber.Ctx) error
	GetPartyByUser(c *fiber.Ctx) error
}

func NewPartyGatewayHandler(pc pg.PartyServiceClient, uc ug.UserServiceClient, sc sg.StoryServiceClient) PartyGatewayHandler {
	return &partyGatewayHandler{
		pc: pc,
		uc: uc,
		sc: sc,
	}
}
