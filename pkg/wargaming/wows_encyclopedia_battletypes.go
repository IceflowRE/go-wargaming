package wargaming

import (
	"strings"
)

type WowsEncyclopediaBattletypes struct {
	// Battle type description
	Description string `json:"description,omitempty"`
	// URL to image
	Image string `json:"image,omitempty"`
	// Battle type name
	Name string `json:"name,omitempty"`
	// Battle type tag
	Tag string `json:"tag,omitempty"`
}

// WowsEncyclopediaBattletypes Method returns information about battle types.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/battletypes
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
func (client *Client) WowsEncyclopediaBattletypes(realm Realm, fields []string, language string) (*WowsEncyclopediaBattletypes, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowsEncyclopediaBattletypes
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/battletypes/", reqParam, &result)
	return result, err
}
