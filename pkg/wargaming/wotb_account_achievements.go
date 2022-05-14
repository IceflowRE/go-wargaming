package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotbAccountAchievements struct {
	// Achievements earned
	Achievements map[string]string `json:"achievements,omitempty"`
	// Maximum values of achievement series
	MaxSeries map[string]string `json:"max_series,omitempty"`
}

// WotbAccountAchievements Method returns players' achievement details.
//
// https://developers.wargaming.net/reference/all/wotb/account/achievements
//
// account_id:
//     Player account ID. Maximum limit: 100.
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
func (client *Client) WotbAccountAchievements(realm Realm, accountId []int, fields []string, language string) (*WotbAccountAchievements, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotbAccountAchievements
	err := client.doGetDataRequest(realm, "/wotb/account/achievements/", reqParam, &result)
	return result, err
}
