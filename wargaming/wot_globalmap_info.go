// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// GlobalmapInfo returns general information about the Global Map.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) GlobalmapInfo(ctx context.Context, realm Realm, options *wot.GlobalmapInfoOptions) (*wot.GlobalmapInfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wot.GlobalmapInfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/info/", reqParam, &data, &metaData)
	return data, metaData, err
}
