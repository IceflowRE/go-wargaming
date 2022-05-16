package wgn

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type ClansMemberhistory struct {
	// Last position in clan
	Role string `json:"role,omitempty"`
	// Date when player left clan
	LeftAt wgnTime.UnixTime `json:"left_at,omitempty"`
	// Date when player joined clan
	JoinedAt wgnTime.UnixTime `json:"joined_at,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
} 
