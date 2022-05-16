package wowp

type EncyclopediaPlanemodules struct {
	// Rear gun
	Turret struct {
		// Autocannon
		WeaponCount int `json:"weapon_count,omitempty"`
		// Rear gun ID
		TurretId int `json:"turret_id,omitempty"`
		// Rate of fire
		Rpm int `json:"rpm,omitempty"`
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// Purchase cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Localized name field
		NameI18N string `json:"name_i18n,omitempty"`
		// Name
		Name string `json:"name,omitempty"`
		// Weight
		Mass float32 `json:"mass,omitempty"`
		// Tier
		Level int `json:"level,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Damage
		Dps int `json:"dps,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
		// Caliber
		Caliber float32 `json:"caliber,omitempty"`
		// Muzzle velocity
		BulletVelocity int `json:"bullet_velocity,omitempty"`
	} `json:"turret,omitempty"`
	// Rocket
	Rocket struct {
		// Rocket ID
		RocketId int `json:"rocket_id,omitempty"`
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// Purchase cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Localized name field
		NameI18N string `json:"name_i18n,omitempty"`
		// Name
		Name string `json:"name,omitempty"`
		// Weight
		Mass float32 `json:"mass,omitempty"`
		// Tier
		Level int `json:"level,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Damage radius
		ExplosionRadius int `json:"explosion_radius,omitempty"`
		// Damage
		ExplosionDamageMax int `json:"explosion_damage_max,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
	} `json:"rocket,omitempty"`
	// Autocannon
	Gun struct {
		// Type
		Type_ string `json:"type,omitempty"`
		// Localized type field
		TypeI18N string `json:"type_i18n,omitempty"`
		// Rate of fire
		Rpm int `json:"rpm,omitempty"`
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// Purchase cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Localized name field
		NameI18N string `json:"name_i18n,omitempty"`
		// Name
		Name string `json:"name,omitempty"`
		// Weight
		Mass float32 `json:"mass,omitempty"`
		// Tier
		Level int `json:"level,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Gun ID
		GunId int `json:"gun_id,omitempty"`
		// Damage
		Dps int `json:"dps,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
		// Caliber
		Caliber float32 `json:"caliber,omitempty"`
		// Muzzle velocity
		BulletVelocity int `json:"bullet_velocity,omitempty"`
	} `json:"gun,omitempty"`
	// Engine
	Engine struct {
		// Thrust
		Thrust int `json:"thrust,omitempty"`
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// Purchase cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Power
		Power int `json:"power,omitempty"`
		// Localized name field
		NameI18N string `json:"name_i18n,omitempty"`
		// Name
		Name string `json:"name,omitempty"`
		// Weight
		Mass float32 `json:"mass,omitempty"`
		// Tier
		Level int `json:"level,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Localized equipment_type field
		EquipmentTypeI18N string `json:"equipment_type_i18n,omitempty"`
		// Type
		EquipmentType string `json:"equipment_type,omitempty"`
		// Engine ID
		EngineId int `json:"engine_id,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
	} `json:"engine,omitempty"`
	// Airframe
	Construction struct {
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// Purchase cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Localized name field
		NameI18N string `json:"name_i18n,omitempty"`
		// Name
		Name string `json:"name,omitempty"`
		// Weight
		Mass float32 `json:"mass,omitempty"`
		// Tier
		Level int `json:"level,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Hit points
		Hp int `json:"hp,omitempty"`
		// Airframe ID
		ConstructionId int `json:"construction_id,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
	} `json:"construction,omitempty"`
	// Bomb
	Bomb struct {
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// Purchase cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Localized name field
		NameI18N string `json:"name_i18n,omitempty"`
		// Name
		Name string `json:"name,omitempty"`
		// Weight
		Mass float32 `json:"mass,omitempty"`
		// Tier
		Level int `json:"level,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Damage radius
		ExplosionRadius int `json:"explosion_radius,omitempty"`
		// Damage
		ExplosionDamageMax int `json:"explosion_damage_max,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
		// Bomb ID
		BombId int `json:"bomb_id,omitempty"`
	} `json:"bomb,omitempty"`
} 
