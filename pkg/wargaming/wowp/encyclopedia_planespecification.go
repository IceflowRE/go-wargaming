package wowp

type EncyclopediaPlanespecification struct {
	// Aircraft specifications
	Specification struct {
		// Stall speed
		StallSpeed int `json:"stall_speed,omitempty"`
		// Airspeed
		SpeedFactor int `json:"speed_factor,omitempty"`
		// Top speed at sea level
		SpeedAtTheGround int `json:"speed_at_the_ground,omitempty"`
		// Rate of roll
		RollManeuverability int `json:"roll_maneuverability,omitempty"`
		// Rate of climb
		RateOfClimbing float32 `json:"rate_of_climbing,omitempty"`
		// Optimum airspeed
		OptimalManeuverSpeed int `json:"optimal_maneuver_speed,omitempty"`
		// Optimum altitude
		OptimalHeight int `json:"optimal_height,omitempty"`
		// Top speed at best altitude
		MaxSpeed int `json:"max_speed,omitempty"`
		// Weight
		Mass float32 `json:"mass,omitempty"`
		// Maneuverability
		Maneuverability int `json:"maneuverability,omitempty"`
		// Hit points
		Hp int `json:"hp,omitempty"`
		// Firepower
		Dps int `json:"dps,omitempty"`
		// Maximum dive speed
		DiveSpeed int `json:"dive_speed,omitempty"`
		// Controllability
		Controllability int `json:"controllability,omitempty"`
		// Average time to turn 360 deg
		AverageTurnTime float32 `json:"average_turn_time,omitempty"`
	} `json:"specification,omitempty"`
	// Slots and modules
	Slots struct {
		// Localized slot_name field
		SlotNameI18N string `json:"slot_name_i18n,omitempty"`
		// Slot name
		SlotName string `json:"slot_name,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Standard module
		IsDefault bool `json:"is_default,omitempty"`
		// ID of unique bind of slot and module
		BindId int `json:"bind_id,omitempty"`
	} `json:"slots,omitempty"`
} 
