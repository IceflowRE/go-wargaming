package wot

type EncyclopediaBadges struct {
	// Badge name
	Name string `json:"name,omitempty"`
	// Image links
	Images struct {
		// URL to 24x24 px badge image
		SmallIcon string `json:"small_icon,omitempty"`
		// URL to 48x48 px badge image
		MediumIcon string `json:"medium_icon,omitempty"`
		// URL to 80x80 px badge image
		BigIcon string `json:"big_icon,omitempty"`
	} `json:"images,omitempty"`
	// Badge description
	Description string `json:"description,omitempty"`
	// Badge ID
	BadgeId int `json:"badge_id,omitempty"`
} 
