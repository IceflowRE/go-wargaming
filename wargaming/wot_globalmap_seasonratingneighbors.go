package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapSeasonratingneighbors Method returns list of adjacent positions in season clan rating.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasonratingneighbors
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
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
func (service *WotService) GlobalmapSeasonratingneighbors(ctx context.Context, realm Realm, clanId int, seasonId string, vehicleLevel string, options *wot.GlobalmapSeasonratingneighborsOptions) ([]*wot.GlobalmapSeasonratingneighbors, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id":       strconv.Itoa(clanId),
		"season_id":     seasonId,
		"vehicle_level": vehicleLevel,
	}
	if options != nil {
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data []*wot.GlobalmapSeasonratingneighbors
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/seasonratingneighbors/", reqParam, &data)
	return data, err
}
