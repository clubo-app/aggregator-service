package commenthandler

import (
	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	cg "github.com/clubo-app/protobuf/comment"
	ug "github.com/clubo-app/protobuf/user"
	"github.com/gofiber/fiber/v2"
)

func (h commentGatewayHandler) CreateReply(c *fiber.Ctx) error {
	req := new(cg.CreateReplyRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	cId := c.Params("id")
	user := middleware.ParseUser(c)

	req.AuthorId = user.Sub
	req.CommentId = cId

	r, err := h.cc.CreateReply(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, _ := h.uc.GetProfile(c.Context(), &ug.GetProfileRequest{Id: r.AuthorId})

	ar := datastruct.AggregatedReply{
		Id:        r.Id,
		CommentId: r.CommentId,
		Author:    profileRes,
		Body:      r.Body,
		CreatedAt: r.CreatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(ar)
}
