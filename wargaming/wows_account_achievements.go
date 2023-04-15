// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strings"
)

// AccountAchievements returns information about players' achievements. Accounts with hidden game profiles are excluded from response. Hidden profiles are listed in the field meta.hidden.
//
// https://developers.wargaming.net/reference/all/wows/account/achievements
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Player account ID. Maximum limit: 100. Min value is 1.
func (service *WowsService) AccountAchievements(ctx context.Context, realm Realm, accountId []int, options *wows.AccountAchievementsOptions) (*wows.AccountAchievements, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"account_id": internal.SliceIntToString(accountId, ","),
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
	}

	var data *wows.AccountAchievements
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWows, realm, "/account/achievements/", reqParam, &data, &metaData)
	return data, metaData, err
}
