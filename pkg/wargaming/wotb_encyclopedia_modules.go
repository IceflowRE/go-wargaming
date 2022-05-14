package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotbEncyclopediaModules struct {
	// Engine characteristics
	Engines struct {
		// Chance of engine fire
		FireChance float32 `json:"fire_chance,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Vehicle name
		Name string `json:"name,omitempty"`
		// Nation
		Nation string `json:"nation,omitempty"`
		// Engine Power (hp)
		Power int `json:"power,omitempty"`
		// List of compatible vehicles
		Tanks []int `json:"tanks,omitempty"`
		// Tier
		Tier int `json:"tier,omitempty"`
		// Weight (kg)
		Weight int `json:"weight,omitempty"`
	} `json:"engines,omitempty"`
	// Gun characteristics
	Guns struct {
		// Aiming time (s)
		AimTime float32 `json:"aim_time,omitempty"`
		// Dispersion at 100 m (m)
		Dispersion float32 `json:"dispersion,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Vehicle name
		Name string `json:"name,omitempty"`
		// Nation
		Nation string `json:"nation,omitempty"`
		// List of compatible vehicles
		Tanks []int `json:"tanks,omitempty"`
		// Tier
		Tier int `json:"tier,omitempty"`
		// Weight (kg)
		Weight int `json:"weight,omitempty"`
		// Gun shells characteristics
		Shells struct {
			// Average damage (HP)
			Damage int `json:"damage,omitempty"`
			// Average penetration (mm)
			Penetration int `json:"penetration,omitempty"`
			// Type
			Type_ string `json:"type,omitempty"`
		} `json:"shells,omitempty"`
	} `json:"guns,omitempty"`
	// Suspension characteristics
	Suspensions struct {
		// Load limit (kg)
		LoadLimit int `json:"load_limit,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Vehicle name
		Name string `json:"name,omitempty"`
		// Nation
		Nation string `json:"nation,omitempty"`
		// List of compatible vehicles
		Tanks []int `json:"tanks,omitempty"`
		// Tier
		Tier int `json:"tier,omitempty"`
		// Traverse speed (deg/s)
		TraverseSpeed int `json:"traverse_speed,omitempty"`
		// Weight (kg)
		Weight int `json:"weight,omitempty"`
	} `json:"suspensions,omitempty"`
	// Turret characteristics
	Turrets struct {
		// Hit points
		Hp int `json:"hp,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Vehicle name
		Name string `json:"name,omitempty"`
		// Nation
		Nation string `json:"nation,omitempty"`
		// List of compatible vehicles
		Tanks []int `json:"tanks,omitempty"`
		// Tier
		Tier int `json:"tier,omitempty"`
		// Traverse angle, left (deg)
		TraverseLeftArc int `json:"traverse_left_arc,omitempty"`
		// Traverse angle, right (deg)
		TraverseRightArc int `json:"traverse_right_arc,omitempty"`
		// View range (m)
		ViewRange int `json:"view_range,omitempty"`
		// Weight (kg)
		Weight int `json:"weight,omitempty"`
		// Armor
		Armor struct {
			// Front (mm)
			Front int `json:"front,omitempty"`
			// Rear (mm)
			Rear int `json:"rear,omitempty"`
			// Sides (mm)
			Sides int `json:"sides,omitempty"`
		} `json:"armor,omitempty"`
	} `json:"turrets,omitempty"`
}

// WotbEncyclopediaModules Method returns list of available modules, such as guns, engines, etc.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/modules
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
// module_id:
//     Module ID. Maximum limit: 100.
// nation:
//     Nation
// type:
//     Module type. Valid values:
//     
//     "vehicleEngine" &mdash; Engines 
//     "vehicleGun" &mdash; Gun 
//     "vehicleChassis" &mdash; Suspension 
//     "vehicleTurret" &mdash; Turret
func (client *Client) WotbEncyclopediaModules(realm Realm, fields []string, language string, moduleId []int, nation string, type_ string) (*WotbEncyclopediaModules, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"module_id": utils.SliceIntToString(moduleId, ","),
		"nation": nation,
		"type": type_,
	}

	var result *WotbEncyclopediaModules
	err := client.doGetDataRequest(realm, "/wotb/encyclopedia/modules/", reqParam, &result)
	return result, err
}
