package datastruct

type FriendshipStatus struct {
	IsFriend        bool   `json:"is_friend"`
	OutgoingRequest bool   `json:"outgoing_request"`
	RequestedAt     string `json:"requested_at,omitempty"`
	AcceptedAt      string `json:"accepted_at,omitempty"`
}

type AggregatedProfile struct {
	Id               string           `json:"id,omitempty"`
	Username         string           `json:"username,omitempty"`
	Firstname        string           `json:"firstname,omitempty"`
	Lastname         string           `json:"lastname,omitempty"`
	Avatar           string           `json:"avatar,omitempty"`
	FriendCount      uint32           `json:"friend_count,omitempty"`
	FriendshipStatus FriendshipStatus `json:"friendship_status,omitempty"`
}
