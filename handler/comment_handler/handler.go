package commenthandler

import (
	cg "github.com/clubo-app/protobuf/comment"
	"github.com/clubo-app/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

type commentGatewayHandler struct {
	cc  cg.CommentServiceClient
	prf profile.ProfileServiceClient
}

type CommentGatewayHandler interface {
	CreateComment(c *fiber.Ctx) error
	DeleteComment(c *fiber.Ctx) error
	GetCommentByParty(c *fiber.Ctx) error
	CreateReply(c *fiber.Ctx) error
	DeleteReply(c *fiber.Ctx) error
	GetReplyByComment(c *fiber.Ctx) error
}

func NewCommentGatewayHandler(cc cg.CommentServiceClient, prf profile.ProfileServiceClient) CommentGatewayHandler {
	return &commentGatewayHandler{
		cc:  cc,
		prf: prf,
	}
}
