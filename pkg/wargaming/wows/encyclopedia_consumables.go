package wows

type EncyclopediaConsumables struct {
	// Consumable type
	Type_ string `json:"type,omitempty"`
	// Consumable characteristics
	Profile struct {
		// Characteristic value
		Value float32 `json:"value,omitempty"`
		// Characteristic description
		Description string `json:"description,omitempty"`
	} `json:"profile,omitempty"`
	// Cost in doubloons
	PriceGold int `json:"price_gold,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Consumable name
	Name string `json:"name,omitempty"`
	// Link to consumable image
	Image string `json:"image,omitempty"`
	// Consumable description
	Description string `json:"description,omitempty"`
	// Consumable ID
	ConsumableId int `json:"consumable_id,omitempty"`
} 
