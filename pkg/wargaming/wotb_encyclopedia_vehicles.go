package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotbEncyclopediaVehicles struct {
	// Vehicle description
	Description string `json:"description,omitempty"`
	// List of compatible engine IDs
	Engines []int `json:"engines,omitempty"`
	// List of compatible gun IDs
	Guns []int `json:"guns,omitempty"`
	// Indicates if the vehicle is Premium vehicle
	IsPremium bool `json:"is_premium,omitempty"`
	// Vehicle name
	Name string `json:"name,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// List of vehicles available for research in form of pairs:
	// 
	// researched vehicle ID
	// cost of research in XP
	NextTanks map[string]string `json:"next_tanks,omitempty"`
	// List of research costs in form of pairs:
	// 
	// parent vehicle ID
	// cost of research in XP
	PricesXp map[string]string `json:"prices_xp,omitempty"`
	// List of compatible suspension IDs
	Suspensions []int `json:"suspensions,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Tier
	Tier int `json:"tier,omitempty"`
	// List of compatible turret IDs
	Turrets []int `json:"turrets,omitempty"`
	// Vehicle type
	Type_ string `json:"type,omitempty"`
	// Cost
	Cost struct {
		// Cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Purchase cost in gold
		PriceGold int `json:"price_gold,omitempty"`
	} `json:"cost,omitempty"`
	// Standard configuration characteristics
	DefaultProfile struct {
		// The highest battle Tier of the vehicle
		BattleLevelRangeMax int `json:"battle_level_range_max,omitempty"`
		// The lowest battle Tier of the vehicle
		BattleLevelRangeMin int `json:"battle_level_range_min,omitempty"`
		// Engine ID
		EngineId int `json:"engine_id,omitempty"`
		// Firepower (%)
		Firepower int `json:"firepower,omitempty"`
		// Gun ID
		GunId int `json:"gun_id,omitempty"`
		// Hit points
		Hp int `json:"hp,omitempty"`
		// Hull HP
		HullHp int `json:"hull_hp,omitempty"`
		// Hull weight (kg)
		HullWeight int `json:"hull_weight,omitempty"`
		// Standard configuration
		IsDefault bool `json:"is_default,omitempty"`
		// Maneuverability (%)
		Maneuverability int `json:"maneuverability,omitempty"`
		// Ammunition
		MaxAmmo int `json:"max_ammo,omitempty"`
		// Load limit (kg)
		MaxWeight int `json:"max_weight,omitempty"`
		// Vehicle Configuration ID
		ProfileId string `json:"profile_id,omitempty"`
		// Armor protection (%)
		Protection int `json:"protection,omitempty"`
		// Shot efficiency (%)
		ShotEfficiency int `json:"shot_efficiency,omitempty"`
		// Signal range
		SignalRange int `json:"signal_range,omitempty"`
		// Top reverse speed (km/h)
		SpeedBackward int `json:"speed_backward,omitempty"`
		// Top speed (km/h)
		SpeedForward int `json:"speed_forward,omitempty"`
		// Suspension ID
		SuspensionId int `json:"suspension_id,omitempty"`
		// Turret ID
		TurretId int `json:"turret_id,omitempty"`
		// Weight (kg)
		Weight int `json:"weight,omitempty"`
		// Armor
		Armor struct {
			// Hull armor
			Hull struct {
				// Front (mm)
				Front int `json:"front,omitempty"`
				// Rear (mm)
				Rear int `json:"rear,omitempty"`
				// Sides (mm)
				Sides int `json:"sides,omitempty"`
			} `json:"hull,omitempty"`
			// Turret armor
			Turret struct {
				// Front (mm)
				Front int `json:"front,omitempty"`
				// Rear (mm)
				Rear int `json:"rear,omitempty"`
				// Sides (mm)
				Sides int `json:"sides,omitempty"`
			} `json:"turret,omitempty"`
		} `json:"armor,omitempty"`
		// Engine characteristics
		Engine struct {
			// Chance of engine fire
			FireChance float32 `json:"fire_chance,omitempty"`
			// Vehicle name
			Name string `json:"name,omitempty"`
			// Engine Power (hp)
			Power int `json:"power,omitempty"`
			// Tier
			Tier int `json:"tier,omitempty"`
			// Weight (kg)
			Weight int `json:"weight,omitempty"`
		} `json:"engine,omitempty"`
		// Gun characteristics
		Gun struct {
			// Aiming time (s)
			AimTime float32 `json:"aim_time,omitempty"`
			// Caliber (mm)
			Caliber int `json:"caliber,omitempty"`
			// Number of shells in the ammo
			ClipCapacity int `json:"clip_capacity,omitempty"`
			// Reload time
			ClipReloadTime float32 `json:"clip_reload_time,omitempty"`
			// Dispersion at 100 m (m)
			Dispersion float32 `json:"dispersion,omitempty"`
			// Rate of fire (rounds/min)
			FireRate float32 `json:"fire_rate,omitempty"`
			// Depression angle (deg)
			MoveDownArc int `json:"move_down_arc,omitempty"`
			// Elevation angle (deg)
			MoveUpArc int `json:"move_up_arc,omitempty"`
			// Vehicle name
			Name string `json:"name,omitempty"`
			// Reload time (s)
			ReloadTime float32 `json:"reload_time,omitempty"`
			// Tier
			Tier int `json:"tier,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed float32 `json:"traverse_speed,omitempty"`
			// Weight (kg)
			Weight int `json:"weight,omitempty"`
		} `json:"gun,omitempty"`
		// Gun shells characteristics
		Shells struct {
			// Average damage (HP)
			Damage int `json:"damage,omitempty"`
			// Average penetration (mm)
			Penetration int `json:"penetration,omitempty"`
			// Type
			Type_ string `json:"type,omitempty"`
		} `json:"shells,omitempty"`
		// Suspension characteristics
		Suspension struct {
			// Load limit (kg)
			LoadLimit int `json:"load_limit,omitempty"`
			// Vehicle name
			Name string `json:"name,omitempty"`
			// Tier
			Tier int `json:"tier,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed int `json:"traverse_speed,omitempty"`
			// Weight (kg)
			Weight int `json:"weight,omitempty"`
		} `json:"suspension,omitempty"`
		// Turret characteristics
		Turret struct {
			// Hit points
			Hp int `json:"hp,omitempty"`
			// Vehicle name
			Name string `json:"name,omitempty"`
			// Tier
			Tier int `json:"tier,omitempty"`
			// Traverse angle, left (deg)
			TraverseLeftArc int `json:"traverse_left_arc,omitempty"`
			// Traverse angle, right (deg)
			TraverseRightArc int `json:"traverse_right_arc,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed int `json:"traverse_speed,omitempty"`
			// View range (m)
			ViewRange int `json:"view_range,omitempty"`
			// Weight (kg)
			Weight int `json:"weight,omitempty"`
		} `json:"turret,omitempty"`
	} `json:"default_profile,omitempty"`
	// Image links
	Images struct {
		// Normal image
		Normal string `json:"normal,omitempty"`
		// Small size image
		Preview string `json:"preview,omitempty"`
	} `json:"images,omitempty"`
	// Module research information
	ModulesTree struct {
		// Indicates if the module is basic
		IsDefault bool `json:"is_default,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Module name
		Name string `json:"name,omitempty"`
		// List of module IDs available after research of the module
		NextModules []int `json:"next_modules,omitempty"`
		// List of vehicle IDs available after research of the module
		NextTanks []int `json:"next_tanks,omitempty"`
		// Cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// Module type
		Type_ string `json:"type,omitempty"`
	} `json:"modules_tree,omitempty"`
}

// WotbEncyclopediaVehicles Method returns list of available vehicles.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/vehicles
//
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
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
// nation:
//     Nation. Maximum limit: 100.
// tank_id:
//     Vehicle ID. Maximum limit: 100.
func (client *Client) WotbEncyclopediaVehicles(realm Realm, fields []string, language string, nation []string, tankId []int) (*WotbEncyclopediaVehicles, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"nation": strings.Join(nation, ","),
		"tank_id": utils.SliceIntToString(tankId, ","),
	}

	var result *WotbEncyclopediaVehicles
	err := client.doGetDataRequest(realm, "/wotb/encyclopedia/vehicles/", reqParam, &result)
	return result, err
}
