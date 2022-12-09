// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wows"
	"strings"
)

// EncyclopediaAccountlevels returns information about Service Record levels.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/accountlevels
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaAccountlevels(ctx context.Context, realm Realm, options *wows.EncyclopediaAccountlevelsOptions) (*wows.EncyclopediaAccountlevels, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wows.EncyclopediaAccountlevels
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/accountlevels/", reqParam, &data, nil)
	return data, err
}
