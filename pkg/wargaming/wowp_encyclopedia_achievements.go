package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wowp"
	"strings"
)

// WowpEncyclopediaAchievements Method returns dictionary of achievements from Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/achievements
//
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
func (client *Client) WowpEncyclopediaAchievements(realm Realm, fields []string, language string) (*wowp.EncyclopediaAchievements, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wowp.EncyclopediaAchievements
	err := client.doGetDataRequest(realm, "/wowp/encyclopedia/achievements/", reqParam, &result)
	return result, err
}
