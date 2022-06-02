package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wows"
	"strings"
)

// ClansSeason Method returns information about Clan Battles season.
//
// https://developers.wargaming.net/reference/all/wows/clans/season
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wowsService) ClansSeason(ctx context.Context, realm Realm, options *wows.ClansSeasonOptions) (*wows.ClansSeason, error) {
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
	}

	var data *wows.ClansSeason
	err := service.client.getRequest(ctx, sectionWows, realm, "/clans/season/", reqParam, &data)
	return data, err
}
