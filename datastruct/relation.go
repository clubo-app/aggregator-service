package datastruct

import (
	"time"

	"github.com/clubo-app/protobuf/relation"
)

type FriendshipStatus struct {
	IsFriend        bool   `json:"is_friend"`
	IncomingRequest bool   `json:"incoming_request"`
	OutgoingRequest bool   `json:"outgoing_request"`
	RequestedAt     string `json:"requested_at,omitempty"`
	AcceptedAt      string `json:"accepted_at,omitempty"`
}

func ParseFriendShipStatus(perspective string, fr *relation.FriendRelation) (fs FriendshipStatus) {
	// if true, we know if the user has request from friend or if user is friends with friend
	if perspective == fr.UserId && !fr.Accepted {
		fs.IncomingRequest = true
		if !fr.RequestedAt.AsTime().IsZero() {
			fs.RequestedAt = fr.RequestedAt.AsTime().UTC().Format(time.RFC3339)
		}
	} else if perspective == fr.FriendId && !fr.Accepted {
		fs.OutgoingRequest = true
		if !fr.RequestedAt.AsTime().IsZero() {
			fs.RequestedAt = fr.RequestedAt.AsTime().UTC().Format(time.RFC3339)
		}
	}

	if fr.Accepted {
		fs.IsFriend = true
		if !fr.AcceptedAt.AsTime().IsZero() {
			fs.AcceptedAt = fr.AcceptedAt.AsTime().UTC().Format(time.RFC3339)
		}
	}

	return fs
}
