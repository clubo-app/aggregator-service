package storyhandler

import (
	"time"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
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

	s, err := h.sc.CreateStory(c.Context(), &sg.CreateStoryRequest{
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
	var profile = new(pg.Profile)
	var fs []*pg.Profile
	for _, p := range profilesRes.Profiles {
		if p.Id != s.UserId {
			fs = append(fs, p)
		} else {
			profile = p
		}
	}

	res := datastruct.AggregatedStory{
		Id:            s.Id,
		PartyId:       s.PartyId,
		Creator:       profile,
		Url:           s.Url,
		TaggedFriends: fs,
		CreatedAt:     s.CreatedAt.AsTime().UTC().Format(time.RFC3339),
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
