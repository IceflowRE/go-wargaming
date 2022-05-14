package wargaming

import (
	"strings"
)

type WowsEncyclopediaAchievements struct {
	// Battle achievements
	Battle struct {
		// Achievement ID
		AchievementId string `json:"achievement_id,omitempty"`
		// Battle types in which players can receive achievements. Battle types according to the Battle Types method response
		BattleTypes []string `json:"battle_types,omitempty"`
		// Indicates how many times achievement can be obtained per battle
		CountPerBattle string `json:"count_per_battle,omitempty"`
		// Achievement description
		Description string `json:"description,omitempty"`
		// Achievement unavailable
		Hidden int `json:"hidden,omitempty"`
		// Image link
		Image string `json:"image,omitempty"`
		// Image of an unearned achievement
		ImageInactive string `json:"image_inactive,omitempty"`
		// Indicates if achievement is in progress
		IsProgress int `json:"is_progress,omitempty"`
		// Maximum progress
		MaxProgress int `json:"max_progress,omitempty"`
		// Maximum tier of ship to receive this achievement
		MaxShipLevel int `json:"max_ship_level,omitempty"`
		// Minimum tier of ship to receive this achievement
		MinShipLevel int `json:"min_ship_level,omitempty"`
		// Achievement that can be received multiple times.
		Multiple int `json:"multiple,omitempty"`
		// Achievement name
		Name string `json:"name,omitempty"`
		// Indicates if a reward is granted for achievement
		Reward bool `json:"reward,omitempty"`
		// Subtype of achievement
		SubType string `json:"sub_type,omitempty"`
		// Type of achievement
		Type_ string `json:"type,omitempty"`
	} `json:"battle,omitempty"`
}

// WowsEncyclopediaAchievements Method returns information about achievements.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/achievements
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
func (client *Client) WowsEncyclopediaAchievements(realm Realm, fields []string, language string) (*WowsEncyclopediaAchievements, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowsEncyclopediaAchievements
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/achievements/", reqParam, &result)
	return result, err
}
