package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotx"
	"strings"
)

// WotxEncyclopediaVehicleupgrades Method returns list of vehicle equipment and consumables.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/vehicleupgrades
//
// tankId:
//     Vehicle ID. Maximum limit: 100.
//     Parameter is required.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "en". Valid values:
//     
//     "en" - English (by default)
//     "ru" - Русский 
//     "pl" - Polski 
//     "de" - Deutsch 
//     "fr" - Français 
//     "es" - Español 
//     "tr" - Türkçe
func (client *Client) WotxEncyclopediaVehicleupgrades(realm Realm, tankId []int, fields []string, language string) (*wotx.EncyclopediaVehicleupgrades, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tank_id": utils.SliceIntToString(tankId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wotx.EncyclopediaVehicleupgrades
	err := client.doGetDataRequest(realm, "/wotx/encyclopedia/vehicleupgrades/", reqParam, &result)
	return result, err
}
