package wargaming

import (
	"strings"
)

type WotEncyclopediaBadges struct {
	// Badge ID
	BadgeId int `json:"badge_id,omitempty"`
	// Badge description
	Description string `json:"description,omitempty"`
	// Badge name
	Name string `json:"name,omitempty"`
	// Image links
	Images struct {
		// URL to 80x80 px badge image
		BigIcon string `json:"big_icon,omitempty"`
		// URL to 48x48 px badge image
		MediumIcon string `json:"medium_icon,omitempty"`
		// URL to 24x24 px badge image
		SmallIcon string `json:"small_icon,omitempty"`
	} `json:"images,omitempty"`
}

// WotEncyclopediaBadges Method returns list of available badges a player can gain in Ranked Battles.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/badges
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
func (client *Client) WotEncyclopediaBadges(realm Realm, fields []string, language string) (*WotEncyclopediaBadges, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotEncyclopediaBadges
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/badges/", reqParam, &result)
	return result, err
}
