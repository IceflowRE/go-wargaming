// Auto generated file!

package wotx

type EncyclopediaVehicleprofileOptions struct {
	// Engine ID. If module is not specified, standard module is used by default.
	EngineId *int `json:"engine_id,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Gun ID. If module is not specified, standard module is used by default.
	GunId *int `json:"gun_id,omitempty"`
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
	// Configuration ID. If specified, parameters of IDs of separate modules are ignored.
	ProfileId *string `json:"profile_id,omitempty"`
	// Radio ID. If module is not specified, standard module is used by default.
	RadioId *int `json:"radio_id,omitempty"`
	// Suspension ID. If module is not specified, standard module is used by default.
	SuspensionId *int `json:"suspension_id,omitempty"`
	// Turret ID. If module is not specified, standard module is used by default.
	TurretId *int `json:"turret_id,omitempty"`
}

type EncyclopediaVehicleprofile struct {
	// Gun shells characteristics
	Ammo *struct {
		// Damage (hp), a list of values: min, avg, max
		Damage []int `json:"damage,omitempty"`
		// Penetration (mm), a list of values: min, avg, max
		Penetration []int `json:"penetration,omitempty"`
		// Stun characteristics
		Stun *struct {
			// Stun duration (s) caused by this shell type, a list of values: min, max
			Duration *struct {
			} `json:"duration,omitempty"`
		} `json:"stun,omitempty"`
		// Shell type
		Type *string `json:"type,omitempty"`
	} `json:"ammo,omitempty"`
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
	// Vehicle characteristics in autosiege mode
	Autosiege *struct {
		// Switch off speed (km/h)
		AutoswitchOffSpeed *int `json:"autoswitch_off_speed,omitempty"`
		// Switch on speed (km/h)
		AutoswitchOnSpeed *int `json:"autoswitch_on_speed,omitempty"`
		// Time needed to switch on the Siege mode
		SwitchOffTime *float32 `json:"switch_off_time,omitempty"`
		// Time required to switch to Siege mode
		SwitchOnTime *float32 `json:"switch_on_time,omitempty"`
	} `json:"autosiege,omitempty"`
	// Engine characteristics
	Engine *struct {
		// Chance of engine fire
		FireChance *float32 `json:"fire_chance,omitempty"`
		// Module name
		Name *string `json:"name,omitempty"`
		// Engine Power (hp)
		Power *int `json:"power,omitempty"`
		// Module tag
		Tag *string `json:"tag,omitempty"`
		// Tier
		Tier *int `json:"tier,omitempty"`
		// Weight (kg)
		Weight *int `json:"weight,omitempty"`
	} `json:"engine,omitempty"`
	// Gun characteristics
	Gun *struct {
		// Aiming time (s)
		AimTime *float32 `json:"aim_time,omitempty"`
		// Caliber (mm)
		Caliber *int `json:"caliber,omitempty"`
		// Dispersion at 100 m (m)
		Dispersion *float32 `json:"dispersion,omitempty"`
		// Rate of fire (rounds/min)
		FireRate *float32 `json:"fire_rate,omitempty"`
		// Depression angle (deg)
		MoveDownArc *int `json:"move_down_arc,omitempty"`
		// Elevation angle (deg)
		MoveUpArc *int `json:"move_up_arc,omitempty"`
		// Module name
		Name *string `json:"name,omitempty"`
		// Reload time (s)
		ReloadTime *float32 `json:"reload_time,omitempty"`
		// Module tag
		Tag *string `json:"tag,omitempty"`
		// Tier
		Tier *int `json:"tier,omitempty"`
		// Traverse speed (deg/s)
		TraverseSpeed *int `json:"traverse_speed,omitempty"`
		// Weight (kg)
		Weight *int `json:"weight,omitempty"`
	} `json:"gun,omitempty"`
	// Hit points
	Hp *int `json:"hp,omitempty"`
	// Hull HP
	HullHp *int `json:"hull_hp,omitempty"`
	// Hull weight (kg)
	HullWeight *int `json:"hull_weight,omitempty"`
	// Standard configuration
	IsDefault *bool `json:"is_default,omitempty"`
	// Ammunition
	MaxAmmo *int `json:"max_ammo,omitempty"`
	// Load limit (kg)
	MaxWeight *int `json:"max_weight,omitempty"`
	// Mounted modules
	Modules *struct {
		// Engine ID
		EngineId *int `json:"engine_id,omitempty"`
		// Gun ID
		GunId *int `json:"gun_id,omitempty"`
		// Radio ID
		RadioId *int `json:"radio_id,omitempty"`
		// Suspension ID
		SuspensionId *int `json:"suspension_id,omitempty"`
		// Turret ID
		TurretId *int `json:"turret_id,omitempty"`
	} `json:"modules,omitempty"`
	// Multi turret vehicle characteristics
	MultiTurret *struct {
		// List of vehicle turrets
		Turrets []int `json:"turrets,omitempty"`
	} `json:"multi_turret,omitempty"`
	// Multi weapon vehicle characteristics
	MultiWeapon *struct {
		// Grouping number on multi gun turrets
		GunGroupNumber *int `json:"gun_group_number,omitempty"`
		// List of turret guns per group number
		Guns []int `json:"guns,omitempty"`
		// Turret ID
		TurretId *int `json:"turret_id,omitempty"`
	} `json:"multi_weapon,omitempty"`
	// Vehicle Configuration ID
	ProfileId *string `json:"profile_id,omitempty"`
	// Radio characteristics
	Radio *struct {
		// Module name
		Name *string `json:"name,omitempty"`
		// Signal range
		SignalRange *int `json:"signal_range,omitempty"`
		// Module tag
		Tag *string `json:"tag,omitempty"`
		// Tier
		Tier *int `json:"tier,omitempty"`
		// Weight (kg)
		Weight *int `json:"weight,omitempty"`
	} `json:"radio,omitempty"`
	// Vehicle characteristics in Rapid mode (for wheeled vehicles)
	Rapid *struct {
		// Top reverse speed (km/h)
		SpeedBackward *int `json:"speed_backward,omitempty"`
		// Top speed (km/h)
		SpeedForward *int `json:"speed_forward,omitempty"`
		// Maximum wheel turning angle
		SuspensionSteeringLockAngle *int `json:"suspension_steering_lock_angle,omitempty"`
		// Time required to switch to Cruise mode
		SwitchOffTime *float32 `json:"switch_off_time,omitempty"`
		// Time required to switch to Rapid mode
		SwitchOnTime *float32 `json:"switch_on_time,omitempty"`
		// Traverse speed (deg/s)
		TurretTraverseSpeed *int `json:"turret_traverse_speed,omitempty"`
	} `json:"rapid,omitempty"`
	// Vehicle characteristics in Siege mode
	Siege *struct {
		// Aiming time (s)
		AimTime *float32 `json:"aim_time,omitempty"`
		// Dispersion at 100 m (m)
		Dispersion *float32 `json:"dispersion,omitempty"`
		// Top reverse speed (km/h)
		SpeedBackward *int `json:"speed_backward,omitempty"`
		// Top speed (km/h)
		SpeedForward *int `json:"speed_forward,omitempty"`
		// Time needed to switch on the Siege mode
		SwitchOffTime *float32 `json:"switch_off_time,omitempty"`
		// Time required to switch to Siege mode
		SwitchOnTime *float32 `json:"switch_on_time,omitempty"`
	} `json:"siege,omitempty"`
	// Top reverse speed (km/h)
	SpeedBackward *int `json:"speed_backward,omitempty"`
	// Top speed (km/h)
	SpeedForward *int `json:"speed_forward,omitempty"`
	// Suspension characteristics
	Suspension *struct {
		// Load limit (kg)
		LoadLimit *int `json:"load_limit,omitempty"`
		// Module name
		Name *string `json:"name,omitempty"`
		// Maximum wheel turning angle (for wheeled vehicles)
		SteeringLockAngle *int `json:"steering_lock_angle,omitempty"`
		// Module tag
		Tag *string `json:"tag,omitempty"`
		// Tier
		Tier *int `json:"tier,omitempty"`
		// Traverse speed (deg/s)
		TraverseSpeed *int `json:"traverse_speed,omitempty"`
		// Weight (kg)
		Weight *int `json:"weight,omitempty"`
	} `json:"suspension,omitempty"`
	// Vehicle ID
	TankId *int `json:"tank_id,omitempty"`
	// Vehicle characteristics in turbo mode
	Turbo *struct {
		// Aiming time (s)
		AimTime *float32 `json:"aim_time,omitempty"`
		// Top reverse speed (km/h)
		SpeedBackward *int `json:"speed_backward,omitempty"`
		// Top speed (km/h)
		SpeedForward *int `json:"speed_forward,omitempty"`
		// Time needed to switch on the Siege mode
		SwitchOffTime *float32 `json:"switch_off_time,omitempty"`
		// Time required to switch to Siege mode
		SwitchOnTime *float32 `json:"switch_on_time,omitempty"`
	} `json:"turbo,omitempty"`
	// Turret characteristics
	Turret *struct {
		// Hit points
		Hp *int `json:"hp,omitempty"`
		// Module name
		Name *string `json:"name,omitempty"`
		// Module tag
		Tag *string `json:"tag,omitempty"`
		// Tier
		Tier *int `json:"tier,omitempty"`
		// Traverse angle, left (deg)
		TraverseLeftArc *int `json:"traverse_left_arc,omitempty"`
		// Traverse angle, right (deg)
		TraverseRightArc *int `json:"traverse_right_arc,omitempty"`
		// Traverse speed (deg/s)
		TraverseSpeed *int `json:"traverse_speed,omitempty"`
		// View range (m)
		ViewRange *int `json:"view_range,omitempty"`
		// Weight (kg)
		Weight *int `json:"weight,omitempty"`
	} `json:"turret,omitempty"`
	// Weight (kg)
	Weight *int `json:"weight,omitempty"`
}
