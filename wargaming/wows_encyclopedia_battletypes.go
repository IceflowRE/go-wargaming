// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wows"
	"strings"
)

// EncyclopediaBattletypes returns information about battle types.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/battletypes
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaBattletypes(ctx context.Context, realm Realm, options *wows.EncyclopediaBattletypesOptions) (*wows.EncyclopediaBattletypes, error) {
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
	}

	var data *wows.EncyclopediaBattletypes
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/battletypes/", reqParam, &data, nil)
	return data, err
}
