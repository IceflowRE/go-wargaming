package wows

type EncyclopediaCrewskills struct {
	// Skill type name
	TypeName string `json:"type_name,omitempty"`
	// Skill type ID
	TypeId int `json:"type_id,omitempty"`
	// Tier
	Tier int `json:"tier,omitempty"`
	// Skills
	Perks struct {
		// Skill ID
		PerkId int `json:"perk_id,omitempty"`
		// Description
		Description string `json:"description,omitempty"`
	} `json:"perks,omitempty"`
	// Name
	Name string `json:"name,omitempty"`
	// URL to skill icon
	Icon string `json:"icon,omitempty"`
} 
