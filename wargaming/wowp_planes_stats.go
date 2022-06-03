package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wowp"
	"strconv"
	"strings"
)

// PlanesStats Method returns statistics on player's aircraft.
//
// https://developers.wargaming.net/reference/all/wowp/planes/stats
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
// accountId:
//     Player account ID
func (service *WowpService) PlanesStats(ctx context.Context, realm Realm, accountId int, options *wowp.PlanesStatsOptions) (*wowp.PlanesStats, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
	}
	if options != nil {
		if options.PlaneId != nil {
			reqParam["plane_id"] = internal.SliceIntToString(options.PlaneId, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.InGarage != nil {
			reqParam["in_garage"] = *options.InGarage
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
	}

	var data *wowp.PlanesStats
	err := service.client.getRequest(ctx, sectionWowp, realm, "/planes/stats/", reqParam, &data)
	return data, err
}
