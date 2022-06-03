package wowp

// EncyclopediaAchievementsOptions options.
type EncyclopediaAchievementsOptions struct {
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
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string
}

type EncyclopediaAchievements struct {
	// Achievement ID
	AchievementId *int `json:"achievement_id,omitempty"`
	// Achievement description
	Description *string `json:"description,omitempty"`
	// URL to image
	Image *string `json:"image,omitempty"`
	// Achievement name
	Name *string `json:"name,omitempty"`
	// Localized name field
	NameI18N *string `json:"name_i18n,omitempty"`
	// Sort order
	Order *int `json:"order,omitempty"`
	// Section
	Section *int `json:"section,omitempty"`
	// Localized section name
	SectionI18N *string `json:"section_i18n,omitempty"`
	// Localized type field
	TypeI18N *string `json:"type_i18n,omitempty"`
	// ID of achievement type
	Type_ *int `json:"type,omitempty"`
}
