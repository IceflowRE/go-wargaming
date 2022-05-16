package wot

type EncyclopediaTanks struct {
	// Vehicle type
	Type_ string `json:"type,omitempty"`
	// Localized vehicle type
	TypeI18N string `json:"type_i18n,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Localized short name of vehicle
	ShortNameI18N string `json:"short_name_i18n,omitempty"`
	// Localized nation field
	NationI18N string `json:"nation_i18n,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Localized name field
	NameI18N string `json:"name_i18n,omitempty"`
	// Vehicle name
	Name string `json:"name,omitempty"`
	// Tier
	Level int `json:"level,omitempty"`
	// Indicates if the vehicle is Premium vehicle
	IsPremium bool `json:"is_premium,omitempty"`
	// URL to 124 x 31 px image of vehicle
	ImageSmall string `json:"image_small,omitempty"`
	// URL to 160 x 100 px image of vehicle
	Image string `json:"image,omitempty"`
	// URL to outline image of vehicle
	ContourImage string `json:"contour_image,omitempty"`
} 
