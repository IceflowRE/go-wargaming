package wotx

type EncyclopediaCrewroles struct {
	// Crew skills and perks
	Skills struct {
		// Name of skill or perk
		Name string `json:"name,omitempty"`
		// Indicates whether a skill or a perk is dealt with
		IsPerk bool `json:"is_perk,omitempty"`
		// Images of skills and perks
		Images struct {
			// Link to large image
			Large string `json:"large,omitempty"`
		} `json:"images,omitempty"`
		// Description of skill or perk
		Description string `json:"description,omitempty"`
	} `json:"skills,omitempty"`
	// Crew member qualification
	Name string `json:"name,omitempty"`
} 
