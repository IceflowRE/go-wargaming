package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotxClansInfo struct {
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Clan color in HEX #RRGGBB
	Color string `json:"color,omitempty"`
	// Clan creation date
	CreatedAt UnixTime `json:"created_at,omitempty"`
	// Clan creator ID
	CreatorId int `json:"creator_id,omitempty"`
	// Clan creator's name
	CreatorName string `json:"creator_name,omitempty"`
	// Clan description
	Description string `json:"description,omitempty"`
	// Clan emblems set ID
	EmblemSetId int `json:"emblem_set_id,omitempty"`
	// Clan has been deleted. The deleted clan data contains valid values for the following fields only: clan_id, is_clan_disbanded, updated_at.
	IsClanDisbanded bool `json:"is_clan_disbanded,omitempty"`
	// Commander ID
	LeaderId int `json:"leader_id,omitempty"`
	// Commander's name
	LeaderName string `json:"leader_name,omitempty"`
	// Number of clan members
	MembersCount int `json:"members_count,omitempty"`
	// List of clan players' IDs
	MembersIds []int `json:"members_ids,omitempty"`
	// Clan motto
	Motto string `json:"motto,omitempty"`
	// Clan name
	Name string `json:"name,omitempty"`
	// Old clan name
	OldName string `json:"old_name,omitempty"`
	// Old clan tag
	OldTag string `json:"old_tag,omitempty"`
	// Clan recruiting policy.
	// Valid values:
	// 
	// open - free to join, if statistics completely meet the required threshold values (by default)
	// restricted - applications to join clan can be sent
	// closed - applications to join clan cannot be sent
	RecruitingPolicy string `json:"recruiting_policy,omitempty"`
	// Time (UTC) when clan name was changed
	RenamedAt UnixTime `json:"renamed_at,omitempty"`
	// Clan tag
	Tag string `json:"tag,omitempty"`
	// Time when clan details were updated
	UpdatedAt UnixTime `json:"updated_at,omitempty"`
	// Threshold statistics values to join clan. Contains null if threshold values are not set.
	JoiningOptions struct {
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average number of battles per day
		BattlesPerDay float32 `json:"battles_per_day,omitempty"`
		// Battles survived
		BattlesSurvived int `json:"battles_survived,omitempty"`
		// Average damage per battle
		DamagePerBattle float32 `json:"damage_per_battle,omitempty"`
		// Hit ratio
		HitsRatio float32 `json:"hits_ratio,omitempty"`
		// Victories/Battles ratio
		WinsRatio float32 `json:"wins_ratio,omitempty"`
		// Average experience per battle
		XpPerBattle float32 `json:"xp_per_battle,omitempty"`
	} `json:"joining_options,omitempty"`
	// Clan members.
	// An extra field.
	Members struct {
		// User ID
		AccountId int `json:"account_id,omitempty"`
		// Player name
		AccountName string `json:"account_name,omitempty"`
		// Date when player joined clan
		JoinedAt UnixTime `json:"joined_at,omitempty"`
		// Technical position name
		Role string `json:"role,omitempty"`
	} `json:"members,omitempty"`
}

// WotxClansInfo Method returns detailed clan information.
//
// https://developers.wargaming.net/reference/all/wotx/clans/info
//
// clan_id:
//     Clan ID. Maximum limit: 100. Min value is 1.
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "members"
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "en". Valid values:
//     
//     "en" &mdash; English (by default)
//     "ru" &mdash; Русский 
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "tr" &mdash; Türkçe
func (client *Client) WotxClansInfo(realm Realm, clanId []int, extra []string, fields []string, language string) (*WotxClansInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": utils.SliceIntToString(clanId, ","),
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotxClansInfo
	err := client.doGetDataRequest(realm, "/wotx/clans/info/", reqParam, &result)
	return result, err
}
