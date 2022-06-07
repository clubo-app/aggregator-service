package commenthandler

import (
	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	cg "github.com/clubo-app/protobuf/comment"
	"github.com/clubo-app/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

func (h commentGatewayHandler) CreateComment(c *fiber.Ctx) error {
	req := new(cg.CreateCommentRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	pId := c.Params("id")
	user := middleware.ParseUser(c)

	req.AuthorId = user.Sub
	req.PartyId = pId

	co, err := h.cc.CreateComment(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, _ := h.prf.GetProfile(c.Context(), &profile.GetProfileRequest{Id: co.AuthorId})

	ac := datastruct.AggregatedComment{
		Id:        co.Id,
		PartyId:   co.PartyId,
		Author:    profileRes,
		Body:      co.Body,
		CreatedAt: co.CreatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(ac)
}
