package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WowsAccountAchievements struct {
	// Battle achievements earned
	Battle map[string]string `json:"battle,omitempty"`
	// Achievement progress
	Progress map[string]string `json:"progress,omitempty"`
}

// WowsAccountAchievements Method returns information about players' achievements. Accounts with hidden game profiles are excluded from response. Hidden profiles are listed in the field meta.hidden.
//
// https://developers.wargaming.net/reference/all/wows/account/achievements
//
// account_id:
//     Player account ID. Maximum limit: 100. Min value is 1.
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
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
func (client *Client) WowsAccountAchievements(realm Realm, accountId []int, accessToken string, fields []string, language string) (*WowsAccountAchievements, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowsAccountAchievements
	err := client.doGetDataRequest(realm, "/wows/account/achievements/", reqParam, &result)
	return result, err
}
