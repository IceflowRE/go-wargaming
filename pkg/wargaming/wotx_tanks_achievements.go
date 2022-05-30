package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotx"
	"strconv"
	"strings"
)

// WotxTanksAchievements Method returns players' achievement details.
// Achievement properties define the achievements field values
// 
// 1-4 for Mastery Badges and Stage Achievements (type: "class");
// maximum value of Achievement series (type: "series");
// number of achievements earned from sections: Battle Hero, Epic Achievements, Group Achievements, Special Achievements, etc. (type: "repeatable, single, custom").
//
// https://developers.wargaming.net/reference/all/wotx/tanks/achievements
//
// accountId:
//     Player account ID. Min value is 0.
//     Parameter is required.
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// inGarage:
//     Filter by vehicle availability in the Garage. If the parameter is not specified, all vehicles are returned. Parameter processing requires a valid access_token for the specified account_id. Valid values:
//     
//     "1" - Return vehicles available in the Garage. 
//     "0" - Return vehicles that are no longer in the Garage.
// language:
//     Localization language. Default is "en". Valid values:
//     
//     "en" - English (by default)
//     "ru" - Русский 
//     "pl" - Polski 
//     "de" - Deutsch 
//     "fr" - Français 
//     "es" - Español 
//     "tr" - Türkçe
// tankId:
//     Player's vehicle ID. Maximum limit: 100.
func (client *Client) WotxTanksAchievements(realm Realm, accountId int, accessToken string, fields []string, inGarage string, language string, tankId []int) (*wotx.TanksAchievements, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
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

	var result *wotx.TanksAchievements
	err := client.doGetDataRequest(realm, "/wotx/tanks/achievements/", reqParam, &result)
	return result, err
}
