package wargaming

import (
	"strconv"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WowsEncyclopediaModules struct {
	// Image link
	Image string `json:"image,omitempty"`
	// Module ID
	ModuleId int `json:"module_id,omitempty"`
	// Module string ID
	ModuleIdStr string `json:"module_id_str,omitempty"`
	// Module name
	Name string `json:"name,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Tag
	Tag string `json:"tag,omitempty"`
	// Module type
	Type_ string `json:"type,omitempty"`
	// Module parameters, values related to the module type
	Profile struct {
		// Main battery
		Artillery struct {
			// Rate of fire (rounds / min)
			GunRate float32 `json:"gun_rate,omitempty"`
			// Maximum Damage caused by Armor Piercing Shells
			MaxDamageAp int `json:"max_damage_AP,omitempty"`
			// Maximum Damage caused by High Explosive Shells
			MaxDamageHe int `json:"max_damage_HE,omitempty"`
			// 180 Degree Turn Time (sec)
			RotationTime float32 `json:"rotation_time,omitempty"`
		} `json:"artillery,omitempty"`
		// Dive bombers
		DiveBomber struct {
			// Chance of Fire on target caused by bomb (%)
			BombBurnProbability float32 `json:"bomb_burn_probability,omitempty"`
			// Cruise Speed (knots)
			CruiseSpeed int `json:"cruise_speed,omitempty"`
			// Maximum Bomb Damage
			MaxDamage int `json:"max_damage,omitempty"`
			// Survivability
			MaxHealth int `json:"max_health,omitempty"`
			// Accuracy
			Accuracy struct {
				// Maximum value
				Max float32 `json:"max,omitempty"`
				// Minimum value
				Min float32 `json:"min,omitempty"`
			} `json:"accuracy,omitempty"`
		} `json:"dive_bomber,omitempty"`
		// Engine
		Engine struct {
			// Top Speed (knots)
			MaxSpeed float32 `json:"max_speed,omitempty"`
		} `json:"engine,omitempty"`
		// Fighters
		Fighter struct {
			// Average damage caused per second
			AvgDamage int `json:"avg_damage,omitempty"`
			// Cruise Speed (knots)
			CruiseSpeed int `json:"cruise_speed,omitempty"`
			// Ammunition
			MaxAmmo int `json:"max_ammo,omitempty"`
			// Survivability
			MaxHealth int `json:"max_health,omitempty"`
		} `json:"fighter,omitempty"`
		// Gun Fire Control System
		FireControl struct {
			// Firing range
			Distance float32 `json:"distance,omitempty"`
			// Firing Range extension (%)
			DistanceIncrease int `json:"distance_increase,omitempty"`
		} `json:"fire_control,omitempty"`
		// Flight Control
		FlightControl struct {
			// Number of bomber squadrons
			BomberSquadrons int `json:"bomber_squadrons,omitempty"`
			// Number of fighter squadrons
			FighterSquadrons int `json:"fighter_squadrons,omitempty"`
			// Number of torpedo bomber squadrons
			TorpedoSquadrons int `json:"torpedo_squadrons,omitempty"`
		} `json:"flight_control,omitempty"`
		// Hull
		Hull struct {
			// AA Mounts
			AntiAircraftBarrels int `json:"anti_aircraft_barrels,omitempty"`
			// Number of main turrets
			ArtilleryBarrels int `json:"artillery_barrels,omitempty"`
			// Secondary gun turrets
			AtbaBarrels int `json:"atba_barrels,omitempty"`
			// Hit points
			Health int `json:"health,omitempty"`
			// Hangar capacity
			PlanesAmount int `json:"planes_amount,omitempty"`
			// Torpedo tubes
			TorpedoesBarrels int `json:"torpedoes_barrels,omitempty"`
			// Armor (mm)
			Range struct {
				// Maximum value
				Max int `json:"max,omitempty"`
				// Minimum value
				Min int `json:"min,omitempty"`
			} `json:"range,omitempty"`
		} `json:"hull,omitempty"`
		// Torpedo Bombers
		TorpedoBomber struct {
			// Cruise Speed (knots)
			CruiseSpeed int `json:"cruise_speed,omitempty"`
			// Firing range
			Distance float32 `json:"distance,omitempty"`
			// Maximum Bomb Damage
			MaxDamage int `json:"max_damage,omitempty"`
			// Survivability
			MaxHealth int `json:"max_health,omitempty"`
			// Maximum torpedo damage
			TorpedoDamage int `json:"torpedo_damage,omitempty"`
			// Top Speed (knots)
			TorpedoMaxSpeed int `json:"torpedo_max_speed,omitempty"`
			// Torpedo name
			TorpedoName string `json:"torpedo_name,omitempty"`
		} `json:"torpedo_bomber,omitempty"`
		// Torpedo tubes
		Torpedoes struct {
			// Firing range
			Distance float32 `json:"distance,omitempty"`
			// Maximum Damage
			MaxDamage int `json:"max_damage,omitempty"`
			// Reload Time (sec)
			ShotSpeed float32 `json:"shot_speed,omitempty"`
			// Torpedo Speed (knots)
			TorpedoSpeed int `json:"torpedo_speed,omitempty"`
		} `json:"torpedoes,omitempty"`
	} `json:"profile,omitempty"`
}

// WowsEncyclopediaModules Method returns list of available modules that can be mounted on a ship (hull, engines, etc.).
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/modules
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "cs" &mdash; Čeština 
//     "de" &mdash; Deutsch 
//     "en" &mdash; English 
//     "es" &mdash; Español 
//     "fr" &mdash; Français 
//     "ja" &mdash; 日本語 
//     "pl" &mdash; Polski 
//     "ru" &mdash; Русский (by default)
//     "th" &mdash; ไทย 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "zh-cn" &mdash; 中文 
//     "pt-br" &mdash; Português do Brasil 
//     "es-mx" &mdash; Español (México)
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// module_id:
//     Module ID. Maximum limit: 100.
// page_no:
//     Result page number. Default is 1. Min value is 1.
// type:
//     Module type. Valid values:
//     
//     "Artillery" &mdash; Main battery 
//     "Torpedoes" &mdash; Torpedo tubes 
//     "Suo" &mdash; Gun Fire Control System 
//     "FlightControl" &mdash; Flight Control 
//     "Hull" &mdash; Hull 
//     "Engine" &mdash; Engine 
//     "Fighter" &mdash; Fighters 
//     "TorpedoBomber" &mdash; Torpedo Bombers 
//     "DiveBomber" &mdash; Dive bombers
func (client *Client) WowsEncyclopediaModules(realm Realm, fields []string, language string, limit int, moduleId []int, pageNo int, type_ string) (*WowsEncyclopediaModules, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"module_id": utils.SliceIntToString(moduleId, ","),
		"page_no": strconv.Itoa(pageNo),
		"type": type_,
	}

	var result *WowsEncyclopediaModules
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/modules/", reqParam, &result)
	return result, err
}
