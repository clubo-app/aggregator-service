package datastruct

import (
	"time"

	sg "github.com/clubo-app/protobuf/story"
)

type AggregatedStory struct {
	Id            string              `json:"id,omitempty"`
	Party         *AggregatedParty    `json:"party,omitempty"`
	Creator       *AggregatedProfile  `json:"creator,omitempty"`
	Url           string              `json:"url,omitempty"`
	TaggedFriends []AggregatedProfile `json:"tagged_friends,omitempty"`
	CreatedAt     string              `json:"created_at,omitempty"`
}

func StoryToAgg(s *sg.Story) AggregatedStory {
	if s == nil {
		return AggregatedStory{}
	}
	return AggregatedStory{
		Id:        s.Id,
		Url:       s.Url,
		CreatedAt: s.CreatedAt.AsTime().UTC().Format(time.RFC3339),
	}
}

func (s AggregatedStory) AddCreator(p AggregatedProfile) AggregatedStory {
	if s.Id != "" {
		s.Creator = &p
	}
	return s
}

func (s AggregatedStory) AddParty(p AggregatedParty) AggregatedStory {
	if p.Id != "" {
		s.Party = &p
	}
	return s
}

func (s AggregatedStory) AddFriends(fs []AggregatedProfile) AggregatedStory {
	if fs[0].Id != "" {
		s.TaggedFriends = fs
	}
	return s
}

type PagedAggregatedStory struct {
	Stories  []AggregatedStory `json:"stories,omitempty"`
	NextPage string            `json:"nextPage"`
}
