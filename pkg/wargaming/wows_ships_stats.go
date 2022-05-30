package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wows"
	"strconv"
	"strings"
)

// WowsShipsStats Method returns general statistics for each ship of a player. Accounts with hidden game profiles are excluded from response. Hidden profiles are listed in the field meta.hidden.
//
// https://developers.wargaming.net/reference/all/wows/ships/stats
//
// accountId:
//     Player account ID
//     Parameter is required.
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "club" 
//     "oper_div" 
//     "oper_div_hard" 
//     "oper_solo" 
//     "pve" 
//     "pve_div2" 
//     "pve_div3" 
//     "pve_solo" 
//     "pvp_div2" 
//     "pvp_div3" 
//     "pvp_solo" 
//     "rank_div2" 
//     "rank_div3" 
//     "rank_solo"
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// inGarage:
//     Filter by ship availability in the Port. If the parameter is not specified, all ships are returned. Parameter processing requires a valid access_token for the specified account_id. Valid values:
//     
//     "1" - Return ships available in the Port. 
//     "0" - Return ships that are no longer in the Port.
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
// shipId:
//     Player's ship ID. Maximum limit: 100.
func (client *Client) WowsShipsStats(realm Realm, accountId int, accessToken string, extra []string, fields []string, inGarage string, language string, shipId []int) (*wows.ShipsStats, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"access_token": accessToken,
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"in_garage": inGarage,
		"language": language,
		"ship_id": utils.SliceIntToString(shipId, ","),
	}

	var result *wows.ShipsStats
	err := client.doGetDataRequest(realm, "/wows/ships/stats/", reqParam, &result)
	return result, err
}
