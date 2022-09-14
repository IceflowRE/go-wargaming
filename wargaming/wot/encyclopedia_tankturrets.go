// Auto generated file!

package wot

// EncyclopediaTankturretsOptions options.
type EncyclopediaTankturretsOptions struct {
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
	// Module ID. Maximum limit: 1000.
	ModuleId []int
	// Nation. Maximum limit: 100.
	Nation []string
}

type EncyclopediaTankturrets struct {
	// Armor: sides
	ArmorBoard *int `json:"armor_board,omitempty"`
	// Armor: rear
	ArmorFedd *int `json:"armor_fedd,omitempty"`
	// Armor: front
	ArmorForehead *int `json:"armor_forehead,omitempty"`
	// View range
	CircularVisionRadius *int `json:"circular_vision_radius,omitempty"`
	// Tier
	Level *int `json:"level,omitempty"`
	// Module ID
	ModuleId *int `json:"module_id,omitempty"`
	// Vehicle name
	Name *string `json:"name,omitempty"`
	// Localized name field
	NameI18N *string `json:"name_i18n,omitempty"`
	// Nation
	Nation *string `json:"nation,omitempty"`
	// Localized nation field
	NationI18N *string `json:"nation_i18n,omitempty"`
	// Purchase cost in credits
	PriceCredit *int `json:"price_credit,omitempty"`
	// Purchase cost in gold
	PriceGold *int `json:"price_gold,omitempty"`
	// Traverse speed
	RotationSpeed *int `json:"rotation_speed,omitempty"`
	// IDs of compatible vehicles
	Tanks []int `json:"tanks,omitempty"`
}
