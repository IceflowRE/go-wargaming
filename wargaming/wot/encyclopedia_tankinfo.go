// Auto generated file!

package wot

type EncyclopediaTankinfoOptions struct {
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
}

type EncyclopediaTankinfo struct {
	// Compatible suspensions
	Chassis *struct {
		// Indicates if the module is standard
		IsDefault *bool `json:"is_default,omitempty"`
		// Module ID
		ModuleId *int `json:"module_id,omitempty"`
	} `json:"chassis,omitempty"`
	// Standard suspension traverse speed
	ChassisRotationSpeed *int `json:"chassis_rotation_speed,omitempty"`
	// Standard turret view range
	CircularVisionRadius *int `json:"circular_vision_radius,omitempty"`
	// URL to outline image of vehicle
	ContourImage *string `json:"contour_image,omitempty"`
	// Crew
	Crew *struct {
		// Additional qualifications of crew member
		AdditionalRoles []string `json:"additional_roles,omitempty"`
		// Additional qualifications of crew member
		AdditionalRolesI18n *struct {
			// Qualification of crew member
			Role *string `json:"role,omitempty"`
			// Localized role field
			RoleI18n *string `json:"role_i18n,omitempty"`
		} `json:"additional_roles_i18n,omitempty"`
		// Qualification of crew member
		Role *string `json:"role,omitempty"`
		// Localized role field
		RoleI18n *string `json:"role_i18n,omitempty"`
	} `json:"crew,omitempty"`
	// Standard engine power
	EnginePower *int `json:"engine_power,omitempty"`
	// Compatible engines
	Engines *struct {
		// Indicates if the module is standard
		IsDefault *bool `json:"is_default,omitempty"`
		// Module ID
		ModuleId *int `json:"module_id,omitempty"`
	} `json:"engines,omitempty"`
	// Maximum damage of standard gun
	GunDamageMax *int `json:"gun_damage_max,omitempty"`
	// Minimum damage of standard gun
	GunDamageMin *int `json:"gun_damage_min,omitempty"`
	// Ammunition of standard gun
	GunMaxAmmo *int `json:"gun_max_ammo,omitempty"`
	// Standard gun name
	GunName *string `json:"gun_name,omitempty"`
	// Maximum penetration of standard gun
	GunPiercingPowerMax *int `json:"gun_piercing_power_max,omitempty"`
	// Minimum penetration of standard gun
	GunPiercingPowerMin *int `json:"gun_piercing_power_min,omitempty"`
	// Rate of fire of standard gun
	GunRate *float32 `json:"gun_rate,omitempty"`
	// Compatible guns
	Guns *struct {
		// Indicates if the module is standard
		IsDefault *bool `json:"is_default,omitempty"`
		// Module ID
		ModuleId *int `json:"module_id,omitempty"`
	} `json:"guns,omitempty"`
	// URL to 160 x 100 px image of vehicle
	Image *string `json:"image,omitempty"`
	// URL to 124 x 31 px image of vehicle
	ImageSmall *string `json:"image_small,omitempty"`
	// Indicates if the vehicle was a gift
	IsGift *bool `json:"is_gift,omitempty"`
	// Indicates if the vehicle is Premium vehicle
	IsPremium *bool `json:"is_premium,omitempty"`
	// Tier
	Level *int `json:"level,omitempty"`
	// Load limit
	LimitWeight *float32 `json:"limit_weight,omitempty"`
	// Localized name field
	LocalizedName *string `json:"localized_name,omitempty"`
	// Hit points
	MaxHealth *int `json:"max_health,omitempty"`
	// Vehicle name
	Name *string `json:"name,omitempty"`
	// Localized name field
	NameI18n *string `json:"name_i18n,omitempty"`
	// Nation
	Nation *string `json:"nation,omitempty"`
	// Localized nation field
	NationI18n *string `json:"nation_i18n,omitempty"`
	// Parent vehicles in Tech Tree
	ParentTanks []int `json:"parent_tanks,omitempty"`
	// Purchase cost in credits
	PriceCredit *int `json:"price_credit,omitempty"`
	// Purchase cost in gold
	PriceGold *int `json:"price_gold,omitempty"`
	// Cost of research in experience
	PriceXp *int `json:"price_xp,omitempty"`
	// Signal range of standard radio
	RadioDistance *int `json:"radio_distance,omitempty"`
	// Compatible radios
	Radios *struct {
		// Indicates if the module is standard
		IsDefault *bool `json:"is_default,omitempty"`
		// Module ID
		ModuleId *int `json:"module_id,omitempty"`
	} `json:"radios,omitempty"`
	// Localized short name of vehicle
	ShortNameI18n *string `json:"short_name_i18n,omitempty"`
	// Speed limit
	SpeedLimit *float32 `json:"speed_limit,omitempty"`
	// Vehicle ID
	TankId *int `json:"tank_id,omitempty"`
	// Standard turret armor: sides
	TurretArmorBoard *int `json:"turret_armor_board,omitempty"`
	// Standard turret armor: sides
	TurretArmorFedd *int `json:"turret_armor_fedd,omitempty"`
	// Standard turret armor: front
	TurretArmorForehead *int `json:"turret_armor_forehead,omitempty"`
	// Standard turret traverse speed
	TurretRotationSpeed *int `json:"turret_rotation_speed,omitempty"`
	// Compatible turrets
	Turrets *struct {
		// Indicates if the module is standard
		IsDefault *bool `json:"is_default,omitempty"`
		// Module ID
		ModuleId *int `json:"module_id,omitempty"`
	} `json:"turrets,omitempty"`
	// Vehicle type
	Type *string `json:"type,omitempty"`
	// Localized vehicle type
	TypeI18n *string `json:"type_i18n,omitempty"`
	// Hull armor: sides
	VehicleArmorBoard *int `json:"vehicle_armor_board,omitempty"`
	// Hull armor: rear
	VehicleArmorFedd *int `json:"vehicle_armor_fedd,omitempty"`
	// Hull armor: front
	VehicleArmorForehead *int `json:"vehicle_armor_forehead,omitempty"`
	// Weight
	Weight *float32 `json:"weight,omitempty"`
}
