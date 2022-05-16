package wowp

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type AccountAchievements struct {
	// Achievements earned
	Achievements struct {
		// Most number
		MaxCount int `json:"max_count,omitempty"`
		// Last time when an achievement was received
		LastAt wgnTime.UnixTime `json:"last_at,omitempty"`
		// First time when an achievement was received
		FirstAt wgnTime.UnixTime `json:"first_at,omitempty"`
		// Number
		Count int `json:"count,omitempty"`
	} `json:"achievements,omitempty"`
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
} 
