package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotbClansInfo struct {
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Clan creation date
	CreatedAt UnixTime `json:"created_at,omitempty"`
	// Clan creator ID
	CreatorId int `json:"creator_id,omitempty"`
	// Clan creator's name
	CreatorName string `json:"creator_name,omitempty"`
	// Clan description
	Description string `json:"description,omitempty"`
	// Emblems set ID
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
	RecruitingPolicy string `json:"recruiting_policy,omitempty"`
	// Time (UTC) when clan name was changed
	RenamedAt UnixTime `json:"renamed_at,omitempty"`
	// Clan tag
	Tag string `json:"tag,omitempty"`
	// Time when clan details were updated
	UpdatedAt UnixTime `json:"updated_at,omitempty"`
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
	// Threshold statistics values to join clan. Contains null if threshold values are not set.
	RecruitingOptions struct {
		// Average number of battles per day
		AverageBattlesPerDay int `json:"average_battles_per_day,omitempty"`
		// Average damage per battle
		AverageDamage int `json:"average_damage,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Minimum vehicle Tier of player
		VehiclesLevel int `json:"vehicles_level,omitempty"`
		// Victories/Battles ratio
		WinsRatio int `json:"wins_ratio,omitempty"`
	} `json:"recruiting_options,omitempty"`
}

// WotbClansInfo Method returns detailed clan information.
//
// https://developers.wargaming.net/reference/all/wotb/clans/info
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
//     Localization language. Default is "ru". Valid values:
//     
//     "en" &mdash; English 
//     "ru" &mdash; Русский (by default)
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "zh-cn" &mdash; 简体中文 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
func (client *Client) WotbClansInfo(realm Realm, clanId []int, extra []string, fields []string, language string) (*WotbClansInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": utils.SliceIntToString(clanId, ","),
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotbClansInfo
	err := client.doGetDataRequest(realm, "/wotb/clans/info/", reqParam, &result)
	return result, err
}
