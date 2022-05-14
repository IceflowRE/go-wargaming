package wargaming

import (
	"strings"
	"strconv"
)

type WotGlobalmapSeasonratingneighbors struct {
	// Award level
	AwardLevel string `json:"award_level,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Clan color
	Color string `json:"color,omitempty"`
	// Clan name
	Name string `json:"name,omitempty"`
	// Current rating
	Rank int `json:"rank,omitempty"`
	// Rating changes during previous turn
	RankDelta int `json:"rank_delta,omitempty"`
	// Clan tag
	Tag string `json:"tag,omitempty"`
	// Date of rating calculation
	UpdatedAt UnixTime `json:"updated_at,omitempty"`
	// Victory Points
	VictoryPoints int `json:"victory_points,omitempty"`
	// Victory Points to get the next award
	VictoryPointsToNextAward int `json:"victory_points_to_next_award,omitempty"`
}

// WotGlobalmapSeasonratingneighbors Method returns list of adjacent positions in season clan rating.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasonratingneighbors
//
// clan_id:
//     Clan ID. To get a clan ID, use the Clans method. Min value is 1.
// limit:
//     Neighbors limit. Default is 10. Min value is 1. Maximum value: 99.
// season_id:
//     Season ID. To get a season ID, use the Seasons method.
// vehicle_level:
//     Vehicle Tier. Valid values:
//     
//     "6" &mdash; Vehicles of Tier 6 
//     "8" &mdash; Vehicles of Tier 8 
//     "10" &mdash; Vehicles of Tier 10
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WotGlobalmapSeasonratingneighbors(realm Realm, clanId int, limit int, seasonId string, vehicleLevel string, fields []string) (*WotGlobalmapSeasonratingneighbors, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": strconv.Itoa(clanId),
		"limit": strconv.Itoa(limit),
		"season_id": seasonId,
		"vehicle_level": vehicleLevel,
		"fields": strings.Join(fields, ","),
	}

	var result *WotGlobalmapSeasonratingneighbors
	err := client.doGetDataRequest(realm, "/wot/globalmap/seasonratingneighbors/", reqParam, &result)
	return result, err
}
