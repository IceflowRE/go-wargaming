package wowp

type EncyclopediaPlanes struct {
	// Type
	Type_ string `json:"type,omitempty"`
	// Aircraft ID
	PlaneId int `json:"plane_id,omitempty"`
	// Localized nation field
	NationI18N string `json:"nation_i18n,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Localized name field
	NameI18N string `json:"name_i18n,omitempty"`
	// Name
	Name string `json:"name,omitempty"`
	// Tier
	Level int `json:"level,omitempty"`
	// Indicates if the aircraft is Premium aircraft
	IsPremium bool `json:"is_premium,omitempty"`
	// Indicates if the aircraft was a gift
	IsGift bool `json:"is_gift,omitempty"`
	// Aircraft images
	Images struct {
		// URL to 51 x 31 px image of aircraft
		Small string `json:"small,omitempty"`
		// URL to 102 x 63 px image of aircraft
		Medium string `json:"medium,omitempty"`
		// URL to 408 x 252 px image of aircraft
		Large string `json:"large,omitempty"`
	} `json:"images,omitempty"`
} 
