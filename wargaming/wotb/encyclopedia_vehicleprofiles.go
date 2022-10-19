// Auto generated file!

package wotb

type EncyclopediaVehicleprofilesOptions struct {
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
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
	// Sorting. Valid values:
	//
	// "price_credit" - by cost in credits
	// "-price_credit" - by cost in credits, in reverse order
	OrderBy *string `json:"order_by,omitempty"`
}

type EncyclopediaVehicleprofiles struct {
	// Armor
	Armor *struct {
		// Hull armor
		Hull *struct {
			// Front (mm)
			Front *int `json:"front,omitempty"`
			// Rear (mm)
			Rear *int `json:"rear,omitempty"`
			// Sides (mm)
			Sides *int `json:"sides,omitempty"`
		} `json:"hull,omitempty"`
		// Turret armor
		Turret *struct {
			// Front (mm)
			Front *int `json:"front,omitempty"`
			// Rear (mm)
			Rear *int `json:"rear,omitempty"`
			// Sides (mm)
			Sides *int `json:"sides,omitempty"`
		} `json:"turret,omitempty"`
	} `json:"armor,omitempty"`
	// The highest battle Tier of the vehicle
	BattleLevelRangeMax *int `json:"battle_level_range_max,omitempty"`
	// The lowest battle Tier of the vehicle
	BattleLevelRangeMin *int `json:"battle_level_range_min,omitempty"`
	// Engine characteristics
	Engine *struct {
		// Engine Power (hp)
		Power *int `json:"power,omitempty"`
	} `json:"engine,omitempty"`
	// Firepower (%)
	Firepower *int `json:"firepower,omitempty"`
	// Gun characteristics
	Gun *struct {
		// Aiming time (s)
		AimTime *float32 `json:"aim_time,omitempty"`
		// Caliber (mm)
		Caliber *int `json:"caliber,omitempty"`
		// Number of shells in the ammo
		ClipCapacity *int `json:"clip_capacity,omitempty"`
		// Reload time
		ClipReloadTime *float32 `json:"clip_reload_time,omitempty"`
		// Dispersion at 100 m (m)
		Dispersion *float32 `json:"dispersion,omitempty"`
		// Rate of fire (rounds/min)
		FireRate *float32 `json:"fire_rate,omitempty"`
		// Depression angle (deg)
		MoveDownArc *int `json:"move_down_arc,omitempty"`
		// Elevation angle (deg)
		MoveUpArc *int `json:"move_up_arc,omitempty"`
		// Reload time (s)
		ReloadTime *float32 `json:"reload_time,omitempty"`
		// Traverse speed (deg/s)
		TraverseSpeed *float32 `json:"traverse_speed,omitempty"`
	} `json:"gun,omitempty"`
	// Hit points
	Hp *int `json:"hp,omitempty"`
	// Hull weight (kg)
	HullWeight *int `json:"hull_weight,omitempty"`
	// Standard configuration
	IsDefault *bool `json:"is_default,omitempty"`
	// Maneuverability (%)
	Maneuverability *int `json:"maneuverability,omitempty"`
	// Ammunition
	MaxAmmo *int `json:"max_ammo,omitempty"`
	// Load limit (kg)
	MaxWeight *int `json:"max_weight,omitempty"`
	// Configuration cost in credits
	PriceCredit *int `json:"price_credit,omitempty"`
	// Configuration cost in experience
	PriceXp *int `json:"price_xp,omitempty"`
	// Vehicle Configuration ID
	ProfileId *string `json:"profile_id,omitempty"`
	// Armor protection (%)
	Protection *int `json:"protection,omitempty"`
	// Gun shells characteristics
	Shells *struct {
		// Average damage (HP)
		Damage *int `json:"damage,omitempty"`
		// Average penetration (mm)
		Penetration *int `json:"penetration,omitempty"`
	} `json:"shells,omitempty"`
	// Shot efficiency (%)
	ShotEfficiency *int `json:"shot_efficiency,omitempty"`
	// Signal range
	SignalRange *int `json:"signal_range,omitempty"`
	// Top reverse speed (km/h)
	SpeedBackward *int `json:"speed_backward,omitempty"`
	// Top speed (km/h)
	SpeedForward *int `json:"speed_forward,omitempty"`
	// Suspension characteristics
	Suspension *struct {
		// Load limit (kg)
		LoadLimit *int `json:"load_limit,omitempty"`
		// Traverse speed (deg/s)
		TraverseSpeed *int `json:"traverse_speed,omitempty"`
	} `json:"suspension,omitempty"`
	// Vehicle ID
	TankId *int `json:"tank_id,omitempty"`
	// Turret characteristics
	Turret *struct {
		// Traverse angle, left (deg)
		TraverseLeftArc *int `json:"traverse_left_arc,omitempty"`
		// Traverse angle, right (deg)
		TraverseRightArc *int `json:"traverse_right_arc,omitempty"`
		// Traverse speed (deg/s)
		TraverseSpeed *int `json:"traverse_speed,omitempty"`
		// View range (m)
		ViewRange *int `json:"view_range,omitempty"`
	} `json:"turret,omitempty"`
	// Weight (kg)
	Weight *int `json:"weight,omitempty"`
}
