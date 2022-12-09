// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wows"
	"strings"
)

// EncyclopediaCollections returns information about collections.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/collections
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaCollections(ctx context.Context, realm Realm, options *wows.EncyclopediaCollectionsOptions) (*wows.EncyclopediaCollections, error) {
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

	var data *wows.EncyclopediaCollections
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/collections/", reqParam, &data, nil)
	return data, err
}
