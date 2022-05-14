package wargaming

import (
	"strconv"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotEncyclopediaVehicles struct {
	// Vehicle description
	Description string `json:"description,omitempty"`
	// List of compatible engine IDs
	Engines []int `json:"engines,omitempty"`
	// List of compatible gun IDs
	Guns []int `json:"guns,omitempty"`
	// Indicates if the vehicle is a gift vehicle
	IsGift bool `json:"is_gift,omitempty"`
	// Indicates if the vehicle is Premium vehicle
	IsPremium bool `json:"is_premium,omitempty"`
	// Indicates the IGR vehicle. Active only for Korea realm
	IsPremiumIgr bool `json:"is_premium_igr,omitempty"`
	// Indicates if the vehicle is a wheeled vehicle
	IsWheeled bool `json:"is_wheeled,omitempty"`
	// Vehicle name
	Name string `json:"name,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// List of vehicles available for research in form of pairs:
	// 
	// researched vehicle ID
	// cost of research in XP
	NextTanks map[string]string `json:"next_tanks,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// List of research costs in form of pairs:
	// 
	// parent vehicle ID
	// cost of research in XP
	PricesXp map[string]string `json:"prices_xp,omitempty"`
	// List of IDs of compatible equipment and consumables
	Provisions []int `json:"provisions,omitempty"`
	// List of compatible radio IDs
	Radios []int `json:"radios,omitempty"`
	// Vehicle short name
	ShortName string `json:"short_name,omitempty"`
	// List of compatible suspension IDs
	Suspensions []int `json:"suspensions,omitempty"`
	// Vehicle tag
	Tag string `json:"tag,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Tier
	Tier int `json:"tier,omitempty"`
	// List of compatible turret IDs
	Turrets []int `json:"turrets,omitempty"`
	// Vehicle type
	Type_ string `json:"type,omitempty"`
	// Crew
	Crew struct {
		// Crew member ID
		MemberId string `json:"member_id,omitempty"`
		// List of crew member roles
		Roles map[string]string `json:"roles,omitempty"`
	} `json:"crew,omitempty"`
	// Standard configuration characteristics
	DefaultProfile struct {
		// Hit points
		Hp int `json:"hp,omitempty"`
		// Hull HP
		HullHp int `json:"hull_hp,omitempty"`
		// Hull weight (kg)
		HullWeight int `json:"hull_weight,omitempty"`
		// Ammunition
		MaxAmmo int `json:"max_ammo,omitempty"`
		// Load limit (kg)
		MaxWeight int `json:"max_weight,omitempty"`
		// Top reverse speed (km/h)
		SpeedBackward int `json:"speed_backward,omitempty"`
		// Top speed (km/h)
		SpeedForward int `json:"speed_forward,omitempty"`
		// Weight (kg)
		Weight int `json:"weight,omitempty"`
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
				Duration []int `json:"duration,omitempty"`
			} `json:"stun,omitempty"`
		} `json:"ammo,omitempty"`
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
			// Module name
			Name string `json:"name,omitempty"`
			// Engine Power (hp)
			Power int `json:"power,omitempty"`
			// Module tag
			Tag string `json:"tag,omitempty"`
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
			// Dispersion at 100 m (m)
			Dispersion float32 `json:"dispersion,omitempty"`
			// Rate of fire (rounds/min)
			FireRate float32 `json:"fire_rate,omitempty"`
			// Depression angle (deg)
			MoveDownArc int `json:"move_down_arc,omitempty"`
			// Elevation angle (deg)
			MoveUpArc int `json:"move_up_arc,omitempty"`
			// Module name
			Name string `json:"name,omitempty"`
			// Reload time (s)
			ReloadTime float32 `json:"reload_time,omitempty"`
			// Module tag
			Tag string `json:"tag,omitempty"`
			// Tier
			Tier int `json:"tier,omitempty"`
			// Traverse speed (deg/s)
			TraverseSpeed int `json:"traverse_speed,omitempty"`
			// Weight (kg)
			Weight int `json:"weight,omitempty"`
		} `json:"gun,omitempty"`
		// Mounted modules
		Modules struct {
			// Engine ID
			EngineId int `json:"engine_id,omitempty"`
			// Gun ID
			GunId int `json:"gun_id,omitempty"`
			// Radio ID
			RadioId int `json:"radio_id,omitempty"`
			// Suspension ID
			SuspensionId int `json:"suspension_id,omitempty"`
			// Turret ID
			TurretId int `json:"turret_id,omitempty"`
		} `json:"modules,omitempty"`
		// Radio characteristics
		Radio struct {
			// Module name
			Name string `json:"name,omitempty"`
			// Signal range
			SignalRange int `json:"signal_range,omitempty"`
			// Module tag
			Tag string `json:"tag,omitempty"`
			// Tier
			Tier int `json:"tier,omitempty"`
			// Weight (kg)
			Weight int `json:"weight,omitempty"`
		} `json:"radio,omitempty"`
		// Vehicle characteristics in Rapid mode (for wheeled vehicles)
		Rapid struct {
			// Top reverse speed (km/h)
			SpeedBackward int `json:"speed_backward,omitempty"`
			// Top speed (km/h)
			SpeedForward int `json:"speed_forward,omitempty"`
			// Maximum wheel turning angle
			SuspensionSteeringLockAngle int `json:"suspension_steering_lock_angle,omitempty"`
			// Time required to switch to Cruise mode
			SwitchOffTime float32 `json:"switch_off_time,omitempty"`
			// Time required to switch to Rapid mode
			SwitchOnTime float32 `json:"switch_on_time,omitempty"`
		} `json:"rapid,omitempty"`
		// Vehicle characteristics in Siege mode
		Siege struct {
			// Aiming time (s)
			AimTime float32 `json:"aim_time,omitempty"`
			// Dispersion at 100 m (m)
			Dispersion float32 `json:"dispersion,omitempty"`
			// Depression angle (deg)
			MoveDownArc int `json:"move_down_arc,omitempty"`
			// Elevation angle (deg)
			MoveUpArc int `json:"move_up_arc,omitempty"`
			// Reload time (s)
			ReloadTime float32 `json:"reload_time,omitempty"`
			// Top reverse speed (km/h)
			SpeedBackward int `json:"speed_backward,omitempty"`
			// Standard suspension traverse speed
			SuspensionTraverseSpeed int `json:"suspension_traverse_speed,omitempty"`
			// Time needed to switch on the Siege mode
			SwitchOffTime float32 `json:"switch_off_time,omitempty"`
			// Time required to switch to Siege mode
			SwitchOnTime float32 `json:"switch_on_time,omitempty"`
		} `json:"siege,omitempty"`
		// Suspension characteristics
		Suspension struct {
			// Load limit (kg)
			LoadLimit int `json:"load_limit,omitempty"`
			// Module name
			Name string `json:"name,omitempty"`
			// Maximum wheel turning angle (for wheeled vehicles)
			SteeringLockAngle int `json:"steering_lock_angle,omitempty"`
			// Module tag
			Tag string `json:"tag,omitempty"`
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
			// Module name
			Name string `json:"name,omitempty"`
			// Module tag
			Tag string `json:"tag,omitempty"`
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
		// URL to 160 x 100 px image of vehicle
		BigIcon string `json:"big_icon,omitempty"`
		// URL to outline image of vehicle
		ContourIcon string `json:"contour_icon,omitempty"`
		// URL to 124 x 31 px image of vehicle
		SmallIcon string `json:"small_icon,omitempty"`
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
	// Информация об мультинации
	Multination struct {
		// Признак нации по-умолчанию
		IsDefault bool `json:"is_default,omitempty"`
		// Vehicle ID
		TankId int `json:"tank_id,omitempty"`
	} `json:"multination,omitempty"`
}

// WotEncyclopediaVehicles Method returns list of available vehicles.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/vehicles
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
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// nation:
//     Nation. Maximum limit: 100.
// page_no:
//     Result page number
// tank_id:
//     Vehicle ID. Maximum limit: 100.
// tier:
//     Tier. Maximum limit: 100.
// type:
//     Vehicle type. Maximum limit: 100. Valid values:
//     
//     "heavyTank" &mdash; Heavy Tank 
//     "AT-SPG" &mdash; Tank Destroyer 
//     "mediumTank" &mdash; Medium Tank 
//     "lightTank" &mdash; Light Tank 
//     "SPG" &mdash; SPG
func (client *Client) WotEncyclopediaVehicles(realm Realm, fields []string, language string, limit int, nation []string, pageNo int, tankId []int, tier []int, type_ []string) (*WotEncyclopediaVehicles, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"nation": strings.Join(nation, ","),
		"page_no": strconv.Itoa(pageNo),
		"tank_id": utils.SliceIntToString(tankId, ","),
		"tier": utils.SliceIntToString(tier, ","),
		"type": strings.Join(type_, ","),
	}

	var result *WotEncyclopediaVehicles
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/vehicles/", reqParam, &result)
	return result, err
}
