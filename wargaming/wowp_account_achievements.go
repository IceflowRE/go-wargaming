// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
	"strings"
)

// AccountAchievements returns players' achievement details.
//
// https://developers.wargaming.net/reference/all/wowp/account/achievements
//
// realm:
//     Valid realms: RealmEu, RealmNa
// accountId:
//     Player account ID. Maximum limit: 100.
func (service *WowpService) AccountAchievements(ctx context.Context, realm Realm, accountId []int, options *wowp.AccountAchievementsOptions) (*wowp.AccountAchievements, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
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

	var data *wowp.AccountAchievements
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/account/achievements/", reqParam, &data, &metaData)
	return data, metaData, err
}
