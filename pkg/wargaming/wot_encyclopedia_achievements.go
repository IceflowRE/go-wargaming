package wargaming

import (
	"strings"
)

type WotEncyclopediaAchievements struct {
	// Condition
	Condition string `json:"condition,omitempty"`
	// Achievement description
	Description string `json:"description,omitempty"`
	// Historical reference
	HeroInfo string `json:"hero_info,omitempty"`
	// URL to image
	Image string `json:"image,omitempty"`
	// 180x180px image
	ImageBig string `json:"image_big,omitempty"`
	// Achievement name
	Name string `json:"name,omitempty"`
	// Localized name field
	NameI18N string `json:"name_i18n,omitempty"`
	// Achievement order in this section. Achievements with a lesser value of the Order field are displayed above achievements with a greater value.
	Order int `json:"order,omitempty"`
	// Indicates, if achievement is outdated and cannot be received anymore
	Outdated bool `json:"outdated,omitempty"`
	// Section
	Section string `json:"section,omitempty"`
	// Section order. Sections with a lesser value of the Section order field are displayed above sections with a greater value.
	SectionOrder int `json:"section_order,omitempty"`
	// Type
	Type_ string `json:"type,omitempty"`
	// Service Record
	Options struct {
		// Achievement description
		Description string `json:"description,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// 180x180px image
		ImageBig string `json:"image_big,omitempty"`
		// Localized name field
		NameI18N string `json:"name_i18n,omitempty"`
		// Information about nation emblems
		NationImages struct {
			// List of links to 180x180 px emblems
			X180 map[string]string `json:"x180,omitempty"`
			// List of links to 67x71 px emblems
			X71 map[string]string `json:"x71,omitempty"`
			// List of links to 95x85 px emblems
			X85 map[string]string `json:"x85,omitempty"`
		} `json:"nation_images,omitempty"`
	} `json:"options,omitempty"`
}

// WotEncyclopediaAchievements Method returns information about achievements.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/achievements
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
func (client *Client) WotEncyclopediaAchievements(realm Realm, fields []string, language string) (*WotEncyclopediaAchievements, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotEncyclopediaAchievements
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/achievements/", reqParam, &result)
	return result, err
}
