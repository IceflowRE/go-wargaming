// Auto generated file!

package wows

type EncyclopediaCrewskillsOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "cs" - Čeština
	// "de" - Deutsch
	// "en" - English (by default)
	// "es" - Español
	// "fr" - Français
	// "it" - Italiano
	// "ja" - 日本語
	// "pl" - Polski
	// "ru" - Русский
	// "th" - ไทย
	// "zh-tw" - 繁體中文
	// "zh-cn" - 中文
	// "tr" - Türkçe
	// "pt-br" - Português do Brasil
	// "es-mx" - Español (México)
	Language *string `json:"language,omitempty"`
	// Skill ID. Maximum limit: 100.
	SkillId []int `json:"skill_id,omitempty"`
}

type EncyclopediaCrewskills struct {
	// Ship Type Customization
	Customization *struct {
		// Column
		Column *int `json:"column,omitempty"`
		// Skills
		Perks *struct {
			// Description
			Description *string `json:"description,omitempty"`
			// Skill ID
			PerkId *int `json:"perk_id,omitempty"`
		} `json:"perks,omitempty"`
		// Tier
		Tier *int `json:"tier,omitempty"`
	} `json:"customization,omitempty"`
	// URL to skill icon
	Icon *string `json:"icon,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	// Skill type ID
	TypeId *int `json:"type_id,omitempty"`
	// Skill type name
	TypeName *string `json:"type_name,omitempty"`
}
