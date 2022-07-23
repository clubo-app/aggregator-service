package datastruct

import "github.com/clubo-app/protobuf/profile"

type AggregatedProfile struct {
	Id          string `json:"id,omitempty"`
	Username    string `json:"username,omitempty"`
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	FriendCount uint32 `json:"friend_count"`
	// the friendship status is a pointer so it won't get marshalled into json if it's nil
	FriendshipStatus *FriendshipStatus `json:"friendship_status,omitempty"`
}

func ProfileToAgg(p *profile.Profile) AggregatedProfile {
	if p == nil {
		return AggregatedProfile{}
	}
	return AggregatedProfile{
		Id:        p.Id,
		Username:  p.Username,
		Firstname: p.Lastname,
		Lastname:  p.Lastname,
		Avatar:    p.Avatar,
	}
}

func (p AggregatedProfile) AddFs(fs FriendshipStatus) AggregatedProfile {
	p.FriendshipStatus = &fs
	return p
}

func (p AggregatedProfile) AddFCount(c uint32) AggregatedProfile {
	p.FriendCount = c
	return p
}

type PagedAggregatedProfile struct {
	AggregatedProfile []AggregatedProfile `json:"profiles"`
	NextPage          string              `json:"nextPage"`
}
