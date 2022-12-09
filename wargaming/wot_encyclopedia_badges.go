// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wot"
	"strings"
)

// EncyclopediaBadges returns list of available badges a player can gain in Ranked Battles.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/badges
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) EncyclopediaBadges(ctx context.Context, realm Realm, options *wot.EncyclopediaBadgesOptions) (*wot.EncyclopediaBadges, error) {
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

	var data *wot.EncyclopediaBadges
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/badges/", reqParam, &data, nil)
	return data, err
}
