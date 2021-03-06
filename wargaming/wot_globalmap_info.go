package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strings"
)

// GlobalmapInfo returns general information about the Global Map.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WotService) GlobalmapInfo(ctx context.Context, realm Realm, options *wot.GlobalmapInfoOptions) (*wot.GlobalmapInfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wot.GlobalmapInfo
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/info/", reqParam, &data)
	return data, err
}
