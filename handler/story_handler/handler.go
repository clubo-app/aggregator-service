package storyhandler

import (
	pg "github.com/clubo-app/protobuf/party"
	prfg "github.com/clubo-app/protobuf/profile"
	sg "github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

type storyGatewayHandler struct {
	sc  sg.StoryServiceClient
	pc  pg.PartyServiceClient
	prf prfg.ProfileServiceClient
}

type StoryGatewayHandler interface {
	CreateStory(c *fiber.Ctx) error
	GetStoryByParty(c *fiber.Ctx) error
	GetStoryByUser(c *fiber.Ctx) error
	DeleteStory(c *fiber.Ctx) error
	PresignURL(c *fiber.Ctx) error
}

func NewStoryGatewayHandler(sc sg.StoryServiceClient, prf prfg.ProfileServiceClient, pc pg.PartyServiceClient) StoryGatewayHandler {
	return &storyGatewayHandler{
		sc:  sc,
		prf: prf,
		pc:  pc,
	}
}
