// Auto generated file!

package wotx

type EncyclopediaVehiclesOptions struct {
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
	// Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
	Limit *int `json:"limit,omitempty"`
	// Nation. Maximum limit: 100.
	Nation []string `json:"nation,omitempty"`
	// Result page number
	PageNo *int `json:"page_no,omitempty"`
	// Vehicle ID. Maximum limit: 100.
	TankId []int `json:"tank_id,omitempty"`
	// Tier. Maximum limit: 100.
	Tier []int `json:"tier,omitempty"`
}

type EncyclopediaVehicles struct {
	// Vehicle description
	Description *string `json:"description,omitempty"`
	// Era
	Era *string `json:"era,omitempty"`
	// Vehicle has multiple weapons
	HasMultiWeapon *bool `json:"has_multi_weapon,omitempty"`
	// Image links
	Images *struct {
		// URL to 160 x 100 px image of vehicle
		BigIcon *string `json:"big_icon,omitempty"`
	} `json:"images,omitempty"`
	// Indicates if the vehicle is Premium vehicle
	IsPremium *bool `json:"is_premium,omitempty"`
	// Module research information
	ModulesTree *struct {
		// Indicates if the module is basic
		IsDefault *bool `json:"is_default,omitempty"`
		// Module ID
		ModuleId *int `json:"module_id,omitempty"`
		// Module name
		Name *string `json:"name,omitempty"`
		// List of module IDs available after research of the module
		NextModules []int `json:"next_modules,omitempty"`
		// List of vehicle IDs available after research of the module
		NextTanks []int `json:"next_tanks,omitempty"`
		// Cost in credits
		PriceCredit *int `json:"price_credit,omitempty"`
		// Research cost
		PriceXp *int `json:"price_xp,omitempty"`
		// Module type
		Type *string `json:"type,omitempty"`
	} `json:"modules_tree,omitempty"`
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
	// Number of turrets
	TotalTurrets *int `json:"total_turrets,omitempty"`
	// Vehicle type
	Type *string `json:"type,omitempty"`
}
