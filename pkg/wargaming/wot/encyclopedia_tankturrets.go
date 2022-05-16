package wot

type EncyclopediaTankturrets struct {
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
	// Tier
	Level int `json:"level,omitempty"`
	// View range
	CircularVisionRadius int `json:"circular_vision_radius,omitempty"`
	// Armor: front
	ArmorForehead int `json:"armor_forehead,omitempty"`
	// Armor: rear
	ArmorFedd int `json:"armor_fedd,omitempty"`
	// Armor: sides
	ArmorBoard int `json:"armor_board,omitempty"`
} 
