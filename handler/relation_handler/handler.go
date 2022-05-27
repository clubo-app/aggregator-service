package relationhandler

import (
	pg "github.com/clubo-app/protobuf/party"
	rg "github.com/clubo-app/protobuf/relation"
	ug "github.com/clubo-app/protobuf/user"
	"github.com/gofiber/fiber/v2"
)

type relationGatewayHandler struct {
	rc rg.RelationServiceClient
	pc pg.PartyServiceClient
	uc ug.UserServiceClient
}

type RelationGatewayHandler interface {
	FriendRequest(c *fiber.Ctx) error
	AcceptFriend(c *fiber.Ctx) error
	RemoveFriend(c *fiber.Ctx) error

	FavorParty(c *fiber.Ctx) error
	DefavorParty(c *fiber.Ctx) error
	GetFavoritePartiesByUser(c *fiber.Ctx) error
	GetFavorisingUsersByParty(c *fiber.Ctx) error
}

func NewRelationGatewayHandler(rc rg.RelationServiceClient, pc pg.PartyServiceClient, uc ug.UserServiceClient) RelationGatewayHandler {
	return &relationGatewayHandler{
		rc: rc,
		pc: pc,
		uc: uc,
	}
}
