// Auto generated file!

package wot

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
	// Condition
	Condition *string `json:"condition,omitempty"`
	// Achievement description
	Description *string `json:"description,omitempty"`
	// Historical reference
	HeroInfo *string `json:"hero_info,omitempty"`
	// URL to image
	Image *string `json:"image,omitempty"`
	// 180x180px image
	ImageBig *string `json:"image_big,omitempty"`
	// Achievement name
	Name *string `json:"name,omitempty"`
	// Localized name field
	NameI18n *string `json:"name_i18n,omitempty"`
	// Service Record
	Options *struct {
		// Achievement description
		Description *string `json:"description,omitempty"`
		// URL to image
		Image *string `json:"image,omitempty"`
		// 180x180px image
		ImageBig *string `json:"image_big,omitempty"`
		// Localized name field
		NameI18n *string `json:"name_i18n,omitempty"`
		// Information about nation emblems
		NationImages *struct {
			// List of links to 180x180 px emblems
			X180 map[string]string `json:"x180,omitempty"`
			// List of links to 67x71 px emblems
			X71 map[string]string `json:"x71,omitempty"`
			// List of links to 95x85 px emblems
			X85 map[string]string `json:"x85,omitempty"`
		} `json:"nation_images,omitempty"`
	} `json:"options,omitempty"`
	// Achievement order in this section. Achievements with a lesser value of the Order field are displayed above achievements with a greater value.
	Order *int `json:"order,omitempty"`
	// Indicates, if achievement is outdated and cannot be received anymore
	Outdated *bool `json:"outdated,omitempty"`
	// Section
	Section *string `json:"section,omitempty"`
	// Section order. Sections with a lesser value of the Section order field are displayed above sections with a greater value.
	SectionOrder *int `json:"section_order,omitempty"`
	// Type
	Type *string `json:"type,omitempty"`
}
