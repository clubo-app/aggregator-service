package datastruct

type AggregatedStory struct {
	Id            string              `json:"id,omitempty"`
	PartyId       string              `json:"party_id,omitempty"`
	Creator       AggregatedProfile   `json:"creator,omitempty"`
	Url           string              `json:"url,omitempty"`
	TaggedFriends []AggregatedProfile `json:"tagged_friends,omitempty"`
	CreatedAt     string              `json:"created_at,omitempty"`
}

type PagedAggregatedStory struct {
	Stories  []AggregatedStory `json:"stories,omitempty"`
	NextPage string            `json:"nextPage"`
}
