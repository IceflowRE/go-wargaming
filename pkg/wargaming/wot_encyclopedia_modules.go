package wargaming

import (
	"strconv"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotEncyclopediaModules struct {
	// Image link
	Image string `json:"image,omitempty"`
	// Module ID
	ModuleId int `json:"module_id,omitempty"`
	// Module name
	Name string `json:"name,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Vehicles compatible with module
	Tanks []int `json:"tanks,omitempty"`
	// Tier
	Tier int `json:"tier,omitempty"`
	// Module type
	Type_ string `json:"type,omitempty"`
	// Weight (kg)
	Weight int `json:"weight,omitempty"`
	// Basic technical characteristics of module.
	// An extra field.
	DefaultProfile struct {
		// Engine characteristics
		Engine struct {
			// Chance of engine fire
			FireChance float32 `json:"fire_chance,omitempty"`
			// Engine Power (hp)
			Power int `json:"power,omitempty"`
		} `json:"engine,omitempty"`
		// Gun characteristics
		Gun struct {
			// Aiming time (s)
			AimTime float32 `json:"aim_time,omitempty"`
			// Dispersion at 100 m (m)
			Dispersion float32 `json:"dispersion,omitempty"`
			// Rate of fire (rounds/min)
			FireRate float32 `json:"fire_rate,omitempty"`
			// Number of shells
			MaxAmmo int `json:"max_ammo,omitempty"`
			// Depression angle (deg)
			MoveDownArc int `json:"move_down_arc,omitempty"`
			// Elevation angle (deg)
			MoveUpArc int `json:"move_up_arc,omitempty"`
			// Reload time (s)
			ReloadTime float32 `json:"reload_time,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed int `json:"traverse_speed,omitempty"`
			// Gun shells characteristics
			Ammo struct {
				// Damage (hp), a list of values: min, avg, max
				Damage []int `json:"damage,omitempty"`
				// Penetration (mm), a list of values: min, avg, max
				Penetration []int `json:"penetration,omitempty"`
				// Shell type
				Type_ string `json:"type,omitempty"`
				// Stun characteristics
				Stun struct {
					// Stun duration (s) caused by this shell type, a list of values: min, max
					Duration struct {
					} `json:"duration,omitempty"`
				} `json:"stun,omitempty"`
			} `json:"ammo,omitempty"`
		} `json:"gun,omitempty"`
		// Radio characteristics
		Radio struct {
			// Signal range
			SignalRange int `json:"signal_range,omitempty"`
		} `json:"radio,omitempty"`
		// Suspension characteristics
		Suspension struct {
			// Load limit (kg)
			LoadLimit int `json:"load_limit,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed int `json:"traverse_speed,omitempty"`
		} `json:"suspension,omitempty"`
		// Turret characteristics
		Turret struct {
			// Armor: front (mm)
			ArmorFront int `json:"armor_front,omitempty"`
			// Armor: rear (mm)
			ArmorRear int `json:"armor_rear,omitempty"`
			// Armor: sides (mm)
			ArmorSides int `json:"armor_sides,omitempty"`
			// Hit points
			Hp int `json:"hp,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed int `json:"traverse_speed,omitempty"`
			// View range (m)
			ViewRange int `json:"view_range,omitempty"`
		} `json:"turret,omitempty"`
	} `json:"default_profile,omitempty"`
}

// WotEncyclopediaModules Method returns list of available modules that can be installed on vehicles, such as engines, turrets, etc. At least one input filter parameter (module ID, type) is required to be indicated.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/modules
//
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "default_profile"
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
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// module_id:
//     Module ID. Maximum limit: 100.
// nation:
//     Nation. Maximum limit: 100.
// page_no:
//     Result page number
// type:
//     Module type. Maximum limit: 100. Valid values:
//     
//     "vehicleRadio" &mdash; Radio 
//     "vehicleEngine" &mdash; Engines 
//     "vehicleGun" &mdash; Gun 
//     "vehicleChassis" &mdash; Suspension 
//     "vehicleTurret" &mdash; Turret
func (client *Client) WotEncyclopediaModules(realm Realm, extra []string, fields []string, language string, limit int, moduleId []int, nation []string, pageNo int, type_ []string) (*WotEncyclopediaModules, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"module_id": utils.SliceIntToString(moduleId, ","),
		"nation": strings.Join(nation, ","),
		"page_no": strconv.Itoa(pageNo),
		"type": strings.Join(type_, ","),
	}

	var result *WotEncyclopediaModules
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/modules/", reqParam, &result)
	return result, err
}
