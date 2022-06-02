package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wotx"
	"strings"
)

// EncyclopediaArenas Method returns information about maps.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/arenas
//
// realm:
//     Valid realms: RealmWgcb
func (service *wotxService) EncyclopediaArenas(ctx context.Context, realm Realm, options *wotx.EncyclopediaArenasOptions) (*wotx.EncyclopediaArenas, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
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

	var data *wotx.EncyclopediaArenas
	err := service.client.getRequest(ctx, sectionWotx, realm, "/encyclopedia/arenas/", reqParam, &data)
	return data, err
}
