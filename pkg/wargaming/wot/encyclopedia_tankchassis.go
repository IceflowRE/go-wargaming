package wot

type EncyclopediaTankchassis struct {
	// IDs of compatible vehicles
	Tanks []int `json:"tanks,omitempty"`
	// Traverse speed
	RotationSpeed int `json:"rotation_speed,omitempty"`
	// Purchase cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Purchase cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Localized nation field
	NationI18N string `json:"nation_i18n,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Localized name field
	NameI18N string `json:"name_i18n,omitempty"`
	// Vehicle name
	Name string `json:"name,omitempty"`
	// Module ID
	ModuleId int `json:"module_id,omitempty"`
	// Maximum load capacity
	MaxLoad float32 `json:"max_load,omitempty"`
	// Tier
	Level int `json:"level,omitempty"`
} 
