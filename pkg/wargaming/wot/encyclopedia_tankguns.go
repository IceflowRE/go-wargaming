package wot

type EncyclopediaTankguns struct {
	// IDs of compatible turrets
	Turrets []int `json:"turrets,omitempty"`
	// IDs of compatible vehicles
	Tanks []int `json:"tanks,omitempty"`
	// Rate of fire
	Rate float32 `json:"rate,omitempty"`
	// Purchase cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Purchase cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Penetration
	PiercingPower []int `json:"piercing_power,omitempty"`
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
	// Tier
	Level int `json:"level,omitempty"`
	// Damage
	Damage []int `json:"damage,omitempty"`
} 
