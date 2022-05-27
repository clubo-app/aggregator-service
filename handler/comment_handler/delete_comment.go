package commenthandler

import (
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	"github.com/clubo-app/protobuf/comment"
	"github.com/gofiber/fiber/v2"
)

func (h commentGatewayHandler) DeleteComment(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	cId := c.Params("id")
	pId := c.Params("pId")

	res, err := h.cc.DeleteComment(c.Context(), &comment.DeleteCommentRequest{AuthorId: user.Sub, PartyId: pId, CommentId: cId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)

}
