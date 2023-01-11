// Auto generated file!

package wotb

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
	// "zh-cn" - 简体中文
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
	// Nation. Maximum limit: 100.
	Nation []string `json:"nation,omitempty"`
	// Vehicle ID. Maximum limit: 100.
	TankId []int `json:"tank_id,omitempty"`
}

type EncyclopediaVehicles struct {
	// Cost
	Cost *struct {
		// Cost in credits
		PriceCredit *int `json:"price_credit,omitempty"`
		// Purchase cost in gold
		PriceGold *int `json:"price_gold,omitempty"`
	} `json:"cost,omitempty"`
	// Standard configuration characteristics
	DefaultProfile *struct {
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
			// Chance of engine fire
			FireChance *float32 `json:"fire_chance,omitempty"`
			// Vehicle name
			Name *string `json:"name,omitempty"`
			// Engine Power (hp)
			Power *int `json:"power,omitempty"`
			// Tier
			Tier *int `json:"tier,omitempty"`
			// Weight (kg)
			Weight *int `json:"weight,omitempty"`
		} `json:"engine,omitempty"`
		// Engine ID
		EngineId *int `json:"engine_id,omitempty"`
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
			// Vehicle name
			Name *string `json:"name,omitempty"`
			// Reload time (s)
			ReloadTime *float32 `json:"reload_time,omitempty"`
			// Tier
			Tier *int `json:"tier,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed *float32 `json:"traverse_speed,omitempty"`
			// Weight (kg)
			Weight *int `json:"weight,omitempty"`
		} `json:"gun,omitempty"`
		// Gun ID
		GunId *int `json:"gun_id,omitempty"`
		// Hit points
		Hp *int `json:"hp,omitempty"`
		// Hull HP
		HullHp *int `json:"hull_hp,omitempty"`
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
			// Type
			Type *string `json:"type,omitempty"`
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
			// Vehicle name
			Name *string `json:"name,omitempty"`
			// Tier
			Tier *int `json:"tier,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed *int `json:"traverse_speed,omitempty"`
			// Weight (kg)
			Weight *int `json:"weight,omitempty"`
		} `json:"suspension,omitempty"`
		// Suspension ID
		SuspensionId *int `json:"suspension_id,omitempty"`
		// Turret characteristics
		Turret *struct {
			// Hit points
			Hp *int `json:"hp,omitempty"`
			// Vehicle name
			Name *string `json:"name,omitempty"`
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
		// Turret ID
		TurretId *int `json:"turret_id,omitempty"`
		// Weight (kg)
		Weight *int `json:"weight,omitempty"`
	} `json:"default_profile,omitempty"`
	// Vehicle description
	Description *string `json:"description,omitempty"`
	// List of compatible engine IDs
	Engines []int `json:"engines,omitempty"`
	// List of compatible gun IDs
	Guns []int `json:"guns,omitempty"`
	// Image links
	Images *struct {
		// Normal image
		Normal *string `json:"normal,omitempty"`
		// Small size image
		Preview *string `json:"preview,omitempty"`
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
	// List of research costs in form of pairs:
	//
	// parent vehicle ID
	// cost of research in XP
	PricesXp map[string]string `json:"prices_xp,omitempty"`
	// List of compatible suspension IDs
	Suspensions []int `json:"suspensions,omitempty"`
	// Vehicle ID
	TankId *int `json:"tank_id,omitempty"`
	// Tier
	Tier *int `json:"tier,omitempty"`
	// List of compatible turret IDs
	Turrets []int `json:"turrets,omitempty"`
	// Vehicle type
	Type *string `json:"type,omitempty"`
}
