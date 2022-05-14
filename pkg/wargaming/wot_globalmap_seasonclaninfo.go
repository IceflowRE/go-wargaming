package wargaming

import (
	"strings"
	"strconv"
)

type WotGlobalmapSeasonclaninfo struct {
	// Clan info by seasons and vehicle Tiers
	Seasons struct {
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Elo rating
		Elo int `json:"elo,omitempty"`
		// Current rating
		Rank int `json:"rank,omitempty"`
		// Rating changes during previous turn
		RankDelta int `json:"rank_delta,omitempty"`
		// Vehicle Tier
		VehicleLevel int `json:"vehicle_level,omitempty"`
		// Victory Points
		VictoryPoints int `json:"victory_points,omitempty"`
		// Change of Victory Points since last turn calculation
		VictoryPointsSinceTurn int `json:"victory_points_since_turn,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
	} `json:"seasons,omitempty"`
}

// WotGlobalmapSeasonclaninfo Method returns clan's statistics for a specific season.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasonclaninfo
//
// clan_id:
//     Clan ID. To get a clan ID, use the Clans method. Min value is 1.
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
func (client *Client) WotGlobalmapSeasonclaninfo(realm Realm, clanId int, seasonId string, vehicleLevel []string, fields []string) (*WotGlobalmapSeasonclaninfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": strconv.Itoa(clanId),
		"season_id": seasonId,
		"vehicle_level": strings.Join(vehicleLevel, ","),
		"fields": strings.Join(fields, ","),
	}

	var result *WotGlobalmapSeasonclaninfo
	err := client.doGetDataRequest(realm, "/wot/globalmap/seasonclaninfo/", reqParam, &result)
	return result, err
}
