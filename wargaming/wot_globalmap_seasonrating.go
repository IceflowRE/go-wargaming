// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapSeasonrating returns season clan rating.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasonrating
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// seasonId:
//     Season ID. To get a season ID, use the Seasons method.
// vehicleLevel:
//     Vehicle Tier. Valid values:
//
//     "6" - Vehicles of Tier 6
//     "8" - Vehicles of Tier 8
//     "10" - Vehicles of Tier 10
func (service *WotService) GlobalmapSeasonrating(ctx context.Context, realm Realm, seasonId string, vehicleLevel string, options *wot.GlobalmapSeasonratingOptions) ([]*wot.GlobalmapSeasonrating, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"season_id":     seasonId,
		"vehicle_level": vehicleLevel,
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
	}

	var data []*wot.GlobalmapSeasonrating
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/seasonrating/", reqParam, &data, &metaData)
	return data, metaData, err
}
