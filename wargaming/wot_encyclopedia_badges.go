// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// EncyclopediaBadges returns list of available badges a player can gain in Ranked Battles.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/badges
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) EncyclopediaBadges(ctx context.Context, realm Realm, options *wot.EncyclopediaBadgesOptions) (*wot.EncyclopediaBadges, *GenericMeta, error) {
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

	var data *wot.EncyclopediaBadges
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/badges/", reqParam, &data, &metaData)
	return data, metaData, err
}
