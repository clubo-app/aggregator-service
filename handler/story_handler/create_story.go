package storyhandler

import (
	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	pg "github.com/clubo-app/protobuf/profile"
	sg "github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

type CreatePartyReq struct {
	PartyId       string   `json:"party_id"`
	Url           string   `json:"url"`
	TaggedFriends []string `json:"tagged_friends"`
}

func (h storyGatewayHandler) CreateStory(c *fiber.Ctx) error {
	req := new(CreatePartyReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	user := middleware.ParseUser(c)

	s, err := h.sc.CreateStory(c.Context(), &sg.CreateStoryRequest{
		RequesterId:   user.Sub,
		PartyId:       req.PartyId,
		Url:           req.Url,
		TaggedFriends: req.TaggedFriends,
	})
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
	var profile datastruct.AggregatedProfile
	taggedFriends := make([]datastruct.AggregatedProfile, len(profilesRes.Profiles))
	for i, p := range profilesRes.Profiles {
		if p.Id != s.UserId {
			taggedFriends[i] = datastruct.ProfileToAgg(p)
		} else {
			profile = datastruct.ProfileToAgg(p)
		}
	}

	res := datastruct.
		StoryToAgg(s).
		AddCreator(profile).
		AddFriends(taggedFriends)

	return c.Status(fiber.StatusCreated).JSON(res)
}
