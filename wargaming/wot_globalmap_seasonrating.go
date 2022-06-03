package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapSeasonrating Method returns season clan rating.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasonrating
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// seasonId:
//     Season ID. To get a season ID, use the Seasons method.
// vehicleLevel:
//     Vehicle Tier. Valid values:
//
//     "6" - Vehicles of Tier 6
//     "8" - Vehicles of Tier 8
//     "10" - Vehicles of Tier 10
func (service *WotService) GlobalmapSeasonrating(ctx context.Context, realm Realm, seasonId string, vehicleLevel string, options *wot.GlobalmapSeasonratingOptions) ([]*wot.GlobalmapSeasonrating, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"season_id":     seasonId,
		"vehicle_level": vehicleLevel,
	}
	if options != nil {
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data []*wot.GlobalmapSeasonrating
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/seasonrating/", reqParam, &data)
	return data, err
}
