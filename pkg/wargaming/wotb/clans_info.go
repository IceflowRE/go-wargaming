package wotb

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type ClansInfo struct {
	// Time when clan details were updated
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Clan tag
	Tag string `json:"tag,omitempty"`
	// Time (UTC) when clan name was changed
	RenamedAt wgnTime.UnixTime `json:"renamed_at,omitempty"`
	// Clan recruiting policy.
	// Valid values:
	// 
	// open - free to join, if statistics completely meet the required threshold values (by default)
	// restricted - applications to join clan can be sent
	RecruitingPolicy string `json:"recruiting_policy,omitempty"`
	// Threshold statistics values to join clan. Contains null if threshold values are not set.
	RecruitingOptions struct {
		// Victories/Battles ratio
		WinsRatio int `json:"wins_ratio,omitempty"`
		// Minimum vehicle Tier of player
		VehiclesLevel int `json:"vehicles_level,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average damage per battle
		AverageDamage int `json:"average_damage,omitempty"`
		// Average number of battles per day
		AverageBattlesPerDay int `json:"average_battles_per_day,omitempty"`
	} `json:"recruiting_options,omitempty"`
	// Old clan tag
	OldTag string `json:"old_tag,omitempty"`
	// Old clan name
	OldName string `json:"old_name,omitempty"`
	// Clan name
	Name string `json:"name,omitempty"`
	// Clan motto
	Motto string `json:"motto,omitempty"`
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
	// Emblems set ID
	EmblemSetId int `json:"emblem_set_id,omitempty"`
	// Clan description
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
