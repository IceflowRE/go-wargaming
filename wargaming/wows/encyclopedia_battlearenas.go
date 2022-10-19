// Auto generated file!

package wows

type EncyclopediaBattlearenasOptions struct {
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

type EncyclopediaBattlearenas struct {
	// Map ID
	BattleArenaId *int `json:"battle_arena_id,omitempty"`
	// Map description
	Description *string `json:"description,omitempty"`
	// URL to the map icon
	Icon *string `json:"icon,omitempty"`
	// Map name
	Name *string `json:"name,omitempty"`
}
