// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// EncyclopediaTankinfo returns vehicle details from Tankopedia.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/tankinfo
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// tankId:
//     Vehicle ID. Maximum limit: 1000.
func (service *WotService) EncyclopediaTankinfo(ctx context.Context, realm Realm, tankId []int, options *wot.EncyclopediaTankinfoOptions) (*wot.EncyclopediaTankinfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tank_id": internal.SliceIntToString(tankId, ","),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wot.EncyclopediaTankinfo
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/tankinfo/", reqParam, &data, nil)
	return data, err
}
