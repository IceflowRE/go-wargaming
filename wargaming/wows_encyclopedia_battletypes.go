package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strings"
)

// EncyclopediaBattletypes Method returns information about battle types.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/battletypes
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WowsService) EncyclopediaBattletypes(ctx context.Context, realm Realm, options *wows.EncyclopediaBattletypesOptions) (*wows.EncyclopediaBattletypes, error) {
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

	var data *wows.EncyclopediaBattletypes
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/battletypes/", reqParam, &data)
	return data, err
}
