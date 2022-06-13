package storyhandler

import (
	"time"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	pg "github.com/clubo-app/protobuf/profile"
	sg "github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h storyGatewayHandler) GetStory(c *fiber.Ctx) error {
	sId := c.Params("id")

	s, err := h.sc.GetStory(c.Context(), &sg.GetStoryRequest{StoryId: sId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	// Get all profiles of the tagged people and the story creator in one call
	ids := append(s.TaggedFriends, s.UserId)

	profilesRes, err := h.prf.GetManyProfiles(c.Context(), &pg.GetManyProfilesRequest{Ids: ids})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	// Remove the creator of the story from the returned array and create a filtered list with only the profiles of the tagged people.
	// Separately store the profile of the creator of the story
	profile := new(pg.Profile)
	taggedFriends := make([]*pg.Profile, len(profilesRes.Profiles))
	for _, p := range profilesRes.Profiles {
		if p.Id != s.UserId {
			taggedFriends = append(taggedFriends, p)
		} else {
			profile = p
		}
	}

	res := datastruct.AggregatedStory{
		Id:            s.Id,
		PartyId:       s.PartyId,
		Creator:       profile,
		Url:           s.Url,
		TaggedFriends: taggedFriends,
		CreatedAt:     s.CreatedAt.AsTime().UTC().Format(time.RFC3339),
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
