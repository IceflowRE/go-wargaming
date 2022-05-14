package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotGlobalmapClaninfo struct {
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Clan name
	Name string `json:"name,omitempty"`
	// Clan tag
	Tag string `json:"tag,omitempty"`
	// Restricted clan information on the Global Map
	Private struct {
		// Influence points spent per day
		DailyWage int `json:"daily_wage,omitempty"`
		// Influence points
		Influence int `json:"influence,omitempty"`
	} `json:"private,omitempty"`
	// Clan rating on the Global Map
	Ratings struct {
		// Clan Elo rating in Absolute division
		Elo10 int `json:"elo_10,omitempty"`
		// Clan Elo rating in Medium division
		Elo6 int `json:"elo_6,omitempty"`
		// Clan Elo rating in Champion division
		Elo8 int `json:"elo_8,omitempty"`
		// Ratings update time
		UpdatedAt UnixTime `json:"updated_at,omitempty"`
	} `json:"ratings,omitempty"`
	// Clan statistics on the Global Map
	Statistics struct {
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Battles fought in Absolute division
		Battles10Level int `json:"battles_10_level,omitempty"`
		// Battles fought in Medium division
		Battles6Level int `json:"battles_6_level,omitempty"`
		// Battles fought in Champion division
		Battles8Level int `json:"battles_8_level,omitempty"`
		// Total number by provinces captured by clan
		Captures int `json:"captures,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Current number of captured provinces
		ProvincesCount int `json:"provinces_count,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Victories in Absolute division
		Wins10Level int `json:"wins_10_level,omitempty"`
		// Victories in Medium division
		Wins6Level int `json:"wins_6_level,omitempty"`
		// Victories in Champion division
		Wins8Level int `json:"wins_8_level,omitempty"`
	} `json:"statistics,omitempty"`
}

// WotGlobalmapClaninfo Method returns clan data on the Global Map.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/claninfo
//
// clan_id:
//     Clan ID. To get a clan ID, use the Clans method. Maximum limit: 10.
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WotGlobalmapClaninfo(realm Realm, clanId []int, accessToken string, fields []string) (*WotGlobalmapClaninfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": utils.SliceIntToString(clanId, ","),
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
	}

	var result *WotGlobalmapClaninfo
	err := client.doGetDataRequest(realm, "/wot/globalmap/claninfo/", reqParam, &result)
	return result, err
}
