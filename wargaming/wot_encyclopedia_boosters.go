// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// EncyclopediaBoosters returns information about Personal Reserves.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/boosters
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) EncyclopediaBoosters(ctx context.Context, realm Realm, options *wot.EncyclopediaBoostersOptions) (*wot.EncyclopediaBoosters, *GenericMeta, error) {
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

	var data *wot.EncyclopediaBoosters
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/boosters/", reqParam, &data, &metaData)
	return data, metaData, err
}
