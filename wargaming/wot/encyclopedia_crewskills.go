package wot

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
	// Сrew qualification ID
	Role *string
	// Skill ID. Maximum limit: 100.
	Skill []string
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
