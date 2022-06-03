package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strings"
)

// EncyclopediaBoosters Method returns information about Personal Reserves.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/boosters
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wotService) EncyclopediaBoosters(ctx context.Context, realm Realm, options *wot.EncyclopediaBoostersOptions) (*wot.EncyclopediaBoosters, error) {
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

	var data *wot.EncyclopediaBoosters
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/boosters/", reqParam, &data)
	return data, err
}
