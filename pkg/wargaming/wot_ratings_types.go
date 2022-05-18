package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strings"
)

// WotRatingsTypes Method returns dictionary of rating periods and ratings details.
//
// https://developers.wargaming.net/reference/all/wot/ratings/types
//
// Deprecated: Attention! The method is deprecated.
//
// battleType:
//     Battle types. Default is "default". Valid values:
//     
//     "company" - Tank Company Battles 
//     "random" - Random Battles 
//     "team" - Team Battles 
//     "default" - any battle type (by default)
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
//     "zh-tw" - 繁體中文 
//     "tr" - Türkçe 
//     "cs" - Čeština 
//     "th" - ไทย 
//     "vi" - Tiếng Việt 
//     "ko" - 한국어
func (client *Client) WotRatingsTypes(realm Realm, battleType string, fields []string, language string) (*wot.RatingsTypes, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"battle_type": battleType,
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wot.RatingsTypes
	err := client.doGetDataRequest(realm, "/wot/ratings/types/", reqParam, &result)
	return result, err
}
