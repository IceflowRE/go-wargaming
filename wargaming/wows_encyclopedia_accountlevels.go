package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strings"
)

// EncyclopediaAccountlevels Method returns information about Service Record levels.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/accountlevels
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wowsService) EncyclopediaAccountlevels(ctx context.Context, realm Realm, options *wows.EncyclopediaAccountlevelsOptions) (*wows.EncyclopediaAccountlevels, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wows.EncyclopediaAccountlevels
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/accountlevels/", reqParam, &data)
	return data, err
}
