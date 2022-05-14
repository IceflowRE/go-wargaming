package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WowsClansAccountinfo struct {
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
		// Clan name
		Name string `json:"name,omitempty"`
		// Clan tag
		Tag string `json:"tag,omitempty"`
	} `json:"clan,omitempty"`
}

// WowsClansAccountinfo Method returns player clan data. Player clan data exist only for accounts, that were participating in clan activities: sent join requests, were clan members etc.
//
// https://developers.wargaming.net/reference/all/wows/clans/accountinfo
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
//     "cs" &mdash; Čeština 
//     "de" &mdash; Deutsch 
//     "en" &mdash; English 
//     "es" &mdash; Español 
//     "fr" &mdash; Français 
//     "ja" &mdash; 日本語 
//     "pl" &mdash; Polski 
//     "ru" &mdash; Русский (by default)
//     "th" &mdash; ไทย 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "zh-cn" &mdash; 中文 
//     "pt-br" &mdash; Português do Brasil 
//     "es-mx" &mdash; Español (México)
func (client *Client) WowsClansAccountinfo(realm Realm, accountId []int, extra []string, fields []string, language string) (*WowsClansAccountinfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowsClansAccountinfo
	err := client.doGetDataRequest(realm, "/wows/clans/accountinfo/", reqParam, &result)
	return result, err
}
