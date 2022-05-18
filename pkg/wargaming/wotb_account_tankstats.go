package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strconv"
	"strings"
)

// WotbAccountTankstats Method returns players' statistics on the vehicle.
//
// https://developers.wargaming.net/reference/all/wotb/account/tankstats
//
// account_id:
//     Player account ID. Maximum limit: 100.
// tank_id:
//     Player's vehicle ID
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
func (client *Client) WotbAccountTankstats(realm Realm, accountId []int, tankId int, fields []string, language string) (*wotb.AccountTankstats, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"tank_id": strconv.Itoa(tankId),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wotb.AccountTankstats
	err := client.doGetDataRequest(realm, "/wotb/account/tankstats/", reqParam, &result)
	return result, err
}
