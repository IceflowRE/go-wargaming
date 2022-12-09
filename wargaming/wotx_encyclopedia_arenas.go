// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wotx"
	"strings"
)

// EncyclopediaArenas returns information about maps.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/arenas
//
// realm:
//     Valid realms: RealmWgcb
func (service *WotxService) EncyclopediaArenas(ctx context.Context, realm Realm, options *wotx.EncyclopediaArenasOptions) (*wotx.EncyclopediaArenas, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
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

	var data *wotx.EncyclopediaArenas
	err := service.client.getRequest(ctx, sectionWotx, realm, "/encyclopedia/arenas/", reqParam, &data, nil)
	return data, err
}
