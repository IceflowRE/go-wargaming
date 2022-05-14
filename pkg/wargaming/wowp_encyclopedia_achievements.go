package wargaming

import (
	"strings"
)

type WowpEncyclopediaAchievements struct {
	// Achievement ID
	AchievementId int `json:"achievement_id,omitempty"`
	// Achievement description
	Description string `json:"description,omitempty"`
	// URL to image
	Image string `json:"image,omitempty"`
	// Achievement name
	Name string `json:"name,omitempty"`
	// Localized name field
	NameI18N string `json:"name_i18n,omitempty"`
	// Sort order
	Order int `json:"order,omitempty"`
	// Section
	Section int `json:"section,omitempty"`
	// Localized section name
	SectionI18N string `json:"section_i18n,omitempty"`
	// ID of achievement type
	Type_ int `json:"type,omitempty"`
	// Localized type field
	TypeI18N string `json:"type_i18n,omitempty"`
}

// WowpEncyclopediaAchievements Method returns dictionary of achievements from Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/achievements
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
func (client *Client) WowpEncyclopediaAchievements(realm Realm, fields []string, language string) (*WowpEncyclopediaAchievements, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowpEncyclopediaAchievements
	err := client.doGetDataRequest(realm, "/wowp/encyclopedia/achievements/", reqParam, &result)
	return result, err
}
