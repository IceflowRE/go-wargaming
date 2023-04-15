// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapSeasons returns information about seasons.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasons
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) GlobalmapSeasons(ctx context.Context, realm Realm, options *wot.GlobalmapSeasonsOptions) ([]*wot.GlobalmapSeasons, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
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
		if options.SeasonId != nil {
			reqParam["season_id"] = *options.SeasonId
		}
		if options.Status != nil {
			reqParam["status"] = *options.Status
		}
	}

	var data []*wot.GlobalmapSeasons
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/seasons/", reqParam, &data, &metaData)
	return data, metaData, err
}
