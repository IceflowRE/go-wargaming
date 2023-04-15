// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strconv"
	"strings"
)

// TanksStats returns general statistics for each vehicle of each player.
//
// https://developers.wargaming.net/reference/all/wotb/tanks/stats
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Player account ID
func (service *WotbService) TanksStats(ctx context.Context, realm Realm, accountId int, options *wotb.TanksStatsOptions) (*wotb.TanksStats, *GenericMeta, error) {
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
		if options.InGarage != nil {
			reqParam["in_garage"] = *options.InGarage
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.TankId != nil {
			reqParam["tank_id"] = internal.SliceIntToString(options.TankId, ",")
		}
	}

	var data *wotb.TanksStats
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tanks/stats/", reqParam, &data, &metaData)
	return data, metaData, err
}
