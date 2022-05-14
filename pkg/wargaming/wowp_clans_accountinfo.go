package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WowpClansAccountinfo struct {
	// User ID
	AccountId int `json:"account_id,omitempty"`
	// Player name
	AccountName string `json:"account_name,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Date when player joined clan
	JoinedAt UnixTime `json:"joined_at,omitempty"`
	// Technical position name
	Role string `json:"role,omitempty"`
	// Short info about clan.
	// An extra field.
	Clan struct {
		// Clan ID
		ClanId int `json:"clan_id,omitempty"`
		// Clan creation date
		CreatedAt UnixTime `json:"created_at,omitempty"`
		// Number of clan members
		MembersCount int `json:"members_count,omitempty"`
		// Clan Name
		Name string `json:"name,omitempty"`
		// Clan Tag
		Tag string `json:"tag,omitempty"`
	} `json:"clan,omitempty"`
}

// WowpClansAccountinfo Method returns player clan data. Player clan data exist only for accounts, that were participating in clan activities: sent join requests, were clan members etc.
//
// https://developers.wargaming.net/reference/all/wowp/clans/accountinfo
//
// account_id:
//     Account ID. Maximum limit: 100. Min value is 1.
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "clan"
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
func (client *Client) WowpClansAccountinfo(realm Realm, accountId []int, extra []string, fields []string, language string) (*WowpClansAccountinfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowpClansAccountinfo
	err := client.doGetDataRequest(realm, "/wowp/clans/accountinfo/", reqParam, &result)
	return result, err
}
