package wotx


// EncyclopediaVehiclesOptions options.
type EncyclopediaVehiclesOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Localization language. Default is "en". Valid values:
	// 
	// "en" - English (by default)
	// "ru" - Русский 
	// "pl" - Polski 
	// "de" - Deutsch 
	// "fr" - Français 
	// "es" - Español 
	// "tr" - Türkçe
	Language *string
	// Nation. Maximum limit: 100.
	Nation []string
	// Vehicle ID. Maximum limit: 100.
	TankId []int
	// Tier. Maximum limit: 100.
	Tier []int
}

type EncyclopediaVehicles struct {
	// Vehicle description
	Description *string `json:"description,omitempty"`
	// Era
	Era *string `json:"era,omitempty"`
	// Image links
	Images *struct {
		// URL to 160 x 100 px image
		BigIcon *string `json:"big_icon,omitempty"`
	} `json:"images,omitempty"`
	// Indicates if the vehicle is Premium vehicle
	IsPremium *bool `json:"is_premium,omitempty"`
	// Vehicle name
	Name *string `json:"name,omitempty"`
	// Nation
	Nation *string `json:"nation,omitempty"`
	// List of vehicles available for research in form of pairs:
	// 
	// researched vehicle ID
	// cost of research in XP
	NextTanks map[string]string `json:"next_tanks,omitempty"`
	// Cost in credits
	PriceCredit *int `json:"price_credit,omitempty"`
	// Cost in gold
	PriceGold *int `json:"price_gold,omitempty"`
	// List of research costs in form of pairs:
	// 
	// parent vehicle ID
	// cost of research in XP
	PricesXp map[string]string `json:"prices_xp,omitempty"`
	// Vehicle short name
	ShortName *string `json:"short_name,omitempty"`
	// Vehicle tag
	Tag *string `json:"tag,omitempty"`
	// Vehicle ID
	TankId *int `json:"tank_id,omitempty"`
	// Tier
	Tier *int `json:"tier,omitempty"`
	// Vehicle type
	Type_ *string `json:"type,omitempty"`
}
