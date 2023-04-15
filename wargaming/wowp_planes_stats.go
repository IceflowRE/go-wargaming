// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
	"strconv"
	"strings"
)

// PlanesStats returns statistics on player's aircraft.
//
// https://developers.wargaming.net/reference/all/wowp/planes/stats
//
// realm:
//     Valid realms: RealmEu, RealmNa
// accountId:
//     Player account ID
func (service *WowpService) PlanesStats(ctx context.Context, realm Realm, accountId int, options *wowp.PlanesStatsOptions) (*wowp.PlanesStats, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
	}

	if options != nil {
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.InGarage != nil {
			reqParam["in_garage"] = *options.InGarage
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.PlaneId != nil {
			reqParam["plane_id"] = internal.SliceIntToString(options.PlaneId, ",")
		}
	}

	var data *wowp.PlanesStats
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/planes/stats/", reqParam, &data, &metaData)
	return data, metaData, err
}
