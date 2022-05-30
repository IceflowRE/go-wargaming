package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wowp"
	"strings"
)

// WowpAccountAchievements Method returns players' achievement details.
//
// https://developers.wargaming.net/reference/all/wowp/account/achievements
//
// accountId:
//     Player account ID. Maximum limit: 100.
//     Parameter is required.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "en" - English 
//     "ru" - Русский (by default)
//     "pl" - Polski 
//     "de" - Deutsch 
//     "fr" - Français 
//     "es" - Español 
//     "zh-cn" - 简体中文 
//     "tr" - Türkçe 
//     "cs" - Čeština 
//     "th" - ไทย 
//     "vi" - Tiếng Việt 
//     "ko" - 한국어
func (client *Client) WowpAccountAchievements(realm Realm, accountId []int, fields []string, language string) (*wowp.AccountAchievements, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wowp.AccountAchievements
	err := client.doGetDataRequest(realm, "/wowp/account/achievements/", reqParam, &result)
	return result, err
}
