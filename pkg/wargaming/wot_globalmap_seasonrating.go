package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapSeasonrating Method returns season clan rating.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasonrating
//
// seasonId:
//     Season ID. To get a season ID, use the Seasons method.
// vehicleLevel:
//     Vehicle Tier. Valid values:
//     
//     "6" - Vehicles of Tier 6 
//     "8" - Vehicles of Tier 8 
//     "10" - Vehicles of Tier 10
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// limit:
//     Clans limit. Default is 10. Min value is 1. Maximum value: 100.
// pageNo:
//     Page number. Default is 1. Min value is 1.
func (client *Client) WotGlobalmapSeasonrating(realm Realm, seasonId string, vehicleLevel string, fields []string, limit int, pageNo int) (*wot.GlobalmapSeasonrating, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"season_id": seasonId,
		"vehicle_level": vehicleLevel,
		"fields": strings.Join(fields, ","),
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
	}

	var result *wot.GlobalmapSeasonrating
	err := client.doGetDataRequest(realm, "/wot/globalmap/seasonrating/", reqParam, &result)
	return result, err
}
