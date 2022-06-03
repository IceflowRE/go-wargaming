package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotx"
	"strings"
)

// EncyclopediaAchievements Method returns list of awards, medals, ribbons.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/achievements
//
// realm:
//     Valid realms: RealmWgcb
func (service *WotxService) EncyclopediaAchievements(ctx context.Context, realm Realm, options *wotx.EncyclopediaAchievementsOptions) (*wotx.EncyclopediaAchievements, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Category != nil {
			reqParam["category"] = strings.Join(options.Category, ",")
		}
	}

	var data *wotx.EncyclopediaAchievements
	err := service.client.getRequest(ctx, sectionWotx, realm, "/encyclopedia/achievements/", reqParam, &data)
	return data, err
}
