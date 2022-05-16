package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wowp"
	"strconv"
	"strings"
)

// WowpPlanesStats Method returns statistics on player's aircraft.
//
// https://developers.wargaming.net/reference/all/wowp/planes/stats
//
// account_id:
//     Player account ID
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// in_garage:
//     Filter by aircraft availability in the Hangar. If the parameter is not specified, all aircraft are returned. Parameter processing requires a valid access_token for the specified account_id. Valid values:
//     
//     "1" &mdash; Return aircraft available in the Hangar. 
//     "0" &mdash; Return aircraft that are no longer available in the Hangar.
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
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
// plane_id:
//     Aircraft ID. Maximum limit: 100.
func (client *Client) WowpPlanesStats(realm Realm, accountId int, accessToken string, fields []string, inGarage string, language string, planeId []int) (*wowp.PlanesStats, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
		"in_garage": inGarage,
		"language": language,
		"plane_id": utils.SliceIntToString(planeId, ","),
	}

	var result *wowp.PlanesStats
	err := client.doGetDataRequest(realm, "/wowp/planes/stats/", reqParam, &result)
	return result, err
}
