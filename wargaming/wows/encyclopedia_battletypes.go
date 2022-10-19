// Auto generated file!

package wows

type EncyclopediaBattletypesOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "cs" - Čeština
	// "de" - Deutsch
	// "en" - English (by default)
	// "es" - Español
	// "fr" - Français
	// "ja" - 日本語
	// "pl" - Polski
	// "ru" - Русский
	// "th" - ไทย
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "zh-cn" - 中文
	// "pt-br" - Português do Brasil
	// "es-mx" - Español (México)
	Language *string `json:"language,omitempty"`
}

type EncyclopediaBattletypes struct {
	// Battle type description
	Description *string `json:"description,omitempty"`
	// URL to image
	Image *string `json:"image,omitempty"`
	// Battle type name
	Name *string `json:"name,omitempty"`
	// Battle type tag
	Tag *string `json:"tag,omitempty"`
}
