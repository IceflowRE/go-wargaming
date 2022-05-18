package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotx"
	"strconv"
	"strings"
)

// WotxEncyclopediaVehicleprofile Method returns vehicle configuration characteristics based on the specified module IDs.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/vehicleprofile
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
//     Localization language. Default is "en". Valid values:
//     
//     "en" &mdash; English (by default)
//     "ru" &mdash; Русский 
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "tr" &mdash; Türkçe
// profile_id:
//     Configuration ID. If specified, parameters of IDs of separate modules are ignored.
// radio_id:
//     Radio ID. If module is not specified, standard module is used by default.
// suspension_id:
//     Suspension ID. If module is not specified, standard module is used by default.
// turret_id:
//     Turret ID. If module is not specified, standard module is used by default.
func (client *Client) WotxEncyclopediaVehicleprofile(realm Realm, engineId int, tankId int, fields []string, gunId int, language string, profileId string, radioId int, suspensionId int, turretId int) (*wotx.EncyclopediaVehicleprofile, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
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

	var result *wotx.EncyclopediaVehicleprofile
	err := client.doGetDataRequest(realm, "/wotx/encyclopedia/vehicleprofile/", reqParam, &result)
	return result, err
}
