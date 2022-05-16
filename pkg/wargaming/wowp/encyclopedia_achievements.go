package wowp

type EncyclopediaAchievements struct {
	// ID of achievement type
	Type_ int `json:"type,omitempty"`
	// Localized type field
	TypeI18N string `json:"type_i18n,omitempty"`
	// Localized section name
	SectionI18N string `json:"section_i18n,omitempty"`
	// Section
	Section int `json:"section,omitempty"`
	// Sort order
	Order int `json:"order,omitempty"`
	// Localized name field
	NameI18N string `json:"name_i18n,omitempty"`
	// Achievement name
	Name string `json:"name,omitempty"`
	// URL to image
	Image string `json:"image,omitempty"`
	// Achievement description
	Description string `json:"description,omitempty"`
	// Achievement ID
	AchievementId int `json:"achievement_id,omitempty"`
} 
