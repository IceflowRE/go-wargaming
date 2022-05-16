package wot

type EncyclopediaTankengines struct {
	// IDs of compatible vehicles
	Tanks []int `json:"tanks,omitempty"`
	// Purchase cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Purchase cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Power
	Power int `json:"power,omitempty"`
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
	// Chance of fire on impact
	FireStartingChance int `json:"fire_starting_chance,omitempty"`
} 
