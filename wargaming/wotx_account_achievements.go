// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotx"
	"strings"
)

// AccountAchievements returns players' achievement details.
// Achievement properties define the achievements field values
//
// 1-4 for Mastery Badges and Stage Achievements (type: "class");
// maximum value of Achievement series (type: "series");
// number of achievements earned from sections: Battle Hero, Epic Achievements, Group Achievements, Special Achievements, etc. (type: "repeatable, single, custom").
//
// https://developers.wargaming.net/reference/all/wotx/account/achievements
//
// realm:
//     Valid realms: RealmWgcb
// accountId:
//     Player account ID. Maximum limit: 100.
func (service *WotxService) AccountAchievements(ctx context.Context, realm Realm, accountId []int, options *wotx.AccountAchievementsOptions) (*wotx.AccountAchievements, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
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

	var data *wotx.AccountAchievements
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotx, realm, "/account/achievements/", reqParam, &data, &metaData)
	return data, metaData, err
}
