// Auto generated file!

package wotb

type EncyclopediaAchievementsOptions struct {
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

type EncyclopediaAchievements struct {
	// Achievement ID
	AchievementId *string `json:"achievement_id,omitempty"`
	// Condition
	Condition *string `json:"condition,omitempty"`
	// Achievement description
	Description *string `json:"description,omitempty"`
	// Image link
	Image *string `json:"image,omitempty"`
	// Link to large image
	ImageBig *string `json:"image_big,omitempty"`
	// Achievement name
	Name *string `json:"name,omitempty"`
	// Service Record
	Options *struct {
		// Achievement description
		Description *string `json:"description,omitempty"`
		// Image link
		Image *string `json:"image,omitempty"`
		// Link to large image
		ImageBig *string `json:"image_big,omitempty"`
		// Achievement name
		Name *string `json:"name,omitempty"`
	} `json:"options,omitempty"`
	// Order of achievements
	Order *int `json:"order,omitempty"`
	// Section
	Section *string `json:"section,omitempty"`
}
