package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotx"
	"strings"
)

// WotxEncyclopediaVehicles Method returns list of available vehicles.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/vehicles
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
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
// nation:
//     Nation. Maximum limit: 100.
// tankId:
//     Vehicle ID. Maximum limit: 100.
// tier:
//     Tier. Maximum limit: 100.
func (client *Client) WotxEncyclopediaVehicles(realm Realm, fields []string, language string, nation []string, tankId []int, tier []int) (*wotx.EncyclopediaVehicles, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"nation": strings.Join(nation, ","),
		"tank_id": utils.SliceIntToString(tankId, ","),
		"tier": utils.SliceIntToString(tier, ","),
	}

	var result *wotx.EncyclopediaVehicles
	err := client.doGetDataRequest(realm, "/wotx/encyclopedia/vehicles/", reqParam, &result)
	return result, err
}
