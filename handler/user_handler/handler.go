package profilehandler

import (
	rg "github.com/clubo-app/protobuf/relation"
	ug "github.com/clubo-app/protobuf/user"
	"github.com/gofiber/fiber/v2"
)

type userGatewayHandler struct {
	uc ug.UserServiceClient
	rc rg.RelationServiceClient
}

type UserGatewayHandler interface {
	GetMe(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	UsernameTaken(c *fiber.Ctx) error
}

func NewUserGatewayHandler(uc ug.UserServiceClient, rc rg.RelationServiceClient) UserGatewayHandler {
	return &userGatewayHandler{
		uc: uc,
		rc: rc,
	}
}
