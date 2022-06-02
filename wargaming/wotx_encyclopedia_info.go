package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wotx"
	"strings"
)

// EncyclopediaInfo Method returns information about Tankopedia.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/info
//
// realm:
//     Valid realms: RealmWgcb
func (service *wotxService) EncyclopediaInfo(ctx context.Context, realm Realm, options *wotx.EncyclopediaInfoOptions) (*wotx.EncyclopediaInfo, error) {
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

	var data *wotx.EncyclopediaInfo
	err := service.client.getRequest(ctx, sectionWotx, realm, "/encyclopedia/info/", reqParam, &data)
	return data, err
}
