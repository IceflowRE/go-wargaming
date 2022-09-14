// Auto generated file!

package wot

// EncyclopediaCrewrolesOptions options.
type EncyclopediaCrewrolesOptions struct {
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
	// Сrew qualification ID. Maximum limit: 100.
	Role []string
}

type EncyclopediaCrewroles struct {
	// Сrew qualification name
	Name *string `json:"name,omitempty"`
	// Сrew qualification ID
	Role *string `json:"role,omitempty"`
	// List of crew member qualifications
	Skills []string `json:"skills,omitempty"`
}
