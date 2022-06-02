package wotx


// EncyclopediaCrewrolesOptions options.
type EncyclopediaCrewrolesOptions struct {
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

type EncyclopediaCrewroles struct {
	// Crew member qualification
	Name *string `json:"name,omitempty"`
	// Crew skills and perks
	Skills *struct {
		// Description of skill or perk
		Description *string `json:"description,omitempty"`
		// Images of skills and perks
		Images *struct {
			// Link to large image
			Large *string `json:"large,omitempty"`
		} `json:"images,omitempty"`
		// Indicates whether a skill or a perk is dealt with
		IsPerk *bool `json:"is_perk,omitempty"`
		// Name of skill or perk
		Name *string `json:"name,omitempty"`
	} `json:"skills,omitempty"`
}
