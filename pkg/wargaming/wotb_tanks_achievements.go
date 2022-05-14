package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
	"strconv"
)

type WotbTanksAchievements struct {
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
	// Achievements earned
	Achievements map[string]string `json:"achievements,omitempty"`
	// Maximum values of achievement series
	MaxSeries map[string]string `json:"max_series,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
}

// WotbTanksAchievements Method returns list of player's achievements on all vehicles.
//
// https://developers.wargaming.net/reference/all/wotb/tanks/achievements
//
// account_id:
//     Player account ID
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// in_garage:
//     Filter by vehicle availability in the Garage. If the parameter is not specified, all vehicles are returned. Parameter processing requires a valid access_token for the specified account_id. Valid values:
//     
//     "1" &mdash; Return vehicles available in the Garage. 
//     "0" &mdash; Return vehicles that are no longer in the Garage.
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
// tank_id:
//     Player's vehicle ID. Maximum limit: 100.
func (client *Client) WotbTanksAchievements(realm Realm, accountId int, accessToken string, fields []string, inGarage string, language string, tankId []int) (*WotbTanksAchievements, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
		"in_garage": inGarage,
		"language": language,
		"tank_id": utils.SliceIntToString(tankId, ","),
	}

	var result *WotbTanksAchievements
	err := client.doGetDataRequest(realm, "/wotb/tanks/achievements/", reqParam, &result)
	return result, err
}
