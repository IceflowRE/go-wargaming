package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotx"
	"strconv"
	"strings"
)

// TanksAchievements Method returns players' achievement details.
// Achievement properties define the achievements field values
//
// 1-4 for Mastery Badges and Stage Achievements (type: "class");
// maximum value of Achievement series (type: "series");
// number of achievements earned from sections: Battle Hero, Epic Achievements, Group Achievements, Special Achievements, etc. (type: "repeatable, single, custom").
//
// https://developers.wargaming.net/reference/all/wotx/tanks/achievements
//
// realm:
//     Valid realms: RealmWgcb
// accountId:
//     Player account ID. Min value is 0.
func (service *wotxService) TanksAchievements(ctx context.Context, realm Realm, accountId int, options *wotx.TanksAchievementsOptions) (*wotx.TanksAchievements, error) {
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

	var data *wotx.TanksAchievements
	err := service.client.getRequest(ctx, sectionWotx, realm, "/tanks/achievements/", reqParam, &data)
	return data, err
}
