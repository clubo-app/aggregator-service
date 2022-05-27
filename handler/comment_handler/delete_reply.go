package commenthandler

import (
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	"github.com/clubo-app/protobuf/comment"
	"github.com/gofiber/fiber/v2"
)

func (h *commentGatewayHandler) DeleteReply(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	cId := c.Params("id")
	rId := c.Params("rId")

	res, err := h.cc.DeleteReply(c.Context(), &comment.DeleteReplyRequest{AuthorId: user.Sub, CommentId: cId, ReplyId: rId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
