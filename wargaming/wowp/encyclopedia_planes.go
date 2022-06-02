package wowp


// EncyclopediaPlanesOptions options.
type EncyclopediaPlanesOptions struct {
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
	// "tr" - Türkçe 
	// "cs" - Čeština 
	// "th" - ไทย 
	// "vi" - Tiếng Việt 
	// "ko" - 한국어
	Language *string
	// Nation. Maximum limit: 100. Valid values:
	// 
	// "ussr" - U.S.S.R. 
	// "germany" - Germany 
	// "usa" - U.S.A. 
	// "france" - France 
	// "uk" - U.K. 
	// "china" - China 
	// "japan" - Japan
	Nation []string
	// Type. Maximum limit: 100. Valid values:
	// 
	// "fighter" - Fighter 
	// "fighterHeavy" - Heavy Fighter 
	// "assault" - Attack Aircraft 
	// "navy" - Multirole Fighter 
	// "bomber" - Bomber
	Type_ []string
}

type EncyclopediaPlanes struct {
	// Aircraft images
	Images *struct {
		// URL to 408 x 252 px image of aircraft
		Large *string `json:"large,omitempty"`
		// URL to 102 x 63 px image of aircraft
		Medium *string `json:"medium,omitempty"`
		// URL to 51 x 31 px image of aircraft
		Small *string `json:"small,omitempty"`
	} `json:"images,omitempty"`
	// Indicates if the aircraft was a gift
	IsGift *bool `json:"is_gift,omitempty"`
	// Indicates if the aircraft is Premium aircraft
	IsPremium *bool `json:"is_premium,omitempty"`
	// Tier
	Level *int `json:"level,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	// Localized name field
	NameI18N *string `json:"name_i18n,omitempty"`
	// Nation
	Nation *string `json:"nation,omitempty"`
	// Localized nation field
	NationI18N *string `json:"nation_i18n,omitempty"`
	// Aircraft ID
	PlaneId *int `json:"plane_id,omitempty"`
	// Type
	Type_ *string `json:"type,omitempty"`
}
