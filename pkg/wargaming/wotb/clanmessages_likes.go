package wotb

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type ClanmessagesLikes struct {
	// Date when message was liked
	LikedAt wgnTime.UnixTime `json:"liked_at,omitempty"`
	// Liker Account ID
	AccountId int `json:"account_id,omitempty"`
} 
