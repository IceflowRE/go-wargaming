package wargaming

import (
	"strings"
	"strconv"
)

type WotGlobalmapSeasonaccountinfo struct {
	// Account information by seasons and vehicle Tiers
	Seasons struct {
		// Account ID
		AccountId int `json:"account_id,omitempty"`
		// Award level
		AwardLevel string `json:"award_level,omitempty"`
		// Battles fought for current clan
		Battles int `json:"battles,omitempty"`
		// Battles to fight in a current clan to get clan award for the season
		BattlesToAward int `json:"battles_to_award,omitempty"`
		// Clan ID
		ClanId int `json:"clan_id,omitempty"`
		// Clan rating
		ClanRank int `json:"clan_rank,omitempty"`
		// Season ID
		SeasonId string `json:"season_id,omitempty"`
		// Date when account data were updated
		UpdatedAt UnixTime `json:"updated_at,omitempty"`
		// Vehicle Tier
		VehicleLevel int `json:"vehicle_level,omitempty"`
	} `json:"seasons,omitempty"`
}

// WotGlobalmapSeasonaccountinfo Method returns player's statistics for a specific season.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasonaccountinfo
//
// account_id:
//     Account ID. Min value is 1.
// season_id:
//     Season ID. To get a season ID, use the Seasons method.
// vehicle_level:
//     List of vehicle Tiers. Maximum limit: 100. Valid values:
//     
//     "6" &mdash; Vehicles of Tier 6 
//     "8" &mdash; Vehicles of Tier 8 
//     "10" &mdash; Vehicles of Tier 10
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WotGlobalmapSeasonaccountinfo(realm Realm, accountId int, seasonId string, vehicleLevel []string, fields []string) (*WotGlobalmapSeasonaccountinfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"season_id": seasonId,
		"vehicle_level": strings.Join(vehicleLevel, ","),
		"fields": strings.Join(fields, ","),
	}

	var result *WotGlobalmapSeasonaccountinfo
	err := client.doGetDataRequest(realm, "/wot/globalmap/seasonaccountinfo/", reqParam, &result)
	return result, err
}
