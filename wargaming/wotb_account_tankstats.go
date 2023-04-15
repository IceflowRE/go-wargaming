// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strconv"
	"strings"
)

// AccountTankstats returns players' statistics on the vehicle.
//
// https://developers.wargaming.net/reference/all/wotb/account/tankstats
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Player account ID. Maximum limit: 100.
// tankId:
//     Player's vehicle ID
func (service *WotbService) AccountTankstats(ctx context.Context, realm Realm, accountId []int, tankId int, options *wotb.AccountTankstatsOptions) (*wotb.AccountTankstats, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"account_id": internal.SliceIntToString(accountId, ","),
		"tank_id":    strconv.Itoa(tankId),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wotb.AccountTankstats
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotb, realm, "/account/tankstats/", reqParam, &data, &metaData)
	return data, metaData, err
}
