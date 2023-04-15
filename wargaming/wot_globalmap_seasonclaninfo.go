// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapSeasonclaninfo returns clan's statistics for a specific season.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasonclaninfo
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// clanId:
//     Clan ID. To get a clan ID, use the Clans method. Min value is 1.
// seasonId:
//     Season ID. To get a season ID, use the Seasons method.
// vehicleLevel:
//     List of vehicle Tiers. Maximum limit: 100. Valid values:
//
//     "6" - Vehicles of Tier 6
//     "8" - Vehicles of Tier 8
//     "10" - Vehicles of Tier 10
func (service *WotService) GlobalmapSeasonclaninfo(ctx context.Context, realm Realm, clanId int, seasonId string, vehicleLevel []string, options *wot.GlobalmapSeasonclaninfoOptions) (*wot.GlobalmapSeasonclaninfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"clan_id":       strconv.Itoa(clanId),
		"season_id":     seasonId,
		"vehicle_level": strings.Join(vehicleLevel, ","),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wot.GlobalmapSeasonclaninfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/seasonclaninfo/", reqParam, &data, &metaData)
	return data, metaData, err
}
