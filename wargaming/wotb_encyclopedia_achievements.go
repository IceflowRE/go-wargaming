// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strings"
)

// EncyclopediaAchievements returns information about achievements.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/achievements
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotbService) EncyclopediaAchievements(ctx context.Context, realm Realm, options *wotb.EncyclopediaAchievementsOptions) (*wotb.EncyclopediaAchievements, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
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

	var data *wotb.EncyclopediaAchievements
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotb, realm, "/encyclopedia/achievements/", reqParam, &data, &metaData)
	return data, metaData, err
}
