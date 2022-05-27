package commenthandler

import (
	"strconv"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	cg "github.com/clubo-app/protobuf/comment"
	"github.com/clubo-app/protobuf/user"
	"github.com/gofiber/fiber/v2"
)

func (h commentGatewayHandler) GetCommentByParty(c *fiber.Ctx) error {
	pId := c.Params("id")
	nextPage := c.Query("nextPage")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)

	cs, err := h.cc.GetCommentByParty(c.Context(), &cg.GetByPartyRequest{PartyId: pId, NextPage: nextPage, Limit: uint32(limit)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	var commentAuthors []string
	for _, c := range cs.Comments {
		commentAuthors = append(commentAuthors, c.AuthorId)
	}

	pRes, _ := h.uc.GetManyProfilesMap(c.Context(), &user.GetManyProfilesRequest{Ids: utils.UniqueStringSlice(commentAuthors)})

	aggC := make([]datastruct.AggregatedComment, len(cs.Comments))
	for i, c := range cs.Comments {
		if author, ok := pRes.Profiles[c.AuthorId]; ok {
			aggC[i] = datastruct.AggregatedComment{
				Id:        c.Id,
				PartyId:   c.PartyId,
				Author:    author,
				Body:      c.Body,
				CreatedAt: c.CreatedAt,
			}
		} else {
			aggC[i] = datastruct.AggregatedComment{
				Id:        c.Id,
				PartyId:   c.PartyId,
				Author:    &user.Profile{},
				Body:      c.Body,
				CreatedAt: c.CreatedAt,
			}
		}
	}

	res := datastruct.PagedAggregatedComment{
		Comments: aggC,
		NextPage: cs.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
