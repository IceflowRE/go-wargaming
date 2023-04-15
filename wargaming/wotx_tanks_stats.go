// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotx"
	"strconv"
	"strings"
)

// TanksStats returns overall statistics, Tank Company statistics, and clan statistics per each vehicle for each user.
//
// https://developers.wargaming.net/reference/all/wotx/tanks/stats
//
// realm:
//     Valid realms: RealmWgcb
// accountId:
//     Player account ID. Min value is 0.
func (service *WotxService) TanksStats(ctx context.Context, realm Realm, accountId int, options *wotx.TanksStatsOptions) (*wotx.TanksStats, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
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

	var data *wotx.TanksStats
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotx, realm, "/tanks/stats/", reqParam, &data, &metaData)
	return data, metaData, err
}
