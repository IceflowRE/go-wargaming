package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wowp"
	"strings"
)

// AccountAchievements Method returns players' achievement details.
//
// https://developers.wargaming.net/reference/all/wowp/account/achievements
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
// accountId:
//     Player account ID. Maximum limit: 100.
func (service *wowpService) AccountAchievements(ctx context.Context, realm Realm, accountId []int, options *wowp.AccountAchievementsOptions) (*wowp.AccountAchievements, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
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

	var data *wowp.AccountAchievements
	err := service.client.getRequest(ctx, sectionWowp, realm, "/account/achievements/", reqParam, &data)
	return data, err
}
