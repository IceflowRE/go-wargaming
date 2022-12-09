// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wows"
	"strings"
)

// EncyclopediaCrewranks returns information about Commanders' ranks.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/crewranks
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaCrewranks(ctx context.Context, realm Realm, options *wows.EncyclopediaCrewranksOptions) (*wows.EncyclopediaCrewranks, error) {
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
		if options.Nation != nil {
			reqParam["nation"] = *options.Nation
		}
	}

	var data *wows.EncyclopediaCrewranks
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/crewranks/", reqParam, &data, nil)
	return data, err
}
