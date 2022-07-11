package datastruct

type AggregatedProfile struct {
	Id               string           `json:"id,omitempty"`
	Username         string           `json:"username,omitempty"`
	Firstname        string           `json:"firstname,omitempty"`
	Lastname         string           `json:"lastname,omitempty"`
	Avatar           string           `json:"avatar,omitempty"`
	FriendCount      uint32           `json:"friend_count"`
	FriendshipStatus FriendshipStatus `json:"friendship_status,omitempty"`
}

type PagedAggregatedProfile struct {
	AggregatedProfile []AggregatedProfile `json:"profiles"`
	NextPage          string              `json:"nextPage"`
}
