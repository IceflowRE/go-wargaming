// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strings"
)

// EncyclopediaCollectioncards returns information about items that are included in the collection.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/collectioncards
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaCollectioncards(ctx context.Context, realm Realm, options *wows.EncyclopediaCollectioncardsOptions) (*wows.EncyclopediaCollectioncards, *GenericMeta, error) {
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

	var data *wows.EncyclopediaCollectioncards
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/collectioncards/", reqParam, &data, &metaData)
	return data, metaData, err
}
