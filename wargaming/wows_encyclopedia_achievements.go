// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wows"
	"strings"
)

// EncyclopediaAchievements returns information about achievements.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/achievements
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaAchievements(ctx context.Context, realm Realm, options *wows.EncyclopediaAchievementsOptions) (*wows.EncyclopediaAchievements, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wows.EncyclopediaAchievements
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/achievements/", reqParam, &data, nil)
	return data, err
}
