// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wot"
	"strings"
)

// EncyclopediaInfo returns information about Tankopedia.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) EncyclopediaInfo(ctx context.Context, realm Realm, options *wot.EncyclopediaInfoOptions) (*wot.EncyclopediaInfo, error) {
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

	var data *wot.EncyclopediaInfo
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/info/", reqParam, &data, nil)
	return data, err
}
