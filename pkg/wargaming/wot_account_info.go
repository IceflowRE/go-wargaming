package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strings"
)

// WotAccountInfo Method returns player details.
//
// https://developers.wargaming.net/reference/all/wot/account/info
//
// account_id:
//     Player account ID. Maximum limit: 100.
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "private.boosters" 
//     "private.garage" 
//     "private.grouped_contacts" 
//     "private.personal_missions" 
//     "private.rented" 
//     "statistics.epic" 
//     "statistics.fallout" 
//     "statistics.globalmap_absolute" 
//     "statistics.globalmap_champion" 
//     "statistics.globalmap_middle" 
//     "statistics.random" 
//     "statistics.ranked_10x10" 
//     "statistics.ranked_15x15" 
//     "statistics.ranked_battles" 
//     "statistics.ranked_battles_current" 
//     "statistics.ranked_battles_previous" 
//     "statistics.ranked_season_1" 
//     "statistics.ranked_season_2" 
//     "statistics.ranked_season_3"
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
func (client *Client) WotAccountInfo(realm Realm, accountId []int, accessToken string, extra []string, fields []string, language string) (map[string]*wot.AccountInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"access_token": accessToken,
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result map[string]*wot.AccountInfo
	err := client.doGetDataRequest(realm, "/wot/account/info/", reqParam, &result)
	return result, err
}
