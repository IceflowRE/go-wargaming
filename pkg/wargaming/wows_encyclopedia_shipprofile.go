package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wows"
	"strconv"
	"strings"
)

// WowsEncyclopediaShipprofile Method returns parameters of ships in all existing configurations.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/shipprofile
//
// artilleryId:
//     Main Battery ID. If the module is not indicated, module of basic configuration is used.
// diveBomberId:
//     Dive bombers' ID. If the module is not indicated, module of basic configuration is used.
// engineId:
//     Engine ID. If the module is not indicated, module of basic configuration is used.
// fighterId:
//     Fighters' ID. If the module is not indicated, module of basic configuration is used.
// shipId:
//     Ship ID
//     Parameter is required.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// fireControlId:
//     ID of Gun Fire Control System. If the module is not indicated, module of basic configuration is used.
// flightControlId:
//     ID of Flight Control System. If the module is not indicated, module of basic configuration is used.
// hullId:
//     Hull ID. If the module is not indicated, module of basic configuration is used.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "cs" - Čeština 
//     "de" - Deutsch 
//     "en" - English 
//     "es" - Español 
//     "fr" - Français 
//     "ja" - 日本語 
//     "pl" - Polski 
//     "ru" - Русский (by default)
//     "th" - ไทย 
//     "zh-tw" - 繁體中文 
//     "tr" - Türkçe 
//     "zh-cn" - 中文 
//     "pt-br" - Português do Brasil 
//     "es-mx" - Español (México)
// torpedoBomberId:
//     Torpedo bombers' ID. If the module is not indicated, module of basic configuration is used.
// torpedoesId:
//     Torpedo tubes' ID. If the module is not indicated, module of basic configuration is used.
func (client *Client) WowsEncyclopediaShipprofile(realm Realm, artilleryId int, diveBomberId int, engineId int, fighterId int, shipId int, fields []string, fireControlId int, flightControlId int, hullId int, language string, torpedoBomberId int, torpedoesId int) (*wows.EncyclopediaShipprofile, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"artillery_id": strconv.Itoa(artilleryId),
		"dive_bomber_id": strconv.Itoa(diveBomberId),
		"engine_id": strconv.Itoa(engineId),
		"fighter_id": strconv.Itoa(fighterId),
		"ship_id": strconv.Itoa(shipId),
		"fields": strings.Join(fields, ","),
		"fire_control_id": strconv.Itoa(fireControlId),
		"flight_control_id": strconv.Itoa(flightControlId),
		"hull_id": strconv.Itoa(hullId),
		"language": language,
		"torpedo_bomber_id": strconv.Itoa(torpedoBomberId),
		"torpedoes_id": strconv.Itoa(torpedoesId),
	}

	var result *wows.EncyclopediaShipprofile
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/shipprofile/", reqParam, &result)
	return result, err
}
