package wotx

// EncyclopediaAchievementsOptions options.
type EncyclopediaAchievementsOptions struct {
	// Filter by award category. Maximum limit: 100. Valid values:
	//
	// "achievements" - Achievements
	// "ribbons" - Ribbons
	Category []string
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "ru" - Русский
	// "pl" - Polski
	// "de" - Deutsch
	// "fr" - Français
	// "es" - Español
	// "tr" - Türkçe
	Language *string
}

type EncyclopediaAchievements struct {
	// Award category
	Category *string `json:"category,omitempty"`
	// Condition
	Condition *string `json:"condition,omitempty"`
	// Award description
	Description *string `json:"description,omitempty"`
	// Historical reference
	HeroInfo *string `json:"hero_info,omitempty"`
	// Award image link
	Image *string `json:"image,omitempty"`
	// Award name
	Name *string `json:"name,omitempty"`
	// Award classes (for mastery badges)
	Options *struct {
		// Award image link
		Image *string `json:"image,omitempty"`
		// Award name
		Name *string `json:"name,omitempty"`
	} `json:"options,omitempty"`
	// Award section
	Section *string `json:"section,omitempty"`
	// Award type
	Type_ *string `json:"type,omitempty"`
	// Award priority value (used to determine place of award in award list)
	Weight *int `json:"weight,omitempty"`
}
