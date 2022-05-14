package wargaming

import (
	"strings"
)

type WotEncyclopediaInfo struct {
	// Game client version
	GameVersion string `json:"game_version,omitempty"`
	// List of supported languages
	Languages map[string]string `json:"languages,omitempty"`
	// Time when vehicles details in Tankopedia were updated
	TanksUpdatedAt UnixTime `json:"tanks_updated_at,omitempty"`
	// Available crew qualifications
	VehicleCrewRoles map[string]string `json:"vehicle_crew_roles,omitempty"`
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

// WotEncyclopediaInfo Method returns information about Tankopedia.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/info
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
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
func (client *Client) WotEncyclopediaInfo(realm Realm, fields []string, language string) (*WotEncyclopediaInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotEncyclopediaInfo
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/info/", reqParam, &result)
	return result, err
}
