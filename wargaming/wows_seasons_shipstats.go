// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strconv"
	"strings"
)

// SeasonsShipstats returns players' ships statistics in Ranked Battles seasons. Accounts with hidden game profiles are excluded from response. Hidden profiles are listed in the field meta.hidden.
//
// https://developers.wargaming.net/reference/all/wows/seasons/shipstats
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Player account ID
func (service *WowsService) SeasonsShipstats(ctx context.Context, realm Realm, accountId int, options *wows.SeasonsShipstatsOptions) (*wows.SeasonsShipstats, *GenericMeta, error) {
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
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.SeasonId != nil {
			reqParam["season_id"] = internal.SliceIntToString(options.SeasonId, ",")
		}
		if options.ShipId != nil {
			reqParam["ship_id"] = internal.SliceIntToString(options.ShipId, ",")
		}
	}

	var data *wows.SeasonsShipstats
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWows, realm, "/seasons/shipstats/", reqParam, &data, &metaData)
	return data, metaData, err
}
