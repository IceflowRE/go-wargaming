package wowp

type EncyclopediaPlaneinfo struct {
	// Type
	Type_ string `json:"type,omitempty"`
	// Localized short name of aircraft
	ShortNameI18N string `json:"short_name_i18n,omitempty"`
	// Localized short name of aircraft
	ShortName18N string `json:"short_name_18n,omitempty"`
	// Purchase cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Purchase cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Ancestors of the vehicle first generation
	PrevPlanes []int `json:"prev_planes,omitempty"`
	// Aircraft ID
	PlaneId int `json:"plane_id,omitempty"`
	// Descendants of the vehicle first generation
	NextPlanes []int `json:"next_planes,omitempty"`
	// Localized nation field
	NationI18N string `json:"nation_i18n,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Localized name field
	NameI18N string `json:"name_i18n,omitempty"`
	// Name
	Name string `json:"name,omitempty"`
	// Tier
	Level int `json:"level,omitempty"`
	// Indicates if the aircraft is Premium aircraft
	IsPremium bool `json:"is_premium,omitempty"`
	// Indicates if the aircraft was a gift
	IsGift bool `json:"is_gift,omitempty"`
	// Aircraft images
	Images struct {
		// URL to 51 x 31 px image of aircraft
		Small string `json:"small,omitempty"`
		// URL to 102 x 63 px image of aircraft
		Medium string `json:"medium,omitempty"`
		// URL to 408 x 252 px image of aircraft
		Large string `json:"large,omitempty"`
	} `json:"images,omitempty"`
	// Calculation of specifications
	Features struct {
		// Stall speed
		StallSpeed int `json:"stall_speed,omitempty"`
		// Airspeed
		SpeedFactor int `json:"speed_factor,omitempty"`
		// Top speed at sea level
		SpeedAtTheGround int `json:"speed_at_the_ground,omitempty"`
		// Number of slots
		SlotsCount int `json:"slots_count,omitempty"`
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
	} `json:"features,omitempty"`
	// Localized description of aircraft
	Description string `json:"description,omitempty"`
	// Crew
	Crew struct {
		// Localized role field
		RoleI18N string `json:"role_i18n,omitempty"`
		// Qualification of crew member
		Role string `json:"role,omitempty"`
		// Number
		Count int `json:"count,omitempty"`
	} `json:"crew,omitempty"`
} 
