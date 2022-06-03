package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotx"
	"strconv"
	"strings"
)

// TanksStats Method returns overall statistics, Tank Company statistics, and clan statistics per each vehicle for each user.
//
// https://developers.wargaming.net/reference/all/wotx/tanks/stats
//
// realm:
//     Valid realms: RealmWgcb
// accountId:
//     Player account ID. Min value is 0.
func (service *wotxService) TanksStats(ctx context.Context, realm Realm, accountId int, options *wotx.TanksStatsOptions) (*wotx.TanksStats, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
	}
	if options != nil {
		if options.TankId != nil {
			reqParam["tank_id"] = internal.SliceIntToString(options.TankId, ",")
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

	var data *wotx.TanksStats
	err := service.client.getRequest(ctx, sectionWotx, realm, "/tanks/stats/", reqParam, &data)
	return data, err
}
