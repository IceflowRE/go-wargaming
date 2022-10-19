// Auto generated file!

package wot

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
	// Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
	Limit *int `json:"limit,omitempty"`
	// Result page number
	PageNo *int `json:"page_no,omitempty"`
	// Equipment or consumables ID. Maximum limit: 100.
	ProvisionId []int `json:"provision_id,omitempty"`
	// Type. Default is "equipment, optionalDevice". Maximum limit: 100. Valid values:
	//
	// "equipment" - Consumables
	// "optionalDevice" - Equipment
	Type []string `json:"type,omitempty"`
}

type EncyclopediaProvisions struct {
	// Achievement description
	Description *string `json:"description,omitempty"`
	// Image link
	Image *string `json:"image,omitempty"`
	// Vehicle name
	Name *string `json:"name,omitempty"`
	// Cost in credits
	PriceCredit *int `json:"price_credit,omitempty"`
	// Cost in gold
	PriceGold *int `json:"price_gold,omitempty"`
	// Equipment or consumables ID
	ProvisionId *int `json:"provision_id,omitempty"`
	// Technical name
	Tag *string `json:"tag,omitempty"`
	// Type: consumable or equipment
	Type *string `json:"type,omitempty"`
	// Weight in kilos. Used for equipment only.
	Weight *int `json:"weight,omitempty"`
}
