// Auto generated file!

package wowp

type EncyclopediaPlanemodulesOptions struct {
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
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
	// Configuration. Default is "engine, bomb, rocket, turret, turretfront, turretupper, turretlower, gun, construction". Maximum limit: 100. Valid values:
	//
	// "engine" - engine
	// "bomb" - bomb
	// "rocket" - rocket
	// "turret" - rear gun
	// "turretfront" - Передняя турель
	// "turretupper" - Верхняя турель
	// "turretlower" - Нижняя турель
	// "gun" - autocannon
	// "construction" - airframe
	Type []string `json:"type,omitempty"`
}

type EncyclopediaPlanemodules struct {
	// Bomb
	Bomb *struct {
		// Bomb ID
		BombId *int `json:"bomb_id,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
		// Damage
		ExplosionDamageMax *int `json:"explosion_damage_max,omitempty"`
		// Damage radius
		ExplosionRadius *int `json:"explosion_radius,omitempty"`
		// URL to image
		Image *string `json:"image,omitempty"`
		// Indicates if the module is standard
		IsDefault *bool `json:"is_default,omitempty"`
		// Tier
		Level *int `json:"level,omitempty"`
		// Weight
		Mass *float32 `json:"mass,omitempty"`
		// Name
		Name *string `json:"name,omitempty"`
		// Localized name field
		NameI18n *string `json:"name_i18n,omitempty"`
		// Purchase cost in credits
		PriceCredit *int `json:"price_credit,omitempty"`
		// Research cost
		PriceXp *int `json:"price_xp,omitempty"`
	} `json:"bomb,omitempty"`
	// Airframe
	Construction *struct {
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
		// Airframe ID
		ConstructionId *int `json:"construction_id,omitempty"`
		// Hit points
		Hp *int `json:"hp,omitempty"`
		// URL to image
		Image *string `json:"image,omitempty"`
		// Indicates if the module is standard
		IsDefault *bool `json:"is_default,omitempty"`
		// Tier
		Level *int `json:"level,omitempty"`
		// Weight
		Mass *float32 `json:"mass,omitempty"`
		// Name
		Name *string `json:"name,omitempty"`
		// Localized name field
		NameI18n *string `json:"name_i18n,omitempty"`
		// Purchase cost in credits
		PriceCredit *int `json:"price_credit,omitempty"`
		// Research cost
		PriceXp *int `json:"price_xp,omitempty"`
	} `json:"construction,omitempty"`
	// Engine
	Engine *struct {
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
		// Engine ID
		EngineId *int `json:"engine_id,omitempty"`
		// Type
		EquipmentType *string `json:"equipment_type,omitempty"`
		// Localized equipment_type field
		EquipmentTypeI18n *string `json:"equipment_type_i18n,omitempty"`
		// URL to image
		Image *string `json:"image,omitempty"`
		// Indicates if the module is standard
		IsDefault *bool `json:"is_default,omitempty"`
		// Tier
		Level *int `json:"level,omitempty"`
		// Weight
		Mass *float32 `json:"mass,omitempty"`
		// Name
		Name *string `json:"name,omitempty"`
		// Localized name field
		NameI18n *string `json:"name_i18n,omitempty"`
		// Power
		Power *int `json:"power,omitempty"`
		// Purchase cost in credits
		PriceCredit *int `json:"price_credit,omitempty"`
		// Research cost
		PriceXp *int `json:"price_xp,omitempty"`
		// Thrust
		Thrust *int `json:"thrust,omitempty"`
	} `json:"engine,omitempty"`
	// Autocannon
	Gun *struct {
		// Muzzle velocity
		BulletVelocity *int `json:"bullet_velocity,omitempty"`
		// Caliber
		Caliber *float32 `json:"caliber,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
		// Damage
		Dps *int `json:"dps,omitempty"`
		// Gun ID
		GunId *int `json:"gun_id,omitempty"`
		// URL to image
		Image *string `json:"image,omitempty"`
		// Indicates if the module is standard
		IsDefault *bool `json:"is_default,omitempty"`
		// Tier
		Level *int `json:"level,omitempty"`
		// Weight
		Mass *float32 `json:"mass,omitempty"`
		// Name
		Name *string `json:"name,omitempty"`
		// Localized name field
		NameI18n *string `json:"name_i18n,omitempty"`
		// Purchase cost in credits
		PriceCredit *int `json:"price_credit,omitempty"`
		// Research cost
		PriceXp *int `json:"price_xp,omitempty"`
		// Rate of fire
		Rpm *int `json:"rpm,omitempty"`
		// Type
		Type *string `json:"type,omitempty"`
		// Localized type field
		TypeI18n *string `json:"type_i18n,omitempty"`
	} `json:"gun,omitempty"`
	// Rocket
	Rocket *struct {
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
		// Damage
		ExplosionDamageMax *int `json:"explosion_damage_max,omitempty"`
		// Damage radius
		ExplosionRadius *int `json:"explosion_radius,omitempty"`
		// URL to image
		Image *string `json:"image,omitempty"`
		// Indicates if the module is standard
		IsDefault *bool `json:"is_default,omitempty"`
		// Tier
		Level *int `json:"level,omitempty"`
		// Weight
		Mass *float32 `json:"mass,omitempty"`
		// Name
		Name *string `json:"name,omitempty"`
		// Localized name field
		NameI18n *string `json:"name_i18n,omitempty"`
		// Purchase cost in credits
		PriceCredit *int `json:"price_credit,omitempty"`
		// Research cost
		PriceXp *int `json:"price_xp,omitempty"`
		// Rocket ID
		RocketId *int `json:"rocket_id,omitempty"`
	} `json:"rocket,omitempty"`
	// Rear gun
	Turret *struct {
		// Muzzle velocity
		BulletVelocity *int `json:"bullet_velocity,omitempty"`
		// Caliber
		Caliber *float32 `json:"caliber,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
		// Damage
		Dps *int `json:"dps,omitempty"`
		// URL to image
		Image *string `json:"image,omitempty"`
		// Indicates if the module is standard
		IsDefault *bool `json:"is_default,omitempty"`
		// Tier
		Level *int `json:"level,omitempty"`
		// Weight
		Mass *float32 `json:"mass,omitempty"`
		// Name
		Name *string `json:"name,omitempty"`
		// Localized name field
		NameI18n *string `json:"name_i18n,omitempty"`
		// Purchase cost in credits
		PriceCredit *int `json:"price_credit,omitempty"`
		// Research cost
		PriceXp *int `json:"price_xp,omitempty"`
		// Rate of fire
		Rpm *int `json:"rpm,omitempty"`
		// Rear gun ID
		TurretId *int `json:"turret_id,omitempty"`
		// Autocannon
		WeaponCount *int `json:"weapon_count,omitempty"`
	} `json:"turret,omitempty"`
}
