// Auto generated file!

package wowp

type EncyclopediaPlaneinfoOptions struct {
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
}

type EncyclopediaPlaneinfo struct {
	// Crew
	Crew *struct {
		// Number
		Count *int `json:"count,omitempty"`
		// Qualification of crew member
		Role *string `json:"role,omitempty"`
		// Localized role field
		RoleI18n *string `json:"role_i18n,omitempty"`
	} `json:"crew,omitempty"`
	// Localized description of aircraft
	Description *string `json:"description,omitempty"`
	// Calculation of specifications
	Features *struct {
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
		// Number of slots
		SlotsCount *int `json:"slots_count,omitempty"`
		// Top speed at sea level
		SpeedAtTheGround *int `json:"speed_at_the_ground,omitempty"`
		// Airspeed
		SpeedFactor *int `json:"speed_factor,omitempty"`
		// Stall speed
		StallSpeed *int `json:"stall_speed,omitempty"`
	} `json:"features,omitempty"`
	// Aircraft images
	Images *struct {
		// URL to 408 x 252 px image of aircraft
		Large *string `json:"large,omitempty"`
		// URL to 102 x 63 px image of aircraft
		Medium *string `json:"medium,omitempty"`
		// URL to 51 x 31 px image of aircraft
		Small *string `json:"small,omitempty"`
	} `json:"images,omitempty"`
	// Indicates if the aircraft was a gift
	IsGift *bool `json:"is_gift,omitempty"`
	// Indicates if the aircraft is Premium aircraft
	IsPremium *bool `json:"is_premium,omitempty"`
	// Tier
	Level *int `json:"level,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	// Localized name field
	NameI18n *string `json:"name_i18n,omitempty"`
	// Nation
	Nation *string `json:"nation,omitempty"`
	// Localized nation field
	NationI18n *string `json:"nation_i18n,omitempty"`
	// Descendants of the vehicle first generation
	NextPlanes []int `json:"next_planes,omitempty"`
	// Aircraft ID
	PlaneId *int `json:"plane_id,omitempty"`
	// Ancestors of the vehicle first generation
	PrevPlanes []int `json:"prev_planes,omitempty"`
	// Purchase cost in credits
	PriceCredit *int `json:"price_credit,omitempty"`
	// Purchase cost in gold
	PriceGold *int `json:"price_gold,omitempty"`
	// Localized short name of aircraft
	ShortName18n *string `json:"short_name_18n,omitempty"`
	// Localized short name of aircraft
	ShortNameI18n *string `json:"short_name_i18n,omitempty"`
	// Type
	Type *string `json:"type,omitempty"`
}
