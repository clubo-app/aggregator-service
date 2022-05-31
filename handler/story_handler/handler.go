package storyhandler

import (
	"github.com/clubo-app/protobuf/profile"
	sc "github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

type storyGatewayHandler struct {
	sc   sc.StoryServiceClient
	prof profile.ProfileServiceClient
}

type StoryGatewayHandler interface {
	CreateStory(c *fiber.Ctx) error
	GetStory(c *fiber.Ctx) error
	GetStoryByParty(c *fiber.Ctx) error
	GetStoryByUser(c *fiber.Ctx) error
	DeleteStory(c *fiber.Ctx) error
	PresignURL(c *fiber.Ctx) error
}

func NewStoryGatewayHandler(sc sc.StoryServiceClient, prof profile.ProfileServiceClient) StoryGatewayHandler {
	return &storyGatewayHandler{
		sc:   sc,
		prof: prof,
	}
}
