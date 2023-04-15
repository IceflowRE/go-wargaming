// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strconv"
	"strings"
)

// ShipsStats returns general statistics for each ship of a player. Accounts with hidden game profiles are excluded from response. Hidden profiles are listed in the field meta.hidden.
//
// https://developers.wargaming.net/reference/all/wows/ships/stats
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Player account ID
func (service *WowsService) ShipsStats(ctx context.Context, realm Realm, accountId int, options *wows.ShipsStatsOptions) (map[int][]*wows.ShipsStats, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
	}

	if options != nil {
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
		if options.Extra != nil {
			reqParam["extra"] = strings.Join(options.Extra, ",")
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
		if options.ShipId != nil {
			reqParam["ship_id"] = internal.SliceIntToString(options.ShipId, ",")
		}
	}

	var data map[int][]*wows.ShipsStats
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWows, realm, "/ships/stats/", reqParam, &data, &metaData)
	return data, metaData, err
}
