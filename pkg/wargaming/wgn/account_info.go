package wgn

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type AccountInfo struct {
	// Player's private data
	Private struct {
		// Premium Account expiration date
		PremiumExpiresAt wgnTime.UnixTime `json:"premium_expires_at,omitempty"`
		// Current gold balance
		Gold int `json:"gold,omitempty"`
		// Amount of Free Experience
		FreeXp int `json:"free_xp,omitempty"`
	} `json:"private,omitempty"`
	// Player name
	Nickname string `json:"nickname,omitempty"`
	// List of games player has played
	Games []string `json:"games,omitempty"`
	// Date when player's account was created
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Player ID
	AccountId int `json:"account_id,omitempty"`
} 
