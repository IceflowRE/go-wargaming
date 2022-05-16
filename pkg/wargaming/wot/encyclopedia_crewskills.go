package wot

type EncyclopediaCrewskills struct {
	// Skill ID
	Skill string `json:"skill,omitempty"`
	// Skill name
	Name string `json:"name,omitempty"`
	// Indicates if it is a perk
	IsPerk bool `json:"is_perk,omitempty"`
	// URL to skill icon
	ImageUrl struct {
		// URL to small image
		SmallIcon string `json:"small_icon,omitempty"`
		// URL to large image
		BigIcon string `json:"big_icon,omitempty"`
	} `json:"image_url,omitempty"`
	// Skill description
	Description string `json:"description,omitempty"`
} 
