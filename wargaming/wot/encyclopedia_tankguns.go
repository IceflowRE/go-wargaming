// Auto generated file!

package wot

type EncyclopediaTankgunsOptions struct {
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
	// Compatible vehicle ID
	TankId *int `json:"tank_id,omitempty"`
	// Compatible turret ID
	TurretId *int `json:"turret_id,omitempty"`
}

type EncyclopediaTankguns struct {
	// Damage
	Damage []int `json:"damage,omitempty"`
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
	// Penetration
	PiercingPower []int `json:"piercing_power,omitempty"`
	// Purchase cost in credits
	PriceCredit *int `json:"price_credit,omitempty"`
	// Purchase cost in gold
	PriceGold *int `json:"price_gold,omitempty"`
	// Rate of fire
	Rate *float32 `json:"rate,omitempty"`
	// IDs of compatible vehicles
	Tanks []int `json:"tanks,omitempty"`
	// IDs of compatible turrets
	Turrets []int `json:"turrets,omitempty"`
}
