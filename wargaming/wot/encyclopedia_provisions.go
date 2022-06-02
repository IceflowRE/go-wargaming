package wot


// EncyclopediaProvisionsOptions options.
type EncyclopediaProvisionsOptions struct {
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
	// Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
	Limit *int
	// Result page number
	PageNo *int
	// Equipment or consumables ID. Maximum limit: 100.
	ProvisionId []int
	// Type. Default is "equipment, optionalDevice". Maximum limit: 100. Valid values:
	// 
	// "equipment" - Consumables 
	// "optionalDevice" - Equipment
	Type_ []string
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
	Type_ *string `json:"type,omitempty"`
	// Weight in kilos. Used for equipment only.
	Weight *int `json:"weight,omitempty"`
}
