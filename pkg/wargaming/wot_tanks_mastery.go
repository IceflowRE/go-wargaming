package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strings"
)

// WotTanksMastery The method returns percentiles of the distribution of average damage or experience values for each piece of equipment
//
// https://developers.wargaming.net/reference/all/wot/tanks/mastery
//
// distribution:
//     Type of data. Valid values:
//     
//     "damage" - Use damage distribution 
//     "xp" - Use a distribution based on experience
// percentile:
//     A list of percentiles to be included in the response. Maximum limit: 10. Min value is 1. Maximum value: 100.
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
// tankId:
//     Player's vehicle ID. Maximum limit: 100.
func (client *Client) WotTanksMastery(realm Realm, distribution string, percentile []int, fields []string, language string, tankId []int) (*wot.TanksMastery, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"distribution": distribution,
		"percentile": utils.SliceIntToString(percentile, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
		"tank_id": utils.SliceIntToString(tankId, ","),
	}

	var result *wot.TanksMastery
	err := client.doGetDataRequest(realm, "/wot/tanks/mastery/", reqParam, &result)
	return result, err
}
