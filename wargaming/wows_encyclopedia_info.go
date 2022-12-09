// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wows"
	"strings"
)

// EncyclopediaInfo returns information about encyclopedia.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaInfo(ctx context.Context, realm Realm, options *wows.EncyclopediaInfoOptions) (*wows.EncyclopediaInfo, error) {
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

	var data *wows.EncyclopediaInfo
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/info/", reqParam, &data, nil)
	return data, err
}
