package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strings"
)

// EncyclopediaAchievements Method returns information about achievements.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/achievements
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WowsService) EncyclopediaAchievements(ctx context.Context, realm Realm, options *wows.EncyclopediaAchievementsOptions) (*wows.EncyclopediaAchievements, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
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
	}

	var data *wows.EncyclopediaAchievements
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/achievements/", reqParam, &data)
	return data, err
}
