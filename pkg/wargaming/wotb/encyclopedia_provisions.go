package wotb

type EncyclopediaProvisions struct {
	// Type
	Type_ string `json:"type,omitempty"`
	// List of compatible vehicle IDs
	Tanks []int `json:"tanks,omitempty"`
	// Equipment or consumable ID
	ProvisionId int `json:"provision_id,omitempty"`
	// Purchase cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Localized name
	NameI18N string `json:"name_i18n,omitempty"`
	// Localized description
	DescriptionI18N string `json:"description_i18n,omitempty"`
} 
