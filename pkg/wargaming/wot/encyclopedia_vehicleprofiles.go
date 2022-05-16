package wot

type EncyclopediaVehicleprofiles struct {
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Vehicle Configuration ID
	ProfileId string `json:"profile_id,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Standard configuration
	IsDefault bool `json:"is_default,omitempty"`
} 
