// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strings"
)

// SeasonsInfo returns information about Ranked Battles seasons.
//
// https://developers.wargaming.net/reference/all/wows/seasons/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) SeasonsInfo(ctx context.Context, realm Realm, options *wows.SeasonsInfoOptions) (*wows.SeasonsInfo, error) {
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
		if options.SeasonId != nil {
			reqParam["season_id"] = internal.SliceIntToString(options.SeasonId, ",")
		}
	}

	var data *wows.SeasonsInfo
	err := service.client.getRequest(ctx, sectionWows, realm, "/seasons/info/", reqParam, &data, nil)
	return data, err
}
