// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapSeasonratingneighbors returns list of adjacent positions in season clan rating.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasonratingneighbors
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// clanId:
//     Clan ID. To get a clan ID, use the Clans method. Min value is 1.
// seasonId:
//     Season ID. To get a season ID, use the Seasons method.
// vehicleLevel:
//     Vehicle Tier. Valid values:
//
//     "6" - Vehicles of Tier 6
//     "8" - Vehicles of Tier 8
//     "10" - Vehicles of Tier 10
func (service *WotService) GlobalmapSeasonratingneighbors(ctx context.Context, realm Realm, clanId int, seasonId string, vehicleLevel string, options *wot.GlobalmapSeasonratingneighborsOptions) ([]*wot.GlobalmapSeasonratingneighbors, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"clan_id":       strconv.Itoa(clanId),
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
	}

	var data []*wot.GlobalmapSeasonratingneighbors
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/seasonratingneighbors/", reqParam, &data, &metaData)
	return data, metaData, err
}
