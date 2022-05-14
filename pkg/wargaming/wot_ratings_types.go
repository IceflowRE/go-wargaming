package wargaming

import (
	"strings"
)

type WotRatingsTypes struct {
	// Rating categories
	RankFields []string `json:"rank_fields,omitempty"`
	// Rating threshold
	Threshold int `json:"threshold,omitempty"`
	// Rating period
	Type_ string `json:"type,omitempty"`
}

// WotRatingsTypes Deprecated: Attention! The method is deprecated.
// Method returns dictionary of rating periods and ratings details.
//
// https://developers.wargaming.net/reference/all/wot/ratings/types
//
// battle_type:
//     Battle types. Default is "default". Valid values:
//     
//     "company" &mdash; Tank Company Battles 
//     "random" &mdash; Random Battles 
//     "team" &mdash; Team Battles 
//     "default" &mdash; any battle type (by default)
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
func (client *Client) WotRatingsTypes(realm Realm, battleType string, fields []string, language string) (*WotRatingsTypes, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"battle_type": battleType,
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotRatingsTypes
	err := client.doGetDataRequest(realm, "/wot/ratings/types/", reqParam, &result)
	return result, err
}
