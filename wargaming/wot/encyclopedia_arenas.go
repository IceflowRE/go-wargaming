// Auto generated file!

package wot

// EncyclopediaArenasOptions options.
type EncyclopediaArenasOptions struct {
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
}

type EncyclopediaArenas struct {
	// Map ID
	ArenaId *string `json:"arena_id,omitempty"`
	// Map type
	CamouflageType *string `json:"camouflage_type,omitempty"`
	// Short map description
	Description *string `json:"description,omitempty"`
	// Localized map name
	NameI18N *string `json:"name_i18n,omitempty"`
}
