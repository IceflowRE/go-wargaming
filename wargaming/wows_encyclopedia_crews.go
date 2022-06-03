package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strings"
)

// EncyclopediaCrews returns information about Commanders.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/crews
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WowsService) EncyclopediaCrews(ctx context.Context, realm Realm, options *wows.EncyclopediaCrewsOptions) (*wows.EncyclopediaCrews, error) {
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
		if options.CommanderId != nil {
			reqParam["commander_id"] = internal.SliceIntToString(options.CommanderId, ",")
		}
	}

	var data *wows.EncyclopediaCrews
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/crews/", reqParam, &data)
	return data, err
}
