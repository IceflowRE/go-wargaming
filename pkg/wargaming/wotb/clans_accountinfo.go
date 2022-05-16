package wotb

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type ClansAccountinfo struct {
	// Technical position name
	Role string `json:"role,omitempty"`
	// Date when player joined clan
	JoinedAt wgnTime.UnixTime `json:"joined_at,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Short info about clan.
	// An extra field.
	Clan struct {
		// Clan tag
		Tag string `json:"tag,omitempty"`
		// Clan name
		Name string `json:"name,omitempty"`
		// Number of clan members
		MembersCount int `json:"members_count,omitempty"`
		// Emblems set ID
		EmblemSetId int `json:"emblem_set_id,omitempty"`
		// Clan creation date
		CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
		// Clan ID
		ClanId int `json:"clan_id,omitempty"`
	} `json:"clan,omitempty"`
	// Player name
	AccountName string `json:"account_name,omitempty"`
	// User ID
	AccountId int `json:"account_id,omitempty"`
} 
