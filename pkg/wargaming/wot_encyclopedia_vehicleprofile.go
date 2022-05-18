package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotEncyclopediaVehicleprofile Method returns vehicle configuration characteristics based on the specified module IDs.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/vehicleprofile
//
// engineId:
//     Engine ID. If module is not specified, standard module is used by default.
// tankId:
//     Vehicle ID
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// gunId:
//     Gun ID. If module is not specified, standard module is used by default.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "en" - English 
//     "ru" - Русский (by default)
//     "pl" - Polski 
//     "de" - Deutsch 
//     "fr" - Français 
//     "es" - Español 
//     "zh-cn" - 简体中文 
//     "zh-tw" - 繁體中文 
//     "tr" - Türkçe 
//     "cs" - Čeština 
//     "th" - ไทย 
//     "vi" - Tiếng Việt 
//     "ko" - 한국어
// profileId:
//     Configuration ID. If specified, parameters of IDs of separate modules are ignored.
// radioId:
//     Radio ID. If module is not specified, standard module is used by default.
// suspensionId:
//     Suspension ID. If module is not specified, standard module is used by default.
// turretId:
//     Turret ID. If module is not specified, standard module is used by default.
func (client *Client) WotEncyclopediaVehicleprofile(realm Realm, engineId int, tankId int, fields []string, gunId int, language string, profileId string, radioId int, suspensionId int, turretId int) (*wot.EncyclopediaVehicleprofile, error) {
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
		"radio_id": strconv.Itoa(radioId),
		"suspension_id": strconv.Itoa(suspensionId),
		"turret_id": strconv.Itoa(turretId),
	}

	var result *wot.EncyclopediaVehicleprofile
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/vehicleprofile/", reqParam, &result)
	return result, err
}
