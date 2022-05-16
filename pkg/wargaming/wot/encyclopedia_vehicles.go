package wot

type EncyclopediaVehicles struct {
	// Vehicle type
	Type_ string `json:"type,omitempty"`
	// List of compatible turret IDs
	Turrets []int `json:"turrets,omitempty"`
	// Tier
	Tier int `json:"tier,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Vehicle tag
	Tag string `json:"tag,omitempty"`
	// List of compatible suspension IDs
	Suspensions []int `json:"suspensions,omitempty"`
	// Vehicle short name
	ShortName string `json:"short_name,omitempty"`
	// List of compatible radio IDs
	Radios []int `json:"radios,omitempty"`
	// List of IDs of compatible equipment and consumables
	Provisions []int `json:"provisions,omitempty"`
	// List of research costs in form of pairs:
	// 
	// parent vehicle ID
	// cost of research in XP
	PricesXp map[string]string `json:"prices_xp,omitempty"`
	// Cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// List of vehicles available for research in form of pairs:
	// 
	// researched vehicle ID
	// cost of research in XP
	NextTanks map[string]string `json:"next_tanks,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Vehicle name
	Name string `json:"name,omitempty"`
	// Информация об мультинации
	Multination struct {
		// Vehicle ID
		TankId int `json:"tank_id,omitempty"`
		// Признак нации по-умолчанию
		IsDefault bool `json:"is_default,omitempty"`
	} `json:"multination,omitempty"`
	// Module research information
	ModulesTree struct {
		// Module type
		Type_ string `json:"type,omitempty"`
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// Cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// List of vehicle IDs available after research of the module
		NextTanks []int `json:"next_tanks,omitempty"`
		// List of module IDs available after research of the module
		NextModules []int `json:"next_modules,omitempty"`
		// Module name
		Name string `json:"name,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Indicates if the module is basic
		IsDefault bool `json:"is_default,omitempty"`
	} `json:"modules_tree,omitempty"`
	// Indicates if the vehicle is a wheeled vehicle
	IsWheeled bool `json:"is_wheeled,omitempty"`
	// Indicates the IGR vehicle. Active only for Korea realm
	IsPremiumIgr bool `json:"is_premium_igr,omitempty"`
	// Indicates if the vehicle is Premium vehicle
	IsPremium bool `json:"is_premium,omitempty"`
	// Indicates if the vehicle is a gift vehicle
	IsGift bool `json:"is_gift,omitempty"`
	// Image links
	Images struct {
		// URL to 124 x 31 px image of vehicle
		SmallIcon string `json:"small_icon,omitempty"`
		// URL to outline image of vehicle
		ContourIcon string `json:"contour_icon,omitempty"`
		// URL to 160 x 100 px image of vehicle
		BigIcon string `json:"big_icon,omitempty"`
	} `json:"images,omitempty"`
	// List of compatible gun IDs
	Guns []int `json:"guns,omitempty"`
	// List of compatible engine IDs
	Engines []int `json:"engines,omitempty"`
	// Vehicle description
	Description string `json:"description,omitempty"`
	// Standard configuration characteristics
	DefaultProfile struct {
		// Weight (kg)
		Weight int `json:"weight,omitempty"`
		// Turret characteristics
		Turret struct {
			// Weight (kg)
			Weight int `json:"weight,omitempty"`
			// View range (m)
			ViewRange int `json:"view_range,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed int `json:"traverse_speed,omitempty"`
			// Traverse angle, right (deg)
			TraverseRightArc int `json:"traverse_right_arc,omitempty"`
			// Traverse angle, left (deg)
			TraverseLeftArc int `json:"traverse_left_arc,omitempty"`
			// Tier
			Tier int `json:"tier,omitempty"`
			// Module tag
			Tag string `json:"tag,omitempty"`
			// Module name
			Name string `json:"name,omitempty"`
			// Hit points
			Hp int `json:"hp,omitempty"`
		} `json:"turret,omitempty"`
		// Suspension characteristics
		Suspension struct {
			// Weight (kg)
			Weight int `json:"weight,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed int `json:"traverse_speed,omitempty"`
			// Tier
			Tier int `json:"tier,omitempty"`
			// Module tag
			Tag string `json:"tag,omitempty"`
			// Maximum wheel turning angle (for wheeled vehicles)
			SteeringLockAngle int `json:"steering_lock_angle,omitempty"`
			// Module name
			Name string `json:"name,omitempty"`
			// Load limit (kg)
			LoadLimit int `json:"load_limit,omitempty"`
		} `json:"suspension,omitempty"`
		// Top speed (km/h)
		SpeedForward int `json:"speed_forward,omitempty"`
		// Top reverse speed (km/h)
		SpeedBackward int `json:"speed_backward,omitempty"`
		// Vehicle characteristics in Siege mode
		Siege struct {
			// Time required to switch to Siege mode
			SwitchOnTime float32 `json:"switch_on_time,omitempty"`
			// Time needed to switch on the Siege mode
			SwitchOffTime float32 `json:"switch_off_time,omitempty"`
			// Standard suspension traverse speed
			SuspensionTraverseSpeed int `json:"suspension_traverse_speed,omitempty"`
			// Top reverse speed (km/h)
			SpeedBackward int `json:"speed_backward,omitempty"`
			// Reload time (s)
			ReloadTime float32 `json:"reload_time,omitempty"`
			// Elevation angle (deg)
			MoveUpArc int `json:"move_up_arc,omitempty"`
			// Depression angle (deg)
			MoveDownArc int `json:"move_down_arc,omitempty"`
			// Dispersion at 100 m (m)
			Dispersion float32 `json:"dispersion,omitempty"`
			// Aiming time (s)
			AimTime float32 `json:"aim_time,omitempty"`
		} `json:"siege,omitempty"`
		// Vehicle characteristics in Rapid mode (for wheeled vehicles)
		Rapid struct {
			// Time required to switch to Rapid mode
			SwitchOnTime float32 `json:"switch_on_time,omitempty"`
			// Time required to switch to Cruise mode
			SwitchOffTime float32 `json:"switch_off_time,omitempty"`
			// Maximum wheel turning angle
			SuspensionSteeringLockAngle int `json:"suspension_steering_lock_angle,omitempty"`
			// Top speed (km/h)
			SpeedForward int `json:"speed_forward,omitempty"`
			// Top reverse speed (km/h)
			SpeedBackward int `json:"speed_backward,omitempty"`
		} `json:"rapid,omitempty"`
		// Radio characteristics
		Radio struct {
			// Weight (kg)
			Weight int `json:"weight,omitempty"`
			// Tier
			Tier int `json:"tier,omitempty"`
			// Module tag
			Tag string `json:"tag,omitempty"`
			// Signal range
			SignalRange int `json:"signal_range,omitempty"`
			// Module name
			Name string `json:"name,omitempty"`
		} `json:"radio,omitempty"`
		// Mounted modules
		Modules struct {
			// Turret ID
			TurretId int `json:"turret_id,omitempty"`
			// Suspension ID
			SuspensionId int `json:"suspension_id,omitempty"`
			// Radio ID
			RadioId int `json:"radio_id,omitempty"`
			// Gun ID
			GunId int `json:"gun_id,omitempty"`
			// Engine ID
			EngineId int `json:"engine_id,omitempty"`
		} `json:"modules,omitempty"`
		// Load limit (kg)
		MaxWeight int `json:"max_weight,omitempty"`
		// Ammunition
		MaxAmmo int `json:"max_ammo,omitempty"`
		// Hull weight (kg)
		HullWeight int `json:"hull_weight,omitempty"`
		// Hull HP
		HullHp int `json:"hull_hp,omitempty"`
		// Hit points
		Hp int `json:"hp,omitempty"`
		// Gun characteristics
		Gun struct {
			// Weight (kg)
			Weight int `json:"weight,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed int `json:"traverse_speed,omitempty"`
			// Tier
			Tier int `json:"tier,omitempty"`
			// Module tag
			Tag string `json:"tag,omitempty"`
			// Reload time (s)
			ReloadTime float32 `json:"reload_time,omitempty"`
			// Module name
			Name string `json:"name,omitempty"`
			// Elevation angle (deg)
			MoveUpArc int `json:"move_up_arc,omitempty"`
			// Depression angle (deg)
			MoveDownArc int `json:"move_down_arc,omitempty"`
			// Rate of fire (rounds/min)
			FireRate float32 `json:"fire_rate,omitempty"`
			// Dispersion at 100 m (m)
			Dispersion float32 `json:"dispersion,omitempty"`
			// Caliber (mm)
			Caliber int `json:"caliber,omitempty"`
			// Aiming time (s)
			AimTime float32 `json:"aim_time,omitempty"`
		} `json:"gun,omitempty"`
		// Engine characteristics
		Engine struct {
			// Weight (kg)
			Weight int `json:"weight,omitempty"`
			// Tier
			Tier int `json:"tier,omitempty"`
			// Module tag
			Tag string `json:"tag,omitempty"`
			// Engine Power (hp)
			Power int `json:"power,omitempty"`
			// Module name
			Name string `json:"name,omitempty"`
			// Chance of engine fire
			FireChance float32 `json:"fire_chance,omitempty"`
		} `json:"engine,omitempty"`
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
		// Gun shells characteristics
		Ammo struct {
			// Shell type
			Type_ string `json:"type,omitempty"`
			// Stun characteristics
			Stun struct {
				// Stun duration (s) caused by this shell type, a list of values: min, max
				Duration []int `json:"duration,omitempty"`
			} `json:"stun,omitempty"`
			// Penetration (mm), a list of values: min, avg, max
			Penetration []int `json:"penetration,omitempty"`
			// Damage (hp), a list of values: min, avg, max
			Damage []int `json:"damage,omitempty"`
		} `json:"ammo,omitempty"`
	} `json:"default_profile,omitempty"`
	// Crew
	Crew struct {
		// List of crew member roles
		Roles map[string]string `json:"roles,omitempty"`
		// Crew member ID
		MemberId string `json:"member_id,omitempty"`
	} `json:"crew,omitempty"`
} 
