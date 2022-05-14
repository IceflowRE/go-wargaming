package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WowpEncyclopediaPlanemodules struct {
	// Bomb
	Bomb struct {
		// Bomb ID
		BombId int `json:"bomb_id,omitempty"`
		// Damage
		ExplosionDamageMax int `json:"explosion_damage_max,omitempty"`
		// Damage radius
		ExplosionRadius int `json:"explosion_radius,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
		// Tier
		Level int `json:"level,omitempty"`
		// Weight
		Mass float32 `json:"mass,omitempty"`
		// Name
		Name string `json:"name,omitempty"`
		// Localized name field
		NameI18N string `json:"name_i18n,omitempty"`
		// Purchase cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
	} `json:"bomb,omitempty"`
	// Airframe
	Construction struct {
		// Airframe ID
		ConstructionId int `json:"construction_id,omitempty"`
		// Hit points
		Hp int `json:"hp,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
		// Tier
		Level int `json:"level,omitempty"`
		// Weight
		Mass float32 `json:"mass,omitempty"`
		// Name
		Name string `json:"name,omitempty"`
		// Localized name field
		NameI18N string `json:"name_i18n,omitempty"`
		// Purchase cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
	} `json:"construction,omitempty"`
	// Engine
	Engine struct {
		// Engine ID
		EngineId int `json:"engine_id,omitempty"`
		// Type
		EquipmentType string `json:"equipment_type,omitempty"`
		// Localized equipment_type field
		EquipmentTypeI18N string `json:"equipment_type_i18n,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
		// Tier
		Level int `json:"level,omitempty"`
		// Weight
		Mass float32 `json:"mass,omitempty"`
		// Name
		Name string `json:"name,omitempty"`
		// Localized name field
		NameI18N string `json:"name_i18n,omitempty"`
		// Purchase cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
		// Power
		Power int `json:"power,omitempty"`
		// Thrust
		Thrust int `json:"thrust,omitempty"`
	} `json:"engine,omitempty"`
	// Autocannon
	Gun struct {
		// Muzzle velocity
		BulletVelocity int `json:"bullet_velocity,omitempty"`
		// Caliber
		Caliber float32 `json:"caliber,omitempty"`
		// Damage
		Dps int `json:"dps,omitempty"`
		// Gun ID
		GunId int `json:"gun_id,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
		// Tier
		Level int `json:"level,omitempty"`
		// Weight
		Mass float32 `json:"mass,omitempty"`
		// Name
		Name string `json:"name,omitempty"`
		// Localized name field
		NameI18N string `json:"name_i18n,omitempty"`
		// Purchase cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// Rate of fire
		Rpm int `json:"rpm,omitempty"`
		// Type
		Type_ string `json:"type,omitempty"`
		// Localized type field
		TypeI18N string `json:"type_i18n,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
	} `json:"gun,omitempty"`
	// Rocket
	Rocket struct {
		// Damage
		ExplosionDamageMax int `json:"explosion_damage_max,omitempty"`
		// Damage radius
		ExplosionRadius int `json:"explosion_radius,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
		// Tier
		Level int `json:"level,omitempty"`
		// Weight
		Mass float32 `json:"mass,omitempty"`
		// Name
		Name string `json:"name,omitempty"`
		// Localized name field
		NameI18N string `json:"name_i18n,omitempty"`
		// Purchase cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// Rocket ID
		RocketId int `json:"rocket_id,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
	} `json:"rocket,omitempty"`
	// Rear gun
	Turret struct {
		// Muzzle velocity
		BulletVelocity int `json:"bullet_velocity,omitempty"`
		// Caliber
		Caliber float32 `json:"caliber,omitempty"`
		// Damage
		Dps int `json:"dps,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
		// Tier
		Level int `json:"level,omitempty"`
		// Weight
		Mass float32 `json:"mass,omitempty"`
		// Name
		Name string `json:"name,omitempty"`
		// Localized name field
		NameI18N string `json:"name_i18n,omitempty"`
		// Purchase cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// Rate of fire
		Rpm int `json:"rpm,omitempty"`
		// Rear gun ID
		TurretId int `json:"turret_id,omitempty"`
		// Autocannon
		WeaponCount int `json:"weapon_count,omitempty"`
		// List of compatible aircraft
		CompatibleVehicles []int `json:"compatible_vehicles,omitempty"`
	} `json:"turret,omitempty"`
}

// WowpEncyclopediaPlanemodules Method returns information from Encyclopedia about modules available for specified aircrafts.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planemodules
//
// plane_id:
//     Aircraft ID. Maximum limit: 1000.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "en" &mdash; English 
//     "ru" &mdash; Русский (by default)
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "zh-cn" &mdash; 简体中文 
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
// type:
//     Configuration. Default is "engine, bomb, rocket, turret, turretfront, turretupper, turretlower, gun, construction". Maximum limit: 100. Valid values:
//     
//     "engine" &mdash; engine 
//     "bomb" &mdash; bomb 
//     "rocket" &mdash; rocket 
//     "turret" &mdash; rear gun 
//     "turretfront" &mdash; Передняя турель 
//     "turretupper" &mdash; Верхняя турель 
//     "turretlower" &mdash; Нижняя турель 
//     "gun" &mdash; autocannon 
//     "construction" &mdash; airframe
func (client *Client) WowpEncyclopediaPlanemodules(realm Realm, planeId []int, fields []string, language string, type_ []string) (*WowpEncyclopediaPlanemodules, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"plane_id": utils.SliceIntToString(planeId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
		"type": strings.Join(type_, ","),
	}

	var result *WowpEncyclopediaPlanemodules
	err := client.doGetDataRequest(realm, "/wowp/encyclopedia/planemodules/", reqParam, &result)
	return result, err
}
