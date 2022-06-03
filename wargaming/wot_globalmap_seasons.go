package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapSeasons Method returns information about seasons.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasons
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WotService) GlobalmapSeasons(ctx context.Context, realm Realm, options *wot.GlobalmapSeasonsOptions) ([]*wot.GlobalmapSeasons, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Status != nil {
			reqParam["status"] = *options.Status
		}
		if options.SeasonId != nil {
			reqParam["season_id"] = *options.SeasonId
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data []*wot.GlobalmapSeasons
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/seasons/", reqParam, &data)
	return data, err
}
