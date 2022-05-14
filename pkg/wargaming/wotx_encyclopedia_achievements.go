package wargaming

import (
	"strings"
)

type WotxEncyclopediaAchievements struct {
	// Award category
	Category string `json:"category,omitempty"`
	// Condition
	Condition string `json:"condition,omitempty"`
	// Award description
	Description string `json:"description,omitempty"`
	// Historical reference
	HeroInfo string `json:"hero_info,omitempty"`
	// Award image link
	Image string `json:"image,omitempty"`
	// Award name
	Name string `json:"name,omitempty"`
	// Award section
	Section string `json:"section,omitempty"`
	// Award type
	Type_ string `json:"type,omitempty"`
	// Award priority value (used to determine place of award in award list)
	Weight int `json:"weight,omitempty"`
	// Award classes (for mastery badges)
	Options struct {
		// Award image link
		Image string `json:"image,omitempty"`
		// Award name
		Name string `json:"name,omitempty"`
	} `json:"options,omitempty"`
}

// WotxEncyclopediaAchievements Method returns list of awards, medals, ribbons.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/achievements
//
// category:
//     Filter by award category. Maximum limit: 100. Valid values:
//     
//     "achievements" &mdash; Achievements 
//     "ribbons" &mdash; Ribbons
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
func (client *Client) WotxEncyclopediaAchievements(realm Realm, category []string, fields []string, language string) (*WotxEncyclopediaAchievements, error) {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"category": strings.Join(category, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotxEncyclopediaAchievements
	err := client.doGetDataRequest(realm, "/wotx/encyclopedia/achievements/", reqParam, &result)
	return result, err
}
