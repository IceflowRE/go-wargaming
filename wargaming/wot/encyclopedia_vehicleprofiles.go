package wot

// EncyclopediaVehicleprofilesOptions options.
type EncyclopediaVehicleprofilesOptions struct {
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
	// Sorting. Valid values:
	//
	// "price_credit" - by cost in credits
	// "-price_credit" - by cost in credits, in reverse order
	OrderBy *string
}

type EncyclopediaVehicleprofiles struct {
	// Standard configuration
	IsDefault *bool `json:"is_default,omitempty"`
	// Cost in credits
	PriceCredit *int `json:"price_credit,omitempty"`
	// Vehicle Configuration ID
	ProfileId *string `json:"profile_id,omitempty"`
	// Vehicle ID
	TankId *int `json:"tank_id,omitempty"`
}
