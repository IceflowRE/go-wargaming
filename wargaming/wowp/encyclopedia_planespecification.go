// Auto generated file!

package wowp

type EncyclopediaPlanespecificationOptions struct {
	// ID of unique bind of slot and module. Maximum limit: 100.
	BindId []int `json:"bind_id,omitempty"`
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
	// "zh-cn" - 简体中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
	// Module ID. Maximum limit: 100.
	ModuleId []int `json:"module_id,omitempty"`
}

type EncyclopediaPlanespecification struct {
	// Slots and modules
	Slots *struct {
		// ID of unique bind of slot and module
		BindId *int `json:"bind_id,omitempty"`
		// Standard module
		IsDefault *bool `json:"is_default,omitempty"`
		// Module ID
		ModuleId *int `json:"module_id,omitempty"`
		// Slot name
		SlotName *string `json:"slot_name,omitempty"`
		// Localized slot_name field
		SlotNameI18n *string `json:"slot_name_i18n,omitempty"`
	} `json:"slots,omitempty"`
	// Aircraft specifications
	Specification *struct {
		// Average time to turn 360 deg
		AverageTurnTime *float32 `json:"average_turn_time,omitempty"`
		// Controllability
		Controllability *int `json:"controllability,omitempty"`
		// Maximum dive speed
		DiveSpeed *int `json:"dive_speed,omitempty"`
		// Firepower
		Dps *int `json:"dps,omitempty"`
		// Hit points
		Hp *int `json:"hp,omitempty"`
		// Maneuverability
		Maneuverability *int `json:"maneuverability,omitempty"`
		// Weight
		Mass *float32 `json:"mass,omitempty"`
		// Top speed at best altitude
		MaxSpeed *int `json:"max_speed,omitempty"`
		// Optimum altitude
		OptimalHeight *int `json:"optimal_height,omitempty"`
		// Optimum airspeed
		OptimalManeuverSpeed *int `json:"optimal_maneuver_speed,omitempty"`
		// Rate of climb
		RateOfClimbing *float32 `json:"rate_of_climbing,omitempty"`
		// Rate of roll
		RollManeuverability *int `json:"roll_maneuverability,omitempty"`
		// Top speed at sea level
		SpeedAtTheGround *int `json:"speed_at_the_ground,omitempty"`
		// Airspeed
		SpeedFactor *int `json:"speed_factor,omitempty"`
		// Stall speed
		StallSpeed *int `json:"stall_speed,omitempty"`
	} `json:"specification,omitempty"`
}
