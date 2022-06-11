package relationhandler

import (
	rg "github.com/clubo-app/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

type relationGatewayHandler struct {
	rc rg.RelationServiceClient
}

type RelationGatewayHandler interface {
	CreateFriendRequest(c *fiber.Ctx) error
	AcceptFriendRequest(c *fiber.Ctx) error
	DeclineFriendRequest(c *fiber.Ctx) error
	RemoveFriend(c *fiber.Ctx) error

	FavorParty(c *fiber.Ctx) error
	DefavorParty(c *fiber.Ctx) error
}

func NewRelationGatewayHandler(rc rg.RelationServiceClient) RelationGatewayHandler {
	return &relationGatewayHandler{
		rc: rc,
	}
}
