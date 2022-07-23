package commenthandler

import (
	"strconv"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	cg "github.com/clubo-app/protobuf/comment"
	"github.com/clubo-app/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

func (h commentGatewayHandler) GetCommentByParty(c *fiber.Ctx) error {
	pId := c.Params("id")
	nextPage := c.Query("nextPage")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)
	if limit > 20 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 20")
	}

	cs, err := h.cc.GetCommentByParty(c.Context(), &cg.GetByPartyRequest{PartyId: pId, NextPage: nextPage, Limit: limit})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	commentAuthors := make([]string, len(cs.Comments))
	for i, c := range cs.Comments {
		commentAuthors[i] = c.AuthorId
	}

	ps, _ := h.prf.GetManyProfilesMap(c.Context(), &profile.GetManyProfilesRequest{Ids: utils.UniqueStringSlice(commentAuthors)})

	aggC := make([]datastruct.AggregatedComment, len(cs.Comments))
	for i, c := range cs.Comments {
		aggC[i] = datastruct.AggregatedComment{
			Id:        c.Id,
			PartyId:   c.PartyId,
			Author:    ps.Profiles[c.AuthorId],
			Body:      c.Body,
			CreatedAt: c.CreatedAt,
		}
	}

	res := datastruct.PagedAggregatedComment{
		Comments: aggC,
		NextPage: cs.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
