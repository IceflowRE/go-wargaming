package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
	"strconv"
)

type WowpEncyclopediaPlanespecification struct {
	// Slots and modules
	Slots struct {
		// ID of unique bind of slot and module
		BindId int `json:"bind_id,omitempty"`
		// Standard module
		IsDefault bool `json:"is_default,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Slot name
		SlotName string `json:"slot_name,omitempty"`
		// Localized slot_name field
		SlotNameI18N string `json:"slot_name_i18n,omitempty"`
	} `json:"slots,omitempty"`
	// Aircraft specifications
	Specification struct {
		// Average time to turn 360 deg
		AverageTurnTime float32 `json:"average_turn_time,omitempty"`
		// Controllability
		Controllability int `json:"controllability,omitempty"`
		// Maximum dive speed
		DiveSpeed int `json:"dive_speed,omitempty"`
		// Firepower
		Dps int `json:"dps,omitempty"`
		// Hit points
		Hp int `json:"hp,omitempty"`
		// Maneuverability
		Maneuverability int `json:"maneuverability,omitempty"`
		// Weight
		Mass float32 `json:"mass,omitempty"`
		// Top speed at best altitude
		MaxSpeed int `json:"max_speed,omitempty"`
		// Optimum altitude
		OptimalHeight int `json:"optimal_height,omitempty"`
		// Optimum airspeed
		OptimalManeuverSpeed int `json:"optimal_maneuver_speed,omitempty"`
		// Rate of climb
		RateOfClimbing float32 `json:"rate_of_climbing,omitempty"`
		// Rate of roll
		RollManeuverability int `json:"roll_maneuverability,omitempty"`
		// Top speed at sea level
		SpeedAtTheGround int `json:"speed_at_the_ground,omitempty"`
		// Airspeed
		SpeedFactor int `json:"speed_factor,omitempty"`
		// Stall speed
		StallSpeed int `json:"stall_speed,omitempty"`
	} `json:"specification,omitempty"`
}

// WowpEncyclopediaPlanespecification Method returns information from Encyclopedia about aircraft specifications depending on installed modules.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planespecification
//
// bind_id:
//     ID of unique bind of slot and module. Maximum limit: 100.
// module_id:
//     Module ID. Maximum limit: 100.
// plane_id:
//     Aircraft ID
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
func (client *Client) WowpEncyclopediaPlanespecification(realm Realm, bindId []int, moduleId []int, planeId int, fields []string, language string) (*WowpEncyclopediaPlanespecification, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"bind_id": utils.SliceIntToString(bindId, ","),
		"module_id": utils.SliceIntToString(moduleId, ","),
		"plane_id": strconv.Itoa(planeId),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowpEncyclopediaPlanespecification
	err := client.doGetDataRequest(realm, "/wowp/encyclopedia/planespecification/", reqParam, &result)
	return result, err
}
