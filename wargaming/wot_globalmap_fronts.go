// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapFronts returns information about the Global Map Fronts.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/fronts
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) GlobalmapFronts(ctx context.Context, realm Realm, options *wot.GlobalmapFrontsOptions) ([]*wot.GlobalmapFronts, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.FrontId != nil {
			reqParam["front_id"] = strings.Join(options.FrontId, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
	}

	var data []*wot.GlobalmapFronts
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/fronts/", reqParam, &data, &metaData)
	return data, metaData, err
}
