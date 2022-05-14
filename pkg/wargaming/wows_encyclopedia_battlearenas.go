package wargaming

import (
	"strings"
)

type WowsEncyclopediaBattlearenas struct {
	// Map ID
	BattleArenaId int `json:"battle_arena_id,omitempty"`
	// Map description
	Description string `json:"description,omitempty"`
	// URL to the map icon
	Icon string `json:"icon,omitempty"`
	// Map name
	Name string `json:"name,omitempty"`
}

// WowsEncyclopediaBattlearenas Method returns the information about maps.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/battlearenas
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
func (client *Client) WowsEncyclopediaBattlearenas(realm Realm, fields []string, language string) (*WowsEncyclopediaBattlearenas, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowsEncyclopediaBattlearenas
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/battlearenas/", reqParam, &result)
	return result, err
}
