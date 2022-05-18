package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

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
// pageNo:
//     Result page number
// tankId:
//     Vehicle ID. Maximum limit: 100.
// tier:
//     Tier. Maximum limit: 100.
// type_:
//     Vehicle type. Maximum limit: 100. Valid values:
//     
//     "heavyTank" &mdash; Heavy Tank 
//     "AT-SPG" &mdash; Tank Destroyer 
//     "mediumTank" &mdash; Medium Tank 
//     "lightTank" &mdash; Light Tank 
//     "SPG" &mdash; SPG
func (client *Client) WotEncyclopediaVehicles(realm Realm, fields []string, language string, limit int, nation []string, pageNo int, tankId []int, tier []int, type_ []string) (*wot.EncyclopediaVehicles, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
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

	var result *wot.EncyclopediaVehicles
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/vehicles/", reqParam, &result)
	return result, err
}
