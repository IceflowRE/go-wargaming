package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strings"
)

// EncyclopediaTanks Method returns list of all vehicles from Tankopedia.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/tanks
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wotService) EncyclopediaTanks(ctx context.Context, realm Realm, options *wot.EncyclopediaTanksOptions) (*wot.EncyclopediaTanks, error) {
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

	var data *wot.EncyclopediaTanks
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/tanks/", reqParam, &data)
	return data, err
}
