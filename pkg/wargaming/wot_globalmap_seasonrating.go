package wargaming

import (
	"strconv"
	"strings"
)

type WotGlobalmapSeasonrating struct {
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

// WotGlobalmapSeasonrating Method returns season clan rating.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasonrating
//
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
// limit:
//     Clans limit. Default is 10. Min value is 1. Maximum value: 100.
// page_no:
//     Page number. Default is 1. Min value is 1.
func (client *Client) WotGlobalmapSeasonrating(realm Realm, seasonId string, vehicleLevel string, fields []string, limit int, pageNo int) (*WotGlobalmapSeasonrating, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"season_id": seasonId,
		"vehicle_level": vehicleLevel,
		"fields": strings.Join(fields, ","),
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
	}

	var result *WotGlobalmapSeasonrating
	err := client.doGetDataRequest(realm, "/wot/globalmap/seasonrating/", reqParam, &result)
	return result, err
}
