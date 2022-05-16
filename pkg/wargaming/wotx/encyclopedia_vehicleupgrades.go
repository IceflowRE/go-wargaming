package wotx

type EncyclopediaVehicleupgrades struct {
	// List of compatible equipment
	Equipment struct {
		// Cost in gold
		PriceGold int `json:"price_gold,omitempty"`
		// Cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Vehicle name
		Name string `json:"name,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Equipment ID
		EquipmentId int `json:"equipment_id,omitempty"`
		// Achievement description
		Description string `json:"description,omitempty"`
	} `json:"equipment,omitempty"`
	// List of compatible consumables
	Consumables struct {
		// Cost in gold
		PriceGold int `json:"price_gold,omitempty"`
		// Cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Vehicle name
		Name string `json:"name,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Achievement description
		Description string `json:"description,omitempty"`
		// Consumable ID
		ConsumableId int `json:"consumable_id,omitempty"`
	} `json:"consumables,omitempty"`
} 
