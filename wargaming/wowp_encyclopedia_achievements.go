// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wowp"
	"strings"
)

// EncyclopediaAchievements returns dictionary of achievements from Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/achievements
//
// realm:
//     Valid realms: RealmEu, RealmNa
func (service *WowpService) EncyclopediaAchievements(ctx context.Context, realm Realm, options *wowp.EncyclopediaAchievementsOptions) (*wowp.EncyclopediaAchievements, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
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

	var data *wowp.EncyclopediaAchievements
	err := service.client.getRequest(ctx, sectionWowp, realm, "/encyclopedia/achievements/", reqParam, &data, nil)
	return data, err
}
