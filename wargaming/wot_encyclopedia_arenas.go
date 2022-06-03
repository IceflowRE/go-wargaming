package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strings"
)

// EncyclopediaArenas Method returns information about maps.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/arenas
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WotService) EncyclopediaArenas(ctx context.Context, realm Realm, options *wot.EncyclopediaArenasOptions) (*wot.EncyclopediaArenas, error) {
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

	var data *wot.EncyclopediaArenas
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/arenas/", reqParam, &data)
	return data, err
}
