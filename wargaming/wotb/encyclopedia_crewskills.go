package wotb


// EncyclopediaCrewskillsOptions options.
type EncyclopediaCrewskillsOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Localization language. Default is "ru". Valid values:
	// 
	// "en" - English 
	// "ru" - Русский (by default)
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
	Language *string
	// Skills IDs. Maximum limit: 100.
	SkillId []string
	// Vehicle types. Maximum limit: 100.
	VehicleType []string
}

type EncyclopediaCrewskills struct {
	// Skill effect
	Effect *string `json:"effect,omitempty"`
	// Features
	Features *string `json:"features,omitempty"`
	// Skill images
	Images *struct {
		// URL to large image
		Large *string `json:"large,omitempty"`
	} `json:"images,omitempty"`
	// Skill name
	Name *string `json:"name,omitempty"`
	// Skill ID
	SkillId *string `json:"skill_id,omitempty"`
	// Tip
	Tip *string `json:"tip,omitempty"`
	// Vehicle type
	VehicleType *string `json:"vehicle_type,omitempty"`
}
