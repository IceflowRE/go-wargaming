package wowp

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type ClansList struct {
	// Clan Tag
	Tag string `json:"tag,omitempty"`
	// Clan Name
	Name string `json:"name,omitempty"`
	// Number of clan members
	MembersCount int `json:"members_count,omitempty"`
	// Clan creation date
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
} 
