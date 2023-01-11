// Auto generated file!

package wotb

type EncyclopediaCrewskillsOptions struct {
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
	// Skills IDs. Maximum limit: 100.
	SkillId []string `json:"skill_id,omitempty"`
	// Vehicle types. Maximum limit: 100.
	VehicleType []string `json:"vehicle_type,omitempty"`
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
