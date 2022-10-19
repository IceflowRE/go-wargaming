// Auto generated file!

package wotb

type EncyclopediaProvisionsOptions struct {
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
	// Equipment or consumable ID. Maximum limit: 100.
	ProvisionId []int `json:"provision_id,omitempty"`
	// Vehicle ID. Maximum limit: 100.
	TankId []int `json:"tank_id,omitempty"`
	// Type. Valid values:
	//
	// "equipment" - Consumables
	// "optionalDevice" - Equipment
	Type *string `json:"type,omitempty"`
}

type EncyclopediaProvisions struct {
	// Localized description
	DescriptionI18n *string `json:"description_i18n,omitempty"`
	// Localized name
	NameI18n *string `json:"name_i18n,omitempty"`
	// Cost in credits
	PriceCredit *int `json:"price_credit,omitempty"`
	// Purchase cost in gold
	PriceGold *int `json:"price_gold,omitempty"`
	// Equipment or consumable ID
	ProvisionId *int `json:"provision_id,omitempty"`
	// List of compatible vehicle IDs
	Tanks []int `json:"tanks,omitempty"`
	// Type
	Type *string `json:"type,omitempty"`
}
