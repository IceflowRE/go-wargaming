package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wows"
	"strconv"
	"strings"
)

// WowsSeasonsShipstats Method returns players' ships statistics in Ranked Battles seasons. Accounts with hidden game profiles are excluded from response. Hidden profiles are listed in the field meta.hidden.
//
// https://developers.wargaming.net/reference/all/wows/seasons/shipstats
//
// accountId:
//     Player account ID
//     Parameter is required.
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
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
// seasonId:
//     Season ID. Maximum limit: 100.
// shipId:
//     Ship ID. Maximum limit: 100.
func (client *Client) WowsSeasonsShipstats(realm Realm, accountId int, accessToken string, fields []string, language string, seasonId []int, shipId []int) (*wows.SeasonsShipstats, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
		"language": language,
		"season_id": utils.SliceIntToString(seasonId, ","),
		"ship_id": utils.SliceIntToString(shipId, ","),
	}

	var result *wows.SeasonsShipstats
	err := client.doGetDataRequest(realm, "/wows/seasons/shipstats/", reqParam, &result)
	return result, err
}
