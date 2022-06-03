package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotb"
	"strings"
)

// EncyclopediaInfo Method returns information about Tankopedia.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WotbService) EncyclopediaInfo(ctx context.Context, realm Realm, options *wotb.EncyclopediaInfoOptions) (*wotb.EncyclopediaInfo, error) {
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

	var data *wotb.EncyclopediaInfo
	err := service.client.getRequest(ctx, sectionWotb, realm, "/encyclopedia/info/", reqParam, &data)
	return data, err
}
