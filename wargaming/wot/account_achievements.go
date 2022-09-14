// Auto generated file!

package wot

// AccountAchievementsOptions options.
type AccountAchievementsOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Localization language. Default is "ru". Valid values:
	//
	// "en" - English
	// "ru" - Русский (by default)
	// "pl" - Polski
	// "de" - Deutsch
	// "fr" - Français
	// "es" - Español
	// "zh-cn" - 简体中文
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string
}

type AccountAchievements struct {
	// Achievements earned
	Achievements map[string]string `json:"achievements,omitempty"`
	// Achievement progress
	Frags map[string]string `json:"frags,omitempty"`
	// Maximum values of achievement series
	MaxSeries map[string]string `json:"max_series,omitempty"`
}
