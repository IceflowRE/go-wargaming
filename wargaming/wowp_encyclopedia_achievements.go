// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
	"strings"
)

// EncyclopediaAchievements returns dictionary of achievements from Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/achievements
//
// realm:
//     Valid realms: RealmEu, RealmNa
func (service *WowpService) EncyclopediaAchievements(ctx context.Context, realm Realm, options *wowp.EncyclopediaAchievementsOptions) (*wowp.EncyclopediaAchievements, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, nil, err
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
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/encyclopedia/achievements/", reqParam, &data, &metaData)
	return data, metaData, err
}
