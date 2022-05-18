package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapSeasonratingneighbors Method returns list of adjacent positions in season clan rating.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasonratingneighbors
//
// clanId:
//     Clan ID. To get a clan ID, use the Clans method. Min value is 1.
// limit:
//     Neighbors limit. Default is 10. Min value is 1. Maximum value: 99.
// seasonId:
//     Season ID. To get a season ID, use the Seasons method.
// vehicleLevel:
//     Vehicle Tier. Valid values:
//     
//     "6" &mdash; Vehicles of Tier 6 
//     "8" &mdash; Vehicles of Tier 8 
//     "10" &mdash; Vehicles of Tier 10
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WotGlobalmapSeasonratingneighbors(realm Realm, clanId int, limit int, seasonId string, vehicleLevel string, fields []string) (*wot.GlobalmapSeasonratingneighbors, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": strconv.Itoa(clanId),
		"limit": strconv.Itoa(limit),
		"season_id": seasonId,
		"vehicle_level": vehicleLevel,
		"fields": strings.Join(fields, ","),
	}

	var result *wot.GlobalmapSeasonratingneighbors
	err := client.doGetDataRequest(realm, "/wot/globalmap/seasonratingneighbors/", reqParam, &result)
	return result, err
}
