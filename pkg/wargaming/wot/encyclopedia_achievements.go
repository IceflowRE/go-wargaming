package wot

type EncyclopediaAchievements struct {
	// Type
	Type_ string `json:"type,omitempty"`
	// Section order. Sections with a lesser value of the Section order field are displayed above sections with a greater value.
	SectionOrder int `json:"section_order,omitempty"`
	// Section
	Section string `json:"section,omitempty"`
	// Indicates, if achievement is outdated and cannot be received anymore
	Outdated bool `json:"outdated,omitempty"`
	// Achievement order in this section. Achievements with a lesser value of the Order field are displayed above achievements with a greater value.
	Order int `json:"order,omitempty"`
	// Service Record
	Options struct {
		// Information about nation emblems
		NationImages struct {
			// List of links to 95x85 px emblems
			X85 map[string]string `json:"x85,omitempty"`
			// List of links to 67x71 px emblems
			X71 map[string]string `json:"x71,omitempty"`
			// List of links to 180x180 px emblems
			X180 map[string]string `json:"x180,omitempty"`
		} `json:"nation_images,omitempty"`
		// Localized name field
		NameI18N string `json:"name_i18n,omitempty"`
		// 180x180px image
		ImageBig string `json:"image_big,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Achievement description
		Description string `json:"description,omitempty"`
	} `json:"options,omitempty"`
	// Localized name field
	NameI18N string `json:"name_i18n,omitempty"`
	// Achievement name
	Name string `json:"name,omitempty"`
	// 180x180px image
	ImageBig string `json:"image_big,omitempty"`
	// URL to image
	Image string `json:"image,omitempty"`
	// Historical reference
	HeroInfo string `json:"hero_info,omitempty"`
	// Achievement description
	Description string `json:"description,omitempty"`
	// Condition
	Condition string `json:"condition,omitempty"`
} 
