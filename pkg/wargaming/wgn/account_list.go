package wgn

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type AccountList struct {
	// Player name
	Nickname string `json:"nickname,omitempty"`
	// List of games player has played
	Games []string `json:"games,omitempty"`
	// Date when player's account was created
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Player ID
	AccountId int `json:"account_id,omitempty"`
} 
