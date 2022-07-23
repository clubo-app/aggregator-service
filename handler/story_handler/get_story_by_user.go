package storyhandler

import (
	"strconv"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/protobuf/party"
	"github.com/clubo-app/protobuf/profile"
	"github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h storyGatewayHandler) GetStoryByUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)
	if limit > 20 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 20")
	}

	nextPage := c.Query("nextPage")

	stories, err := h.sc.GetByUser(c.Context(), &story.GetByUserRequest{
		UserId:   userId,
		NextPage: nextPage,
		Limit:    uint32(limit),
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profile, _ := h.prf.GetProfile(c.Context(), &profile.GetProfileRequest{Id: userId})

	// Get all ids of the parties of the stories
	ids := make([]string, len(stories.Stories))
	for i, s := range stories.Stories {
		ids[i] = s.PartyId
	}
	// here we aggregated all the profiles if the story creators
	parties, _ := h.pc.GetManyPartiesMap(c.Context(), &party.GetManyPartiesRequest{Ids: utils.UniqueStringSlice(ids)})

	aggS := make([]datastruct.AggregatedStory, len(stories.Stories))
	for i, s := range stories.Stories {
		s := datastruct.
			StoryToAgg(s).
			AddCreator(datastruct.ProfileToAgg(profile)).
			AddParty(datastruct.AggregatedParty{Id: s.PartyId})

		if parties != nil {
			s = s.AddParty(datastruct.PartyToAgg(parties.Parties[s.Party.Id]))
		}

		aggS[i] = s
	}

	res := datastruct.PagedAggregatedStory{
		Stories:  aggS,
		NextPage: stories.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
