package wot

type EncyclopediaTankinfo struct {
	// Weight
	Weight float32 `json:"weight,omitempty"`
	// Hull armor: front
	VehicleArmorForehead int `json:"vehicle_armor_forehead,omitempty"`
	// Hull armor: rear
	VehicleArmorFedd int `json:"vehicle_armor_fedd,omitempty"`
	// Hull armor: sides
	VehicleArmorBoard int `json:"vehicle_armor_board,omitempty"`
	// Vehicle type
	Type_ string `json:"type,omitempty"`
	// Localized vehicle type
	TypeI18N string `json:"type_i18n,omitempty"`
	// Compatible turrets
	Turrets struct {
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
	} `json:"turrets,omitempty"`
	// Standard turret traverse speed
	TurretRotationSpeed int `json:"turret_rotation_speed,omitempty"`
	// Standard turret armor: front
	TurretArmorForehead int `json:"turret_armor_forehead,omitempty"`
	// Standard turret armor: sides
	TurretArmorFedd int `json:"turret_armor_fedd,omitempty"`
	// Standard turret armor: sides
	TurretArmorBoard int `json:"turret_armor_board,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Speed limit
	SpeedLimit float32 `json:"speed_limit,omitempty"`
	// Localized short name of vehicle
	ShortNameI18N string `json:"short_name_i18n,omitempty"`
	// Compatible radios
	Radios struct {
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
	} `json:"radios,omitempty"`
	// Signal range of standard radio
	RadioDistance int `json:"radio_distance,omitempty"`
	// Cost of research in experience
	PriceXp int `json:"price_xp,omitempty"`
	// Purchase cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Purchase cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Parent vehicles in Tech Tree
	ParentTanks []int `json:"parent_tanks,omitempty"`
	// Localized nation field
	NationI18N string `json:"nation_i18n,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Localized name field
	NameI18N string `json:"name_i18n,omitempty"`
	// Vehicle name
	Name string `json:"name,omitempty"`
	// Hit points
	MaxHealth int `json:"max_health,omitempty"`
	// Localized name field
	LocalizedName string `json:"localized_name,omitempty"`
	// Load limit
	LimitWeight float32 `json:"limit_weight,omitempty"`
	// Tier
	Level int `json:"level,omitempty"`
	// Indicates if the vehicle is Premium vehicle
	IsPremium bool `json:"is_premium,omitempty"`
	// Indicates if the vehicle was a gift
	IsGift bool `json:"is_gift,omitempty"`
	// URL to 124 x 31 px image of vehicle
	ImageSmall string `json:"image_small,omitempty"`
	// URL to 160 x 100 px image of vehicle
	Image string `json:"image,omitempty"`
	// Compatible guns
	Guns struct {
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
	} `json:"guns,omitempty"`
	// Rate of fire of standard gun
	GunRate float32 `json:"gun_rate,omitempty"`
	// Minimum penetration of standard gun
	GunPiercingPowerMin int `json:"gun_piercing_power_min,omitempty"`
	// Maximum penetration of standard gun
	GunPiercingPowerMax int `json:"gun_piercing_power_max,omitempty"`
	// Standard gun name
	GunName string `json:"gun_name,omitempty"`
	// Ammunition of standard gun
	GunMaxAmmo int `json:"gun_max_ammo,omitempty"`
	// Minimum damage of standard gun
	GunDamageMin int `json:"gun_damage_min,omitempty"`
	// Maximum damage of standard gun
	GunDamageMax int `json:"gun_damage_max,omitempty"`
	// Compatible engines
	Engines struct {
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
	} `json:"engines,omitempty"`
	// Standard engine power
	EnginePower int `json:"engine_power,omitempty"`
	// Crew
	Crew struct {
		// Localized role field
		RoleI18N string `json:"role_i18n,omitempty"`
		// Qualification of crew member
		Role string `json:"role,omitempty"`
		// Additional qualifications of crew member
		AdditionalRolesI18N struct {
			// Localized role field
			RoleI18N string `json:"role_i18n,omitempty"`
			// Qualification of crew member
			Role string `json:"role,omitempty"`
		} `json:"additional_roles_i18n,omitempty"`
		// Additional qualifications of crew member
		AdditionalRoles []string `json:"additional_roles,omitempty"`
	} `json:"crew,omitempty"`
	// URL to outline image of vehicle
	ContourImage string `json:"contour_image,omitempty"`
	// Standard turret view range
	CircularVisionRadius int `json:"circular_vision_radius,omitempty"`
	// Standard suspension traverse speed
	ChassisRotationSpeed int `json:"chassis_rotation_speed,omitempty"`
	// Compatible suspensions
	Chassis struct {
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
	} `json:"chassis,omitempty"`
} 
