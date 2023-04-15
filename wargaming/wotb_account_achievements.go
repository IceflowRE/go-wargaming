// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strings"
)

// AccountAchievements returns players' achievement details.
//
// https://developers.wargaming.net/reference/all/wotb/account/achievements
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Player account ID. Maximum limit: 100.
func (service *WotbService) AccountAchievements(ctx context.Context, realm Realm, accountId []int, options *wotb.AccountAchievementsOptions) (*wotb.AccountAchievements, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"account_id": internal.SliceIntToString(accountId, ","),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wotb.AccountAchievements
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotb, realm, "/account/achievements/", reqParam, &data, &metaData)
	return data, metaData, err
}
