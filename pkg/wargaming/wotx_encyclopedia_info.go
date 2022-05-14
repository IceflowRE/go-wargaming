package wargaming

import (
	"strings"
)

type WotxEncyclopediaInfo struct {
	// Game client version
	GameVersion string `json:"game_version,omitempty"`
	// Nations available
	VehicleNations map[string]string `json:"vehicle_nations,omitempty"`
	// Available vehicle types
	VehicleTypes map[string]string `json:"vehicle_types,omitempty"`
	// Award sections
	AchievementSections struct {
		// Award section name
		Name string `json:"name,omitempty"`
		// Award section order
		Order int `json:"order,omitempty"`
	} `json:"achievement_sections,omitempty"`
}

// WotxEncyclopediaInfo Method returns information about Tankopedia.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/info
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "en". Valid values:
//     
//     "en" &mdash; English (by default)
//     "ru" &mdash; Русский 
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "tr" &mdash; Türkçe
func (client *Client) WotxEncyclopediaInfo(realm Realm, fields []string, language string) (*WotxEncyclopediaInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotxEncyclopediaInfo
	err := client.doGetDataRequest(realm, "/wotx/encyclopedia/info/", reqParam, &result)
	return result, err
}
