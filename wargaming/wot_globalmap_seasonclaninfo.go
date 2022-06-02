package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapSeasonclaninfo Method returns clan's statistics for a specific season.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasonclaninfo
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
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
func (service *wotService) GlobalmapSeasonclaninfo(ctx context.Context, realm Realm, clanId int, seasonId string, vehicleLevel []string, options *wot.GlobalmapSeasonclaninfoOptions) (*wot.GlobalmapSeasonclaninfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
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
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/seasonclaninfo/", reqParam, &data)
	return data, err
}