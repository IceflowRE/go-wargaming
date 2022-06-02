package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/wargaming/wotb"
	"strconv"
	"strings"
)

// TanksStats Method returns general statistics for each vehicle of each player.
//
// https://developers.wargaming.net/reference/all/wotb/tanks/stats
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accountId:
//     Player account ID
func (service *wotbService) TanksStats(ctx context.Context, realm Realm, accountId int, options *wotb.TanksStatsOptions) (*wotb.TanksStats, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
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

	var data *wotb.TanksStats
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tanks/stats/", reqParam, &data)
	return data, err
}
