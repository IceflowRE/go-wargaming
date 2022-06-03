package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strings"
)

// EncyclopediaCollections Method returns information about collections.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/collections
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wowsService) EncyclopediaCollections(ctx context.Context, realm Realm, options *wows.EncyclopediaCollectionsOptions) (*wows.EncyclopediaCollections, error) {
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

	var data *wows.EncyclopediaCollections
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/collections/", reqParam, &data)
	return data, err
}
