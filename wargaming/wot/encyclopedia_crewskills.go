// Auto generated file!

package wot

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
	// Сrew qualification ID
	Role *string `json:"role,omitempty"`
	// Skill ID. Maximum limit: 100.
	Skill []string `json:"skill,omitempty"`
}

type EncyclopediaCrewskills struct {
	// Skill description
	Description *string `json:"description,omitempty"`
	// URL to skill icon
	ImageUrl *struct {
		// URL to large image
		BigIcon *string `json:"big_icon,omitempty"`
		// URL to small image
		SmallIcon *string `json:"small_icon,omitempty"`
	} `json:"image_url,omitempty"`
	// Indicates if it is a perk
	IsPerk *bool `json:"is_perk,omitempty"`
	// Skill name
	Name *string `json:"name,omitempty"`
	// Skill ID
	Skill *string `json:"skill,omitempty"`
}
