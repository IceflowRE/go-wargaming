// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strings"
)

// ClansSeason returns information about Clan Battles season.
//
// https://developers.wargaming.net/reference/all/wows/clans/season
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) ClansSeason(ctx context.Context, realm Realm, options *wows.ClansSeasonOptions) (*wows.ClansSeason, *GenericMeta, error) {
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
	}

	var data *wows.ClansSeason
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWows, realm, "/clans/season/", reqParam, &data, &metaData)
	return data, metaData, err
}
