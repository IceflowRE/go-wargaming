package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strings"
)

// EncyclopediaCrewranks Method returns information about Commanders' ranks.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/crewranks
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WowsService) EncyclopediaCrewranks(ctx context.Context, realm Realm, options *wows.EncyclopediaCrewranksOptions) (*wows.EncyclopediaCrewranks, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Nation != nil {
			reqParam["nation"] = *options.Nation
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wows.EncyclopediaCrewranks
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/crewranks/", reqParam, &data)
	return data, err
}
