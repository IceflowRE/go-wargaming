package wargaming

import (
	"strings"
)

type WowsEncyclopediaCrewranks struct {
	// Experience
	Experience int `json:"experience,omitempty"`
	// Rank name
	Name string `json:"name,omitempty"`
	// Rank Name by subnation index
	Names map[string]string `json:"names,omitempty"`
	// Rank
	Rank int `json:"rank,omitempty"`
}

// WowsEncyclopediaCrewranks Method returns information about Commanders' ranks.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/crewranks
//
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
// nation:
//     Nation
func (client *Client) WowsEncyclopediaCrewranks(realm Realm, fields []string, language string, nation string) (*WowsEncyclopediaCrewranks, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"nation": nation,
	}

	var result *WowsEncyclopediaCrewranks
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/crewranks/", reqParam, &result)
	return result, err
}
