// Auto generated file!

package wotx

type EncyclopediaVehicleupgradesOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "ru" - Русский
	// "pl" - Polski
	// "de" - Deutsch
	// "fr" - Français
	// "es" - Español
	// "tr" - Türkçe
	Language *string `json:"language,omitempty"`
}

type EncyclopediaVehicleupgrades struct {
	// List of compatible consumables
	Consumables *struct {
		// Consumable ID
		ConsumableId *int `json:"consumable_id,omitempty"`
		// Achievement description
		Description *string `json:"description,omitempty"`
		// URL to image
		Image *string `json:"image,omitempty"`
		// Vehicle name
		Name *string `json:"name,omitempty"`
		// Cost in credits
		PriceCredit *int `json:"price_credit,omitempty"`
		// Cost in gold
		PriceGold *int `json:"price_gold,omitempty"`
	} `json:"consumables,omitempty"`
	// List of compatible equipment
	Equipment *struct {
		// Achievement description
		Description *string `json:"description,omitempty"`
		// Equipment ID
		EquipmentId *int `json:"equipment_id,omitempty"`
		// URL to image
		Image *string `json:"image,omitempty"`
		// Vehicle name
		Name *string `json:"name,omitempty"`
		// Cost in credits
		PriceCredit *int `json:"price_credit,omitempty"`
		// Cost in gold
		PriceGold *int `json:"price_gold,omitempty"`
	} `json:"equipment,omitempty"`
}
