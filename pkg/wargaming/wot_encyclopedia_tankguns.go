package wargaming

import (
	"strconv"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotEncyclopediaTankguns struct {
	// Damage
	Damage []int `json:"damage,omitempty"`
	// Tier
	Level int `json:"level,omitempty"`
	// Module ID
	ModuleId int `json:"module_id,omitempty"`
	// Vehicle name
	Name string `json:"name,omitempty"`
	// Localized name field
	NameI18N string `json:"name_i18n,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Localized nation field
	NationI18N string `json:"nation_i18n,omitempty"`
	// Penetration
	PiercingPower []int `json:"piercing_power,omitempty"`
	// Purchase cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Purchase cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Rate of fire
	Rate float32 `json:"rate,omitempty"`
	// IDs of compatible vehicles
	Tanks []int `json:"tanks,omitempty"`
	// IDs of compatible turrets
	Turrets []int `json:"turrets,omitempty"`
}

// WotEncyclopediaTankguns Deprecated: Attention! The method is deprecated.
// Method returns a tanks' gun list.
// The logic of this method and some field values may vary according to optional parameters passed.
// Changeable fields:
// 
// damage
// piercing_power
// rate
// price_credit
// price_gold
// 
// Optional input parameters work as follows:
// 
// correct turret_id passed — tank guns are filtered by whether they are placed on the turret and the abovementioned values change according to the turret;
// correct turret_id and module_id passed — the method returns details on each module with the abovementioned values changed according to the turret, or returns null if the module is not compatible with the turret;
// correct tank_id passed — if tank type matches one of AT-SPG, SPG, mediumTank, tank guns are filtered by whether they belong to the tank, the abovementioned values change according to the tank; otherwise, returns an error and requests turret_id. If module_id is also passed, the method returns details on each module with the abovementioned values changed according to the tank, or returns null if the module is not compatible with the tank;
// compatible turret_id and tank_id passed — tank guns are filtered by whether they belong to the tank and are placed on the turret, the abovementioned values changed according to the turret.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/tankguns
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
//     Module ID. Maximum limit: 1000.
// nation:
//     Nation. Maximum limit: 100.
// tank_id:
//     Compatible vehicle ID
// turret_id:
//     Compatible turret ID
func (client *Client) WotEncyclopediaTankguns(realm Realm, fields []string, language string, moduleId []int, nation []string, tankId int, turretId int) (*WotEncyclopediaTankguns, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"module_id": utils.SliceIntToString(moduleId, ","),
		"nation": strings.Join(nation, ","),
		"tank_id": strconv.Itoa(tankId),
		"turret_id": strconv.Itoa(turretId),
	}

	var result *WotEncyclopediaTankguns
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/tankguns/", reqParam, &result)
	return result, err
}
