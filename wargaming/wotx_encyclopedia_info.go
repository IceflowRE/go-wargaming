// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotx"
	"strings"
)

// EncyclopediaInfo returns information about Tankopedia.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/info
//
// realm:
//     Valid realms: RealmWgcb
func (service *WotxService) EncyclopediaInfo(ctx context.Context, realm Realm, options *wotx.EncyclopediaInfoOptions) (*wotx.EncyclopediaInfo, error) {
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

	var data *wotx.EncyclopediaInfo
	err := service.client.getRequest(ctx, sectionWotx, realm, "/encyclopedia/info/", reqParam, &data, nil)
	return data, err
}
