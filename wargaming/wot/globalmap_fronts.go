// Auto generated file!

package wot

type GlobalmapFrontsOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// List of Front IDs, to specify what fronts need to be returned. Maximum limit: 100.
	FrontId []string `json:"front_id,omitempty"`
	// Language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "de" - German
	// "fr" - French
	// "es" - Spanish
	// "pl" - Polish
	// "tr" - Turkish
	Language *string `json:"language,omitempty"`
	// Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
	Limit *int `json:"limit,omitempty"`
	// Page number. Default is 1. Min value is 1.
	PageNo *int `json:"page_no,omitempty"`
}

type GlobalmapFronts struct {
	// Available modules
	AvailableExtensions []*struct {
		// Cost of modules
		Cost *int `json:"cost,omitempty"`
		// Localized description of strategic effect
		DescriptionStrategic *string `json:"description_strategic,omitempty"`
		// Localized description of tactical effect
		DescriptionTactical *string `json:"description_tactical,omitempty"`
		// Consumable ID
		ExtensionId *string `json:"extension_id,omitempty"`
		// Localized consumable name
		Name *string `json:"name,omitempty"`
		// Modules upkeep cost
		Wage *int `json:"wage,omitempty"`
	} `json:"available_extensions,omitempty"`
	// Average clans rating
	AvgClansRating *int `json:"avg_clans_rating,omitempty"`
	// Average minimum bid
	AvgMinBet *int `json:"avg_min_bet,omitempty"`
	// Average winning bid
	AvgWonBet *int `json:"avg_won_bet,omitempty"`
	// Maximum battle duration in minutes
	BattleTimeLimit *int `json:"battle_time_limit,omitempty"`
	// Division cost
	DivisionCost *int `json:"division_cost,omitempty"`
	// Indicates presence of Fog of War
	FogOfWar *bool `json:"fog_of_war,omitempty"`
	// Front ID
	FrontId *string `json:"front_id,omitempty"`
	// Front name
	FrontName *string `json:"front_name,omitempty"`
	// Indicates if a map is active
	IsActive *bool `json:"is_active,omitempty"`
	// Indicates the map type: regular map or events map
	IsEvent *bool `json:"is_event,omitempty"`
	// Maximum number of vehicles in division
	MaxTanksPerDivision *int `json:"max_tanks_per_division,omitempty"`
	// Maximum vehicle Tier
	MaxVehicleLevel *int `json:"max_vehicle_level,omitempty"`
	// Minimum number of vehicles in division
	MinTanksPerDivision *int `json:"min_tanks_per_division,omitempty"`
	// Minimum vehicle Tier
	MinVehicleLevel *int `json:"min_vehicle_level,omitempty"`
	// Amount of Provinces
	ProvincesCount *int `json:"provinces_count,omitempty"`
	// Indicates if vehicles blocking is active
	VehicleFreeze *bool `json:"vehicle_freeze,omitempty"`
}
