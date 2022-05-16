package wotx

type EncyclopediaVehicles struct {
	// Vehicle type
	Type_ string `json:"type,omitempty"`
	// Tier
	Tier int `json:"tier,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Vehicle tag
	Tag string `json:"tag,omitempty"`
	// Vehicle short name
	ShortName string `json:"short_name,omitempty"`
	// List of research costs in form of pairs:
	// 
	// parent vehicle ID
	// cost of research in XP
	PricesXp map[string]string `json:"prices_xp,omitempty"`
	// Cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// List of vehicles available for research in form of pairs:
	// 
	// researched vehicle ID
	// cost of research in XP
	NextTanks map[string]string `json:"next_tanks,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Vehicle name
	Name string `json:"name,omitempty"`
	// Indicates if the vehicle is Premium vehicle
	IsPremium bool `json:"is_premium,omitempty"`
	// Image links
	Images struct {
		// URL to 160 x 100 px image
		BigIcon string `json:"big_icon,omitempty"`
	} `json:"images,omitempty"`
	// Era
	Era string `json:"era,omitempty"`
	// Vehicle description
	Description string `json:"description,omitempty"`
} 
