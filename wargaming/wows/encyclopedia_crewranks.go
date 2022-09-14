// Auto generated file!

package wows

// EncyclopediaCrewranksOptions options.
type EncyclopediaCrewranksOptions struct {
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
	// Nation
	Nation *string
}

type EncyclopediaCrewranks struct {
	// Experience
	Experience *int `json:"experience,omitempty"`
	// Rank name
	Name *string `json:"name,omitempty"`
	// Rank Name by subnation index
	Names map[string]string `json:"names,omitempty"`
	// Rank
	Rank *int `json:"rank,omitempty"`
}
