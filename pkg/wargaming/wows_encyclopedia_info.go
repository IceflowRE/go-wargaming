package wargaming

import (
	"strings"
)

type WowsEncyclopediaInfo struct {
	// Game client version
	GameVersion string `json:"game_version,omitempty"`
	// List of languages supported by encyclopedia methods
	Languages map[string]string `json:"languages,omitempty"`
	// Types of Modifications available
	ShipModifications map[string]string `json:"ship_modifications,omitempty"`
	// Types of modules available
	ShipModules map[string]string `json:"ship_modules,omitempty"`
	// Nations available
	ShipNations map[string]string `json:"ship_nations,omitempty"`
	// Types of ships available
	ShipTypes map[string]string `json:"ship_types,omitempty"`
	// Time when details on ships were updated
	ShipsUpdatedAt UnixTime `json:"ships_updated_at,omitempty"`
	// Images of ship types
	ShipTypeImages struct {
		// Ship type image
		Image string `json:"image,omitempty"`
		// Elite ships icon
		ImageElite string `json:"image_elite,omitempty"`
		// Premium ships icon
		ImagePremium string `json:"image_premium,omitempty"`
	} `json:"ship_type_images,omitempty"`
}

// WowsEncyclopediaInfo Method returns information about encyclopedia.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/info
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
func (client *Client) WowsEncyclopediaInfo(realm Realm, fields []string, language string) (*WowsEncyclopediaInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowsEncyclopediaInfo
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/info/", reqParam, &result)
	return result, err
}
