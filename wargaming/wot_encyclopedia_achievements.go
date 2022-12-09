// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wot"
	"strings"
)

// EncyclopediaAchievements returns information about achievements.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/achievements
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) EncyclopediaAchievements(ctx context.Context, realm Realm, options *wot.EncyclopediaAchievementsOptions) (*wot.EncyclopediaAchievements, error) {
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

	var data *wot.EncyclopediaAchievements
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/achievements/", reqParam, &data, nil)
	return data, err
}
