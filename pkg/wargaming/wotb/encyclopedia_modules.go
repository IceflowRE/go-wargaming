package wotb

type EncyclopediaModules struct {
	// Turret characteristics
	Turrets struct {
		// Weight (kg)
		Weight int `json:"weight,omitempty"`
		// View range (m)
		ViewRange int `json:"view_range,omitempty"`
		// Traverse angle, right (deg)
		TraverseRightArc int `json:"traverse_right_arc,omitempty"`
		// Traverse angle, left (deg)
		TraverseLeftArc int `json:"traverse_left_arc,omitempty"`
		// Tier
		Tier int `json:"tier,omitempty"`
		// List of compatible vehicles
		Tanks []int `json:"tanks,omitempty"`
		// Nation
		Nation string `json:"nation,omitempty"`
		// Vehicle name
		Name string `json:"name,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Hit points
		Hp int `json:"hp,omitempty"`
		// Armor
		Armor struct {
			// Sides (mm)
			Sides int `json:"sides,omitempty"`
			// Rear (mm)
			Rear int `json:"rear,omitempty"`
			// Front (mm)
			Front int `json:"front,omitempty"`
		} `json:"armor,omitempty"`
	} `json:"turrets,omitempty"`
	// Suspension characteristics
	Suspensions struct {
		// Weight (kg)
		Weight int `json:"weight,omitempty"`
		// Traverse speed (deg/s)
		TraverseSpeed int `json:"traverse_speed,omitempty"`
		// Tier
		Tier int `json:"tier,omitempty"`
		// List of compatible vehicles
		Tanks []int `json:"tanks,omitempty"`
		// Nation
		Nation string `json:"nation,omitempty"`
		// Vehicle name
		Name string `json:"name,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Load limit (kg)
		LoadLimit int `json:"load_limit,omitempty"`
	} `json:"suspensions,omitempty"`
	// Gun characteristics
	Guns struct {
		// Weight (kg)
		Weight int `json:"weight,omitempty"`
		// Tier
		Tier int `json:"tier,omitempty"`
		// List of compatible vehicles
		Tanks []int `json:"tanks,omitempty"`
		// Gun shells characteristics
		Shells struct {
			// Type
			Type_ string `json:"type,omitempty"`
			// Average penetration (mm)
			Penetration int `json:"penetration,omitempty"`
			// Average damage (HP)
			Damage int `json:"damage,omitempty"`
		} `json:"shells,omitempty"`
		// Nation
		Nation string `json:"nation,omitempty"`
		// Vehicle name
		Name string `json:"name,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Dispersion at 100 m (m)
		Dispersion float32 `json:"dispersion,omitempty"`
		// Aiming time (s)
		AimTime float32 `json:"aim_time,omitempty"`
	} `json:"guns,omitempty"`
	// Engine characteristics
	Engines struct {
		// Weight (kg)
		Weight int `json:"weight,omitempty"`
		// Tier
		Tier int `json:"tier,omitempty"`
		// List of compatible vehicles
		Tanks []int `json:"tanks,omitempty"`
		// Engine Power (hp)
		Power int `json:"power,omitempty"`
		// Nation
		Nation string `json:"nation,omitempty"`
		// Vehicle name
		Name string `json:"name,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Chance of engine fire
		FireChance float32 `json:"fire_chance,omitempty"`
	} `json:"engines,omitempty"`
} 
