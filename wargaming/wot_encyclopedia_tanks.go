// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// EncyclopediaTanks returns list of all vehicles from Tankopedia.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/tanks
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) EncyclopediaTanks(ctx context.Context, realm Realm, options *wot.EncyclopediaTanksOptions) (*wot.EncyclopediaTanks, error) {
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

	var data *wot.EncyclopediaTanks
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/tanks/", reqParam, &data, nil)
	return data, err
}
