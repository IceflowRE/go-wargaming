package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapSeasonaccountinfo Method returns player's statistics for a specific season.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasonaccountinfo
//
// accountId:
//     Account ID. Min value is 1.
//     Parameter is required.
// seasonId:
//     Season ID. To get a season ID, use the Seasons method.
//     Parameter is required.
// vehicleLevel:
//     List of vehicle Tiers. Maximum limit: 100. Valid values:
//     
//     "6" - Vehicles of Tier 6 
//     "8" - Vehicles of Tier 8 
//     "10" - Vehicles of Tier 10
//     Parameter is required.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WotGlobalmapSeasonaccountinfo(realm Realm, accountId int, seasonId string, vehicleLevel []string, fields []string) (*wot.GlobalmapSeasonaccountinfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"season_id": seasonId,
		"vehicle_level": strings.Join(vehicleLevel, ","),
		"fields": strings.Join(fields, ","),
	}

	var result *wot.GlobalmapSeasonaccountinfo
	err := client.doGetDataRequest(realm, "/wot/globalmap/seasonaccountinfo/", reqParam, &result)
	return result, err
}