package wotx

type EncyclopediaAchievements struct {
	// Award priority value (used to determine place of award in award list)
	Weight int `json:"weight,omitempty"`
	// Award type
	Type_ string `json:"type,omitempty"`
	// Award section
	Section string `json:"section,omitempty"`
	// Award classes (for mastery badges)
	Options struct {
		// Award name
		Name string `json:"name,omitempty"`
		// Award image link
		Image string `json:"image,omitempty"`
	} `json:"options,omitempty"`
	// Award name
	Name string `json:"name,omitempty"`
	// Award image link
	Image string `json:"image,omitempty"`
	// Historical reference
	HeroInfo string `json:"hero_info,omitempty"`
	// Award description
	Description string `json:"description,omitempty"`
	// Condition
	Condition string `json:"condition,omitempty"`
	// Award category
	Category string `json:"category,omitempty"`
} 
