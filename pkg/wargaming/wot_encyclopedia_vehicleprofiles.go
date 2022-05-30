package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotEncyclopediaVehicleprofiles Method returns vehicle configuration characteristics.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/vehicleprofiles
//
// tankId:
//     Vehicle ID
//     Parameter is required.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
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
// orderBy:
//     Sorting. Valid values:
//     
//     "price_credit" - by cost in credits 
//     "-price_credit" - by cost in credits, in reverse order
func (client *Client) WotEncyclopediaVehicleprofiles(realm Realm, tankId int, fields []string, language string, orderBy string) (*wot.EncyclopediaVehicleprofiles, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tank_id": strconv.Itoa(tankId),
		"fields": strings.Join(fields, ","),
		"language": language,
		"order_by": orderBy,
	}

	var result *wot.EncyclopediaVehicleprofiles
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/vehicleprofiles/", reqParam, &result)
	return result, err
}
