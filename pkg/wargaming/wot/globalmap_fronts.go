package wot

type GlobalmapFronts struct {
	// Indicates if vehicles blocking is active
	VehicleFreeze bool `json:"vehicle_freeze,omitempty"`
	// Amount of Provinces
	ProvincesCount int `json:"provinces_count,omitempty"`
	// Minimum vehicle Tier
	MinVehicleLevel int `json:"min_vehicle_level,omitempty"`
	// Minimum number of vehicles in division
	MinTanksPerDivision int `json:"min_tanks_per_division,omitempty"`
	// Maximum vehicle Tier
	MaxVehicleLevel int `json:"max_vehicle_level,omitempty"`
	// Maximum number of vehicles in division
	MaxTanksPerDivision int `json:"max_tanks_per_division,omitempty"`
	// Indicates the map type: regular map or events map
	IsEvent bool `json:"is_event,omitempty"`
	// Indicates if a map is active
	IsActive bool `json:"is_active,omitempty"`
	// Front name
	FrontName string `json:"front_name,omitempty"`
	// Front ID
	FrontId string `json:"front_id,omitempty"`
	// Indicates presence of Fog of War
	FogOfWar bool `json:"fog_of_war,omitempty"`
	// Division cost
	DivisionCost int `json:"division_cost,omitempty"`
	// Maximum battle duration in minutes
	BattleTimeLimit int `json:"battle_time_limit,omitempty"`
	// Average winning bid
	AvgWonBet int `json:"avg_won_bet,omitempty"`
	// Average minimum bid
	AvgMinBet int `json:"avg_min_bet,omitempty"`
	// Average clans rating
	AvgClansRating int `json:"avg_clans_rating,omitempty"`
	// Available modules
	AvailableExtensions struct {
		// Modules upkeep cost
		Wage int `json:"wage,omitempty"`
		// Localized consumable name
		Name string `json:"name,omitempty"`
		// Consumable ID
		ExtensionId string `json:"extension_id,omitempty"`
		// Localized description of tactical effect
		DescriptionTactical string `json:"description_tactical,omitempty"`
		// Localized description of strategic effect
		DescriptionStrategic string `json:"description_strategic,omitempty"`
		// Cost of modules
		Cost int `json:"cost,omitempty"`
	} `json:"available_extensions,omitempty"`
} 
