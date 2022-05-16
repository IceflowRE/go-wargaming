package wot

type EncyclopediaProvisions struct {
	// Weight in kilos. Used for equipment only.
	Weight int `json:"weight,omitempty"`
	// Type: consumable or equipment
	Type_ string `json:"type,omitempty"`
	// Technical name
	Tag string `json:"tag,omitempty"`
	// Equipment or consumables ID
	ProvisionId int `json:"provision_id,omitempty"`
	// Cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Vehicle name
	Name string `json:"name,omitempty"`
	// Image link
	Image string `json:"image,omitempty"`
	// Achievement description
	Description string `json:"description,omitempty"`
} 
