// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strings"
)

// EncyclopediaCollections returns information about collections.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/collections
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaCollections(ctx context.Context, realm Realm, options *wows.EncyclopediaCollectionsOptions) (*wows.EncyclopediaCollections, *GenericMeta, error) {
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

	var data *wows.EncyclopediaCollections
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/collections/", reqParam, &data, &metaData)
	return data, metaData, err
}
