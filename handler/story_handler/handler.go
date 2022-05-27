package storyhandler

import (
	sc "github.com/clubo-app/protobuf/story"
	uc "github.com/clubo-app/protobuf/user"
	"github.com/gofiber/fiber/v2"
)

type storyGatewayHandler struct {
	sc sc.StoryServiceClient
	uc uc.UserServiceClient
}

type StoryGatewayHandler interface {
	CreateStory(c *fiber.Ctx) error
	GetStory(c *fiber.Ctx) error
	GetStoryByParty(c *fiber.Ctx) error
	GetStoryByUser(c *fiber.Ctx) error
	DeleteStory(c *fiber.Ctx) error
	PresignURL(c *fiber.Ctx) error
}

func NewStoryGatewayHandler(sc sc.StoryServiceClient, uc uc.UserServiceClient) StoryGatewayHandler {
	return &storyGatewayHandler{
		sc: sc,
		uc: uc,
	}
}
