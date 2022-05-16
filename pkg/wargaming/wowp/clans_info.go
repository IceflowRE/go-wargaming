package wowp

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type ClansInfo struct {
	// Time when clan details were updated
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Clan Tag
	Tag string `json:"tag,omitempty"`
	// Time (UTC) when clan name was changed
	RenamedAt wgnTime.UnixTime `json:"renamed_at,omitempty"`
	// Old clan tag
	OldTag string `json:"old_tag,omitempty"`
	// Old clan name
	OldName string `json:"old_name,omitempty"`
	// Clan Name
	Name string `json:"name,omitempty"`
	// List of clan players' IDs
	MembersIds []int `json:"members_ids,omitempty"`
	// Number of clan members
	MembersCount int `json:"members_count,omitempty"`
	// Clan members.
	// An extra field.
	Members struct {
		// Technical position name
		Role string `json:"role,omitempty"`
		// Date when player joined clan
		JoinedAt wgnTime.UnixTime `json:"joined_at,omitempty"`
		// Player name
		AccountName string `json:"account_name,omitempty"`
		// User ID
		AccountId int `json:"account_id,omitempty"`
	} `json:"members,omitempty"`
	// Commander's name
	LeaderName string `json:"leader_name,omitempty"`
	// Commander ID
	LeaderId int `json:"leader_id,omitempty"`
	// Clan has been deleted. The deleted clan data contains valid values for the following fields only: clan_id, is_clan_disbanded, updated_at.
	IsClanDisbanded bool `json:"is_clan_disbanded,omitempty"`
	// Clan Description
	Description string `json:"description,omitempty"`
	// Clan creator's name
	CreatorName string `json:"creator_name,omitempty"`
	// Clan creator ID
	CreatorId int `json:"creator_id,omitempty"`
	// Clan creation date
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
} 
