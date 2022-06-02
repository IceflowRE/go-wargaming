package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/wargaming/wotb"
	"strconv"
	"strings"
)

// AccountTankstats Method returns players' statistics on the vehicle.
//
// https://developers.wargaming.net/reference/all/wotb/account/tankstats
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accountId:
//     Player account ID. Maximum limit: 100.
// tankId:
//     Player's vehicle ID
func (service *wotbService) AccountTankstats(ctx context.Context, realm Realm, accountId []int, tankId int, options *wotb.AccountTankstatsOptions) (*wotb.AccountTankstats, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": internal.SliceIntToString(accountId, ","),
		"tank_id":    strconv.Itoa(tankId),
	}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wotb.AccountTankstats
	err := service.client.getRequest(ctx, sectionWotb, realm, "/account/tankstats/", reqParam, &data)
	return data, err
}