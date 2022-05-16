package wotx

type EncyclopediaModules struct {
	// Weight (kg)
	Weight int `json:"weight,omitempty"`
	// Module type
	Type_ string `json:"type,omitempty"`
	// Tier
	Tier int `json:"tier,omitempty"`
	// Vehicles compatible with module
	Tanks []int `json:"tanks,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Module name
	Name string `json:"name,omitempty"`
	// Module ID
	ModuleId int `json:"module_id,omitempty"`
	// Image link
	Image string `json:"image,omitempty"`
	// Basic technical characteristics of module.
	// An extra field.
	DefaultProfile struct {
		// Turret characteristics
		Turret struct {
			// View range (m)
			ViewRange int `json:"view_range,omitempty"`
			// Hit points
			Hp int `json:"hp,omitempty"`
			// Armor: sides (mm)
			ArmorSides int `json:"armor_sides,omitempty"`
			// Armor: rear (mm)
			ArmorRear int `json:"armor_rear,omitempty"`
			// Armor: front (mm)
			ArmorFront int `json:"armor_front,omitempty"`
		} `json:"turret,omitempty"`
		// Suspension characteristics
		Suspension struct {
			// Traverse speed (deg/s)
			TraverseSpeed int `json:"traverse_speed,omitempty"`
			// Load limit (kg)
			LoadLimit int `json:"load_limit,omitempty"`
		} `json:"suspension,omitempty"`
		// Radio characteristics
		Radio struct {
			// Signal range
			SignalRange int `json:"signal_range,omitempty"`
		} `json:"radio,omitempty"`
		// Gun characteristics
		Gun struct {
			// Traverse speed (deg/s)
			TraverseSpeed int `json:"traverse_speed,omitempty"`
			// Reload time (s)
			ReloadTime float32 `json:"reload_time,omitempty"`
			// Elevation angle (deg)
			MoveUpArc int `json:"move_up_arc,omitempty"`
			// Depression angle (deg)
			MoveDownArc int `json:"move_down_arc,omitempty"`
			// Number of shells
			MaxAmmo int `json:"max_ammo,omitempty"`
			// Rate of fire (rounds/min)
			FireRate float32 `json:"fire_rate,omitempty"`
			// Dispersion at 100 m (m)
			Dispersion float32 `json:"dispersion,omitempty"`
			// Gun shells characteristics
			Ammo struct {
				// Shell type
				Type_ string `json:"type,omitempty"`
				// Stun characteristics
				Stun struct {
					// Stun duration (s) caused by this shell type, a list of values: min, max
					Duration struct {
					} `json:"duration,omitempty"`
				} `json:"stun,omitempty"`
				// Penetration (mm), a list of values: min, avg, max
				Penetration []int `json:"penetration,omitempty"`
				// Damage (hp), a list of values: min, avg, max
				Damage []int `json:"damage,omitempty"`
			} `json:"ammo,omitempty"`
			// Aiming time (s)
			AimTime float32 `json:"aim_time,omitempty"`
		} `json:"gun,omitempty"`
		// Engine characteristics
		Engine struct {
			// Engine Power (hp)
			Power int `json:"power,omitempty"`
			// Chance of engine fire
			FireChance float32 `json:"fire_chance,omitempty"`
		} `json:"engine,omitempty"`
	} `json:"default_profile,omitempty"`
} 
