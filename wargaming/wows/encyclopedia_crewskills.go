package wows

// EncyclopediaCrewskillsOptions options.
type EncyclopediaCrewskillsOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Localization language. Default is "ru". Valid values:
	//
	// "cs" - Čeština
	// "de" - Deutsch
	// "en" - English
	// "es" - Español
	// "fr" - Français
	// "ja" - 日本語
	// "pl" - Polski
	// "ru" - Русский (by default)
	// "th" - ไทย
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "zh-cn" - 中文
	// "pt-br" - Português do Brasil
	// "es-mx" - Español (México)
	Language *string
	// Skill ID. Maximum limit: 100.
	SkillId []int
}

type EncyclopediaCrewskills struct {
	// URL to skill icon
	Icon *string `json:"icon,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	// Skills
	Perks *struct {
		// Description
		Description *string `json:"description,omitempty"`
		// Skill ID
		PerkId *int `json:"perk_id,omitempty"`
	} `json:"perks,omitempty"`
	// Tier
	Tier *int `json:"tier,omitempty"`
	// Skill type ID
	TypeId *int `json:"type_id,omitempty"`
	// Skill type name
	TypeName *string `json:"type_name,omitempty"`
}
