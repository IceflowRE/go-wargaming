// Auto generated file!

package wot

type EncyclopediaBadgesOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "ru" - Русский
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
	Language *string `json:"language,omitempty"`
}

type EncyclopediaBadges struct {
	// Badge ID
	BadgeId *int `json:"badge_id,omitempty"`
	// Badge description
	Description *string `json:"description,omitempty"`
	// Image links
	Images *struct {
		// URL to 80x80 px badge image
		BigIcon *string `json:"big_icon,omitempty"`
		// URL to 48x48 px badge image
		MediumIcon *string `json:"medium_icon,omitempty"`
		// URL to 24x24 px badge image
		SmallIcon *string `json:"small_icon,omitempty"`
	} `json:"images,omitempty"`
	// Badge name
	Name *string `json:"name,omitempty"`
}
