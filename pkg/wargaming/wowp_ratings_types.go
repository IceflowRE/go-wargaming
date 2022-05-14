package wargaming

import (
	"strings"
)

type WowpRatingsTypes struct {
	// Rating categories
	RankFields []string `json:"rank_fields,omitempty"`
	// Rating threshold
	Threshold int `json:"threshold,omitempty"`
	// Rating period
	Type_ string `json:"type,omitempty"`
}

// WowpRatingsTypes Method returns dictionary of rating periods and ratings details.
//
// https://developers.wargaming.net/reference/all/wowp/ratings/types
//
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
func (client *Client) WowpRatingsTypes(realm Realm, fields []string, language string) (*WowpRatingsTypes, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowpRatingsTypes
	err := client.doGetDataRequest(realm, "/wowp/ratings/types/", reqParam, &result)
	return result, err
}
