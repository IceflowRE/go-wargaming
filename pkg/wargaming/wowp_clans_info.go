package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WowpClansInfo struct {
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Clan creation date
	CreatedAt UnixTime `json:"created_at,omitempty"`
	// Clan creator ID
	CreatorId int `json:"creator_id,omitempty"`
	// Clan creator's name
	CreatorName string `json:"creator_name,omitempty"`
	// Clan Description
	Description string `json:"description,omitempty"`
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
	// Clan Name
	Name string `json:"name,omitempty"`
	// Old clan name
	OldName string `json:"old_name,omitempty"`
	// Old clan tag
	OldTag string `json:"old_tag,omitempty"`
	// Time (UTC) when clan name was changed
	RenamedAt UnixTime `json:"renamed_at,omitempty"`
	// Clan Tag
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
}

// WowpClansInfo Method returns detailed clan information.
//
// https://developers.wargaming.net/reference/all/wowp/clans/info
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
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
func (client *Client) WowpClansInfo(realm Realm, clanId []int, extra []string, fields []string, language string) (*WowpClansInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": utils.SliceIntToString(clanId, ","),
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowpClansInfo
	err := client.doGetDataRequest(realm, "/wowp/clans/info/", reqParam, &result)
	return result, err
}
