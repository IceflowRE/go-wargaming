package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strings"
)

// EncyclopediaTankinfo Method returns vehicle details from Tankopedia.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/tankinfo
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// tankId:
//     Vehicle ID. Maximum limit: 1000.
func (service *wotService) EncyclopediaTankinfo(ctx context.Context, realm Realm, tankId []int, options *wot.EncyclopediaTankinfoOptions) (*wot.EncyclopediaTankinfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tank_id": internal.SliceIntToString(tankId, ","),
	}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wot.EncyclopediaTankinfo
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/tankinfo/", reqParam, &data)
	return data, err
}
