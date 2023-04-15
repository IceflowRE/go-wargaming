// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strings"
)

// EncyclopediaCrews returns information about Commanders.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/crews
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaCrews(ctx context.Context, realm Realm, options *wows.EncyclopediaCrewsOptions) (*wows.EncyclopediaCrews, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.CommanderId != nil {
			reqParam["commander_id"] = internal.SliceIntToString(options.CommanderId, ",")
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wows.EncyclopediaCrews
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/crews/", reqParam, &data, &metaData)
	return data, metaData, err
}
