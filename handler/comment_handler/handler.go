package commenthandler

import (
	cg "github.com/clubo-app/protobuf/comment"
	uc "github.com/clubo-app/protobuf/user"
	"github.com/gofiber/fiber/v2"
)

type commentGatewayHandler struct {
	cc cg.CommentServiceClient
	uc uc.UserServiceClient
}

type CommentGatewayHandler interface {
	CreateComment(c *fiber.Ctx) error
	DeleteComment(c *fiber.Ctx) error
	GetCommentByParty(c *fiber.Ctx) error
	CreateReply(c *fiber.Ctx) error
	DeleteReply(c *fiber.Ctx) error
	GetReplyByComment(c *fiber.Ctx) error
}

func NewCommentGatewayHandler(cc cg.CommentServiceClient, uc uc.UserServiceClient) CommentGatewayHandler {
	return &commentGatewayHandler{
		cc: cc,
		uc: uc,
	}
}
