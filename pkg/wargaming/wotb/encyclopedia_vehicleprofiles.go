package wotb

type EncyclopediaVehicleprofiles struct {
	// Weight (kg)
	Weight int `json:"weight,omitempty"`
	// Turret characteristics
	Turret struct {
		// View range (m)
		ViewRange int `json:"view_range,omitempty"`
		// Traverse speed (deg/s)
		TraverseSpeed int `json:"traverse_speed,omitempty"`
		// Traverse angle, right (deg)
		TraverseRightArc int `json:"traverse_right_arc,omitempty"`
		// Traverse angle, left (deg)
		TraverseLeftArc int `json:"traverse_left_arc,omitempty"`
	} `json:"turret,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Suspension characteristics
	Suspension struct {
		// Traverse speed (deg/s)
		TraverseSpeed int `json:"traverse_speed,omitempty"`
		// Load limit (kg)
		LoadLimit int `json:"load_limit,omitempty"`
	} `json:"suspension,omitempty"`
	// Top speed (km/h)
	SpeedForward int `json:"speed_forward,omitempty"`
	// Top reverse speed (km/h)
	SpeedBackward int `json:"speed_backward,omitempty"`
	// Signal range
	SignalRange int `json:"signal_range,omitempty"`
	// Shot efficiency (%)
	ShotEfficiency int `json:"shot_efficiency,omitempty"`
	// Gun shells characteristics
	Shells struct {
		// Average penetration (mm)
		Penetration int `json:"penetration,omitempty"`
		// Average damage (HP)
		Damage int `json:"damage,omitempty"`
	} `json:"shells,omitempty"`
	// Armor protection (%)
	Protection int `json:"protection,omitempty"`
	// Vehicle Configuration ID
	ProfileId string `json:"profile_id,omitempty"`
	// Configuration cost in experience
	PriceXp int `json:"price_xp,omitempty"`
	// Configuration cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Load limit (kg)
	MaxWeight int `json:"max_weight,omitempty"`
	// Ammunition
	MaxAmmo int `json:"max_ammo,omitempty"`
	// Maneuverability (%)
	Maneuverability int `json:"maneuverability,omitempty"`
	// Standard configuration
	IsDefault bool `json:"is_default,omitempty"`
	// Hull weight (kg)
	HullWeight int `json:"hull_weight,omitempty"`
	// Hit points
	Hp int `json:"hp,omitempty"`
	// Gun characteristics
	Gun struct {
		// Traverse speed (deg/s)
		TraverseSpeed float32 `json:"traverse_speed,omitempty"`
		// Reload time (s)
		ReloadTime float32 `json:"reload_time,omitempty"`
		// Elevation angle (deg)
		MoveUpArc int `json:"move_up_arc,omitempty"`
		// Depression angle (deg)
		MoveDownArc int `json:"move_down_arc,omitempty"`
		// Rate of fire (rounds/min)
		FireRate float32 `json:"fire_rate,omitempty"`
		// Dispersion at 100 m (m)
		Dispersion float32 `json:"dispersion,omitempty"`
		// Reload time
		ClipReloadTime float32 `json:"clip_reload_time,omitempty"`
		// Number of shells in the ammo
		ClipCapacity int `json:"clip_capacity,omitempty"`
		// Caliber (mm)
		Caliber int `json:"caliber,omitempty"`
		// Aiming time (s)
		AimTime float32 `json:"aim_time,omitempty"`
	} `json:"gun,omitempty"`
	// Firepower (%)
	Firepower int `json:"firepower,omitempty"`
	// Engine characteristics
	Engine struct {
		// Engine Power (hp)
		Power int `json:"power,omitempty"`
	} `json:"engine,omitempty"`
	// The lowest battle Tier of the vehicle
	BattleLevelRangeMin int `json:"battle_level_range_min,omitempty"`
	// The highest battle Tier of the vehicle
	BattleLevelRangeMax int `json:"battle_level_range_max,omitempty"`
	// Armor
	Armor struct {
		// Turret armor
		Turret struct {
			// Sides (mm)
			Sides int `json:"sides,omitempty"`
			// Rear (mm)
			Rear int `json:"rear,omitempty"`
			// Front (mm)
			Front int `json:"front,omitempty"`
		} `json:"turret,omitempty"`
		// Hull armor
		Hull struct {
			// Sides (mm)
			Sides int `json:"sides,omitempty"`
			// Rear (mm)
			Rear int `json:"rear,omitempty"`
			// Front (mm)
			Front int `json:"front,omitempty"`
		} `json:"hull,omitempty"`
	} `json:"armor,omitempty"`
} 
