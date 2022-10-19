// Auto generated file!

package wot

type EncyclopediaTanksOptions struct {
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
}

type EncyclopediaTanks struct {
	// URL to outline image of vehicle
	ContourImage *string `json:"contour_image,omitempty"`
	// URL to 160 x 100 px image of vehicle
	Image *string `json:"image,omitempty"`
	// URL to 124 x 31 px image of vehicle
	ImageSmall *string `json:"image_small,omitempty"`
	// Indicates if the vehicle is Premium vehicle
	IsPremium *bool `json:"is_premium,omitempty"`
	// Tier
	Level *int `json:"level,omitempty"`
	// Vehicle name
	Name *string `json:"name,omitempty"`
	// Localized name field
	NameI18n *string `json:"name_i18n,omitempty"`
	// Nation
	Nation *string `json:"nation,omitempty"`
	// Localized nation field
	NationI18n *string `json:"nation_i18n,omitempty"`
	// Localized short name of vehicle
	ShortNameI18n *string `json:"short_name_i18n,omitempty"`
	// Vehicle ID
	TankId *int `json:"tank_id,omitempty"`
	// Vehicle type
	Type *string `json:"type,omitempty"`
	// Localized vehicle type
	TypeI18n *string `json:"type_i18n,omitempty"`
}
