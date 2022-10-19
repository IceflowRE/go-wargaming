// Auto generated file!

package wot

type EncyclopediaTankenginesOptions struct {
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
	// Module ID. Maximum limit: 1000.
	ModuleId []int `json:"module_id,omitempty"`
	// Nation. Maximum limit: 100.
	Nation []string `json:"nation,omitempty"`
}

type EncyclopediaTankengines struct {
	// Chance of fire on impact
	FireStartingChance *int `json:"fire_starting_chance,omitempty"`
	// Tier
	Level *int `json:"level,omitempty"`
	// Module ID
	ModuleId *int `json:"module_id,omitempty"`
	// Vehicle name
	Name *string `json:"name,omitempty"`
	// Localized name field
	NameI18n *string `json:"name_i18n,omitempty"`
	// Nation
	Nation *string `json:"nation,omitempty"`
	// Localized nation field
	NationI18n *string `json:"nation_i18n,omitempty"`
	// Power
	Power *int `json:"power,omitempty"`
	// Purchase cost in credits
	PriceCredit *int `json:"price_credit,omitempty"`
	// Purchase cost in gold
	PriceGold *int `json:"price_gold,omitempty"`
	// IDs of compatible vehicles
	Tanks []int `json:"tanks,omitempty"`
}
