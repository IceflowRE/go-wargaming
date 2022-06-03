package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strings"
)

// EncyclopediaInfo Method returns information about encyclopedia.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WowsService) EncyclopediaInfo(ctx context.Context, realm Realm, options *wows.EncyclopediaInfoOptions) (*wows.EncyclopediaInfo, error) {
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

	var data *wows.EncyclopediaInfo
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/info/", reqParam, &data)
	return data, err
}
