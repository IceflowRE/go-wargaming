// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strings"
)

// EncyclopediaAccountlevels returns information about Service Record levels.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/accountlevels
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaAccountlevels(ctx context.Context, realm Realm, options *wows.EncyclopediaAccountlevelsOptions) (*wows.EncyclopediaAccountlevels, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wows.EncyclopediaAccountlevels
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/accountlevels/", reqParam, &data, &metaData)
	return data, metaData, err
}
