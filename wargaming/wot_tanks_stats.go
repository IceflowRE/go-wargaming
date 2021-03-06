package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strconv"
	"strings"
)

// TanksStats returns overall statistics, Tank Company statistics, and clan statistics per each vehicle for each user.
//
// https://developers.wargaming.net/reference/all/wot/tanks/stats
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accountId:
//     Player account ID
func (service *WotService) TanksStats(ctx context.Context, realm Realm, accountId int, options *wot.TanksStatsOptions) (*wot.TanksStats, error) {
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
		if options.Extra != nil {
			reqParam["extra"] = strings.Join(options.Extra, ",")
		}
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
	}

	var data *wot.TanksStats
	err := service.client.getRequest(ctx, sectionWot, realm, "/tanks/stats/", reqParam, &data)
	return data, err
}
