// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wot"
	"strings"
)

// EncyclopediaArenas returns information about maps.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/arenas
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) EncyclopediaArenas(ctx context.Context, realm Realm, options *wot.EncyclopediaArenasOptions) (*wot.EncyclopediaArenas, error) {
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

	var data *wot.EncyclopediaArenas
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/arenas/", reqParam, &data, nil)
	return data, err
}
