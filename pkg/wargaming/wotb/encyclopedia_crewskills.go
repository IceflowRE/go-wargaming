package wotb

type EncyclopediaCrewskills struct {
	// Vehicle type
	VehicleType string `json:"vehicle_type,omitempty"`
	// Tip
	Tip string `json:"tip,omitempty"`
	// Skill ID
	SkillId string `json:"skill_id,omitempty"`
	// Skill name
	Name string `json:"name,omitempty"`
	// Skill images
	Images struct {
		// URL to large image
		Large string `json:"large,omitempty"`
	} `json:"images,omitempty"`
	// Features
	Features string `json:"features,omitempty"`
	// Skill effect
	Effect string `json:"effect,omitempty"`
} 
