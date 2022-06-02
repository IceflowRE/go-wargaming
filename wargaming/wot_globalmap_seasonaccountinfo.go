package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapSeasonaccountinfo Method returns player's statistics for a specific season.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasonaccountinfo
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accountId:
//     Account ID. Min value is 1.
// seasonId:
//     Season ID. To get a season ID, use the Seasons method.
// vehicleLevel:
//     List of vehicle Tiers. Maximum limit: 100. Valid values:
//
//     "6" - Vehicles of Tier 6
//     "8" - Vehicles of Tier 8
//     "10" - Vehicles of Tier 10
func (service *wotService) GlobalmapSeasonaccountinfo(ctx context.Context, realm Realm, accountId int, seasonId string, vehicleLevel []string, options *wot.GlobalmapSeasonaccountinfoOptions) (*wot.GlobalmapSeasonaccountinfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id":    strconv.Itoa(accountId),
		"season_id":     seasonId,
		"vehicle_level": strings.Join(vehicleLevel, ","),
	}
	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wot.GlobalmapSeasonaccountinfo
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/seasonaccountinfo/", reqParam, &data)
	return data, err
}
