package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wows"
	"strconv"
	"strings"
)

// WowsAccountStatsbydate Method returns statistics slices by dates in specified time span.
//
// https://developers.wargaming.net/reference/all/wows/account/statsbydate
//
// accountId:
//     Player account ID
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// dates:
//     List of dates to return statistics slices for. Format: YYYYMMDD. Max. dates range - 28 days from the current date. Statistics slice for yesterday will be returned by default. Maximum limit: 10.
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "pve"
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
func (client *Client) WowsAccountStatsbydate(realm Realm, accountId int, accessToken string, dates []string, extra []string, fields []string, language string) (*wows.AccountStatsbydate, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"access_token": accessToken,
		"dates": strings.Join(dates, ","),
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wows.AccountStatsbydate
	err := client.doGetDataRequest(realm, "/wows/account/statsbydate/", reqParam, &result)
	return result, err
}
