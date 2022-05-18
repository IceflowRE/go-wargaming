package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wows"
	"strings"
)

// WowsAccountInfo Method returns player details. Players may hide their game profiles, use field hidden_profile for determination.
//
// https://developers.wargaming.net/reference/all/wows/account/info
//
// accountId:
//     Player account ID. Maximum limit: 100. Min value is 1.
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "private.grouped_contacts" 
//     "private.port" 
//     "statistics.club" 
//     "statistics.oper_div" 
//     "statistics.oper_div_hard" 
//     "statistics.oper_solo" 
//     "statistics.pve" 
//     "statistics.pve_div2" 
//     "statistics.pve_div3" 
//     "statistics.pve_solo" 
//     "statistics.pvp_div2" 
//     "statistics.pvp_div3" 
//     "statistics.pvp_solo" 
//     "statistics.rank_div2" 
//     "statistics.rank_div3" 
//     "statistics.rank_solo"
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "cs" &mdash; Čeština 
//     "de" &mdash; Deutsch 
//     "en" &mdash; English 
//     "es" &mdash; Español 
//     "fr" &mdash; Français 
//     "ja" &mdash; 日本語 
//     "pl" &mdash; Polski 
//     "ru" &mdash; Русский (by default)
//     "th" &mdash; ไทย 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "zh-cn" &mdash; 中文 
//     "pt-br" &mdash; Português do Brasil 
//     "es-mx" &mdash; Español (México)
func (client *Client) WowsAccountInfo(realm Realm, accountId []int, accessToken string, extra []string, fields []string, language string) (map[string]*wows.AccountInfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"access_token": accessToken,
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result map[string]*wows.AccountInfo
	err := client.doGetDataRequest(realm, "/wows/account/info/", reqParam, &result)
	return result, err
}
