package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strings"
)

// AccountAchievements Method returns players' achievement details.
// Achievement properties define the achievements field values:
//
// 1-4 for Mastery Badges and Stage Achievements (type: "class");
// maximum value of Achievement series (type: "series");
// number of achievements earned from sections: Battle Hero, Epic Achievements, Group Achievements, Special Achievements, etc. (type: "repeatable, single, custom").
//
// https://developers.wargaming.net/reference/all/wot/account/achievements
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accountId:
//     Player account ID. Maximum limit: 100.
func (service *WotService) AccountAchievements(ctx context.Context, realm Realm, accountId []int, options *wot.AccountAchievementsOptions) (*wot.AccountAchievements, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": internal.SliceIntToString(accountId, ","),
	}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wot.AccountAchievements
	err := service.client.getRequest(ctx, sectionWot, realm, "/account/achievements/", reqParam, &data)
	return data, err
}
