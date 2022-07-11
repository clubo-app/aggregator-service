package authhandler

import (
	ag "github.com/clubo-app/protobuf/auth"
	pg "github.com/clubo-app/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

type authGatewayHandler struct {
	ac ag.AuthServiceClient
	pc pg.ProfileServiceClient
}

type AuthGatewayHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	GoogleLogin(c *fiber.Ctx) error
	VerifyEmail(c *fiber.Ctx) error
	RefreshAccessToken(c *fiber.Ctx) error
}

func NewAuthGatewayHandler(ac ag.AuthServiceClient, pc pg.ProfileServiceClient) AuthGatewayHandler {
	return &authGatewayHandler{
		ac: ac,
		pc: pc,
	}
}
