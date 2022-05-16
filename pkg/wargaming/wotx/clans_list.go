package wotx

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type ClansList struct {
	// Clan tag
	Tag string `json:"tag,omitempty"`
	// Clan name
	Name string `json:"name,omitempty"`
	// Number of clan members
	MembersCount int `json:"members_count,omitempty"`
	// Clan emblems set ID
	EmblemSetId int `json:"emblem_set_id,omitempty"`
	// Clan creation date
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Clan color in HEX #RRGGBB
	Color string `json:"color,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
} 
