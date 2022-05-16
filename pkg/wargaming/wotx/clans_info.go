package wotx

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
	// closed - applications to join clan cannot be sent
	RecruitingPolicy string `json:"recruiting_policy,omitempty"`
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
	// Threshold statistics values to join clan. Contains null if threshold values are not set.
	JoiningOptions struct {
		// Average experience per battle
		XpPerBattle float32 `json:"xp_per_battle,omitempty"`
		// Victories/Battles ratio
		WinsRatio float32 `json:"wins_ratio,omitempty"`
		// Hit ratio
		HitsRatio float32 `json:"hits_ratio,omitempty"`
		// Average damage per battle
		DamagePerBattle float32 `json:"damage_per_battle,omitempty"`
		// Battles survived
		BattlesSurvived int `json:"battles_survived,omitempty"`
		// Average number of battles per day
		BattlesPerDay float32 `json:"battles_per_day,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
	} `json:"joining_options,omitempty"`
	// Clan has been deleted. The deleted clan data contains valid values for the following fields only: clan_id, is_clan_disbanded, updated_at.
	IsClanDisbanded bool `json:"is_clan_disbanded,omitempty"`
	// Clan emblems set ID
	EmblemSetId int `json:"emblem_set_id,omitempty"`
	// Clan description
	Description string `json:"description,omitempty"`
	// Clan creator's name
	CreatorName string `json:"creator_name,omitempty"`
	// Clan creator ID
	CreatorId int `json:"creator_id,omitempty"`
	// Clan creation date
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Clan color in HEX #RRGGBB
	Color string `json:"color,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
} 
