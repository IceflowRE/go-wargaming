// Auto generated file!

package wotb

type EncyclopediaModulesOptions struct {
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
	// Module ID. Maximum limit: 100.
	ModuleId []int `json:"module_id,omitempty"`
	// Nation
	Nation *string `json:"nation,omitempty"`
	// Module type. Valid values:
	//
	// "vehicleEngine" - Engines
	// "vehicleGun" - Gun
	// "vehicleChassis" - Suspension
	// "vehicleTurret" - Turret
	Type *string `json:"type,omitempty"`
}

type EncyclopediaModules struct {
	// Engine characteristics
	Engines []*struct {
		// Chance of engine fire
		FireChance *float32 `json:"fire_chance,omitempty"`
		// Module ID
		ModuleId *int `json:"module_id,omitempty"`
		// Vehicle name
		Name *string `json:"name,omitempty"`
		// Nation
		Nation *string `json:"nation,omitempty"`
		// Engine Power (hp)
		Power *int `json:"power,omitempty"`
		// List of compatible vehicles
		Tanks []int `json:"tanks,omitempty"`
		// Tier
		Tier *int `json:"tier,omitempty"`
		// Weight (kg)
		Weight *int `json:"weight,omitempty"`
	} `json:"engines,omitempty"`
	// Gun characteristics
	Guns []*struct {
		// Aiming time (s)
		AimTime *float32 `json:"aim_time,omitempty"`
		// Dispersion at 100 m (m)
		Dispersion *float32 `json:"dispersion,omitempty"`
		// Module ID
		ModuleId *int `json:"module_id,omitempty"`
		// Vehicle name
		Name *string `json:"name,omitempty"`
		// Nation
		Nation *string `json:"nation,omitempty"`
		// Gun shells characteristics
		Shells []*struct {
			// Average damage (HP)
			Damage *int `json:"damage,omitempty"`
			// Average penetration (mm)
			Penetration *int `json:"penetration,omitempty"`
			// Type
			Type *string `json:"type,omitempty"`
		} `json:"shells,omitempty"`
		// List of compatible vehicles
		Tanks []int `json:"tanks,omitempty"`
		// Tier
		Tier *int `json:"tier,omitempty"`
		// Weight (kg)
		Weight *int `json:"weight,omitempty"`
	} `json:"guns,omitempty"`
	// Suspension characteristics
	Suspensions []*struct {
		// Load limit (kg)
		LoadLimit *int `json:"load_limit,omitempty"`
		// Module ID
		ModuleId *int `json:"module_id,omitempty"`
		// Vehicle name
		Name *string `json:"name,omitempty"`
		// Nation
		Nation *string `json:"nation,omitempty"`
		// List of compatible vehicles
		Tanks []int `json:"tanks,omitempty"`
		// Tier
		Tier *int `json:"tier,omitempty"`
		// Traverse speed (deg/s)
		TraverseSpeed *int `json:"traverse_speed,omitempty"`
		// Weight (kg)
		Weight *int `json:"weight,omitempty"`
	} `json:"suspensions,omitempty"`
	// Turret characteristics
	Turrets []*struct {
		// Armor
		Armor *struct {
			// Front (mm)
			Front *int `json:"front,omitempty"`
			// Rear (mm)
			Rear *int `json:"rear,omitempty"`
			// Sides (mm)
			Sides *int `json:"sides,omitempty"`
		} `json:"armor,omitempty"`
		// Hit points
		Hp *int `json:"hp,omitempty"`
		// Module ID
		ModuleId *int `json:"module_id,omitempty"`
		// Vehicle name
		Name *string `json:"name,omitempty"`
		// Nation
		Nation *string `json:"nation,omitempty"`
		// List of compatible vehicles
		Tanks []int `json:"tanks,omitempty"`
		// Tier
		Tier *int `json:"tier,omitempty"`
		// Traverse angle, left (deg)
		TraverseLeftArc *int `json:"traverse_left_arc,omitempty"`
		// Traverse angle, right (deg)
		TraverseRightArc *int `json:"traverse_right_arc,omitempty"`
		// View range (m)
		ViewRange *int `json:"view_range,omitempty"`
		// Weight (kg)
		Weight *int `json:"weight,omitempty"`
	} `json:"turrets,omitempty"`
}
