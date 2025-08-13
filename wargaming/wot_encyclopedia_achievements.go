// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// EncyclopediaAchievements returns information about achievements.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/achievements
//
// realm:
//
//	Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) EncyclopediaAchievements(ctx context.Context, realm Realm, options *wot.EncyclopediaAchievementsOptions) (*wot.EncyclopediaAchievements, *GenericMeta, error) {
	if !containsRealm([]Realm{RealmAsia, RealmEu, RealmNa}, realm) {
		return nil, nil, InvalidRealm{realm}
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
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/achievements/", reqParam, &data, &metaData)
	return data, metaData, err
}
