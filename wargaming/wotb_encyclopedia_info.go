// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strings"
)

// EncyclopediaInfo returns information about Tankopedia.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotbService) EncyclopediaInfo(ctx context.Context, realm Realm, options *wotb.EncyclopediaInfoOptions) (*wotb.EncyclopediaInfo, error) {
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

	var data *wotb.EncyclopediaInfo
	err := service.client.getRequest(ctx, sectionWotb, realm, "/encyclopedia/info/", reqParam, &data, nil)
	return data, err
}
