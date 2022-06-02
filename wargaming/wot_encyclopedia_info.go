package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wot"
	"strings"
)

// EncyclopediaInfo Method returns information about Tankopedia.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wotService) EncyclopediaInfo(ctx context.Context, realm Realm, options *wot.EncyclopediaInfoOptions) (*wot.EncyclopediaInfo, error) {
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

	var data *wot.EncyclopediaInfo
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/info/", reqParam, &data)
	return data, err
}
