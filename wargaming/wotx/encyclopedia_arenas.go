// Auto generated file!

package wotx

type EncyclopediaArenasOptions struct {
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
	// "tr" - Türkçe
	Language *string `json:"language,omitempty"`
}

type EncyclopediaArenas struct {
	// Map ID
	ArenaId *string `json:"arena_id,omitempty"`
	// Map type
	CamouflageType *string `json:"camouflage_type,omitempty"`
	// Short map description
	Description *string `json:"description,omitempty"`
	// Localized map name
	NameI18n *string `json:"name_i18n,omitempty"`
}
