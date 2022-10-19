// Auto generated file!

package wows

type EncyclopediaConsumablesOptions struct {
	// Consumable ID. Maximum limit: 100.
	ConsumableId []int `json:"consumable_id,omitempty"`
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
	// Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
	Limit *int `json:"limit,omitempty"`
	// Page limit. Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default). Default is 1.
	PageNo *int `json:"page_no,omitempty"`
	// Consumable type. Valid values:
	//
	// "Camouflage" - Camouflages
	// "Flags" - Flags
	// "Permoflage" - Permanent camouflages
	// "Modernization" - Upgrades
	// "Skin" - Ship camouflages
	Type *string `json:"type,omitempty"`
}

type EncyclopediaConsumables struct {
	// Consumable ID
	ConsumableId *int `json:"consumable_id,omitempty"`
	// Consumable description
	Description *string `json:"description,omitempty"`
	// Link to consumable image
	Image *string `json:"image,omitempty"`
	// Consumable name
	Name *string `json:"name,omitempty"`
	// Cost in credits
	PriceCredit *int `json:"price_credit,omitempty"`
	// Cost in doubloons
	PriceGold *int `json:"price_gold,omitempty"`
	// Consumable characteristics
	Profile *struct {
		// Characteristic description
		Description *string `json:"description,omitempty"`
		// Characteristic value
		Value *float32 `json:"value,omitempty"`
	} `json:"profile,omitempty"`
	// Consumable type
	Type *string `json:"type,omitempty"`
}
