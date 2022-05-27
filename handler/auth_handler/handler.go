package authhandler

import (
	ug "github.com/clubo-app/protobuf/user"
	"github.com/gofiber/fiber/v2"
)

type authGatewayHandler struct {
	uc ug.UserServiceClient
}

type AuthGatewayHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	GoogleLogin(c *fiber.Ctx) error
	VerifyEmail(c *fiber.Ctx) error
}

func NewAuthGatewayHandler(uc ug.UserServiceClient) AuthGatewayHandler {
	return &authGatewayHandler{
		uc: uc,
	}
}
