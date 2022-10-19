// Auto generated file!

package wot

type EncyclopediaModulesOptions struct {
	// Extra fields that will be added to the response. Valid values:
	//
	// "default_profile"
	Extra []string `json:"extra,omitempty"`
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
	// Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
	Limit *int `json:"limit,omitempty"`
	// Module ID. Maximum limit: 100.
	ModuleId []int `json:"module_id,omitempty"`
	// Nation. Maximum limit: 100.
	Nation []string `json:"nation,omitempty"`
	// Result page number
	PageNo *int `json:"page_no,omitempty"`
	// Module type. Maximum limit: 100. Valid values:
	//
	// "vehicleRadio" - Radio
	// "vehicleEngine" - Engines
	// "vehicleGun" - Gun
	// "vehicleChassis" - Suspension
	// "vehicleTurret" - Turret
	Type []string `json:"type,omitempty"`
}

type EncyclopediaModules struct {
	// Basic technical characteristics of module.
	// An extra field.
	DefaultProfile *struct {
		// Engine characteristics
		Engine *struct {
			// Chance of engine fire
			FireChance *float32 `json:"fire_chance,omitempty"`
			// Engine Power (hp)
			Power *int `json:"power,omitempty"`
		} `json:"engine,omitempty"`
		// Gun characteristics
		Gun *struct {
			// Aiming time (s)
			AimTime *float32 `json:"aim_time,omitempty"`
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
			// Dispersion at 100 m (m)
			Dispersion *float32 `json:"dispersion,omitempty"`
			// Rate of fire (rounds/min)
			FireRate *float32 `json:"fire_rate,omitempty"`
			// Number of shells
			MaxAmmo *int `json:"max_ammo,omitempty"`
			// Depression angle (deg)
			MoveDownArc *int `json:"move_down_arc,omitempty"`
			// Elevation angle (deg)
			MoveUpArc *int `json:"move_up_arc,omitempty"`
			// Reload time (s)
			ReloadTime *float32 `json:"reload_time,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed *int `json:"traverse_speed,omitempty"`
		} `json:"gun,omitempty"`
		// Radio characteristics
		Radio *struct {
			// Signal range
			SignalRange *int `json:"signal_range,omitempty"`
		} `json:"radio,omitempty"`
		// Suspension characteristics
		Suspension *struct {
			// Load limit (kg)
			LoadLimit *int `json:"load_limit,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed *int `json:"traverse_speed,omitempty"`
		} `json:"suspension,omitempty"`
		// Turret characteristics
		Turret *struct {
			// Armor: front (mm)
			ArmorFront *int `json:"armor_front,omitempty"`
			// Armor: rear (mm)
			ArmorRear *int `json:"armor_rear,omitempty"`
			// Armor: sides (mm)
			ArmorSides *int `json:"armor_sides,omitempty"`
			// Hit points
			Hp *int `json:"hp,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed *int `json:"traverse_speed,omitempty"`
			// View range (m)
			ViewRange *int `json:"view_range,omitempty"`
		} `json:"turret,omitempty"`
	} `json:"default_profile,omitempty"`
	// Image link
	Image *string `json:"image,omitempty"`
	// Module ID
	ModuleId *int `json:"module_id,omitempty"`
	// Module name
	Name *string `json:"name,omitempty"`
	// Nation
	Nation *string `json:"nation,omitempty"`
	// Cost in credits
	PriceCredit *int `json:"price_credit,omitempty"`
	// Vehicles compatible with module
	Tanks []int `json:"tanks,omitempty"`
	// Tier
	Tier *int `json:"tier,omitempty"`
	// Module type
	Type *string `json:"type,omitempty"`
	// Weight (kg)
	Weight *int `json:"weight,omitempty"`
}
