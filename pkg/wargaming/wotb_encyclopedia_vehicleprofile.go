package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strconv"
	"strings"
)

// WotbEncyclopediaVehicleprofile Method returns vehicle configuration characteristics based on the specified module IDs.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/vehicleprofile
//
// engine_id:
//     Engine ID. If module is not specified, standard module is used by default.
// tank_id:
//     Vehicle ID
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// gun_id:
//     Gun ID. If module is not specified, standard module is used by default.
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
// profile_id:
//     Configuration ID. If specified, parameters of IDs of separate modules are ignored.
// suspension_id:
//     Suspension ID. If module is not specified, standard module is used by default.
// turret_id:
//     Turret ID. If module is not specified, standard module is used by default.
func (client *Client) WotbEncyclopediaVehicleprofile(realm Realm, engineId int, tankId int, fields []string, gunId int, language string, profileId string, suspensionId int, turretId int) (*wotb.EncyclopediaVehicleprofile, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"engine_id": strconv.Itoa(engineId),
		"tank_id": strconv.Itoa(tankId),
		"fields": strings.Join(fields, ","),
		"gun_id": strconv.Itoa(gunId),
		"language": language,
		"profile_id": profileId,
		"suspension_id": strconv.Itoa(suspensionId),
		"turret_id": strconv.Itoa(turretId),
	}

	var result *wotb.EncyclopediaVehicleprofile
	err := client.doGetDataRequest(realm, "/wotb/encyclopedia/vehicleprofile/", reqParam, &result)
	return result, err
}
