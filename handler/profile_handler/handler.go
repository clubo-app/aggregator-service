package profilehandler

import (
	ag "github.com/clubo-app/protobuf/auth"
	pg "github.com/clubo-app/protobuf/profile"
	rg "github.com/clubo-app/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

type profileGatewayHandler struct {
	pc pg.ProfileServiceClient
	rc rg.RelationServiceClient
	ac ag.AuthServiceClient
}

type ProfileGatewayHandler interface {
	GetMe(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	UsernameTaken(c *fiber.Ctx) error
}

func NewUserGatewayHandler(pc pg.ProfileServiceClient, rc rg.RelationServiceClient, ac ag.AuthServiceClient) ProfileGatewayHandler {
	return &profileGatewayHandler{
		pc: pc,
		rc: rc,
		ac: ac,
	}
}
