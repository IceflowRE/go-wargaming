package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strings"
)

// EncyclopediaCrewroles Method returns full description of all crew qualifications.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/crewroles
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WotService) EncyclopediaCrewroles(ctx context.Context, realm Realm, options *wot.EncyclopediaCrewrolesOptions) (*wot.EncyclopediaCrewroles, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Role != nil {
			reqParam["role"] = strings.Join(options.Role, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wot.EncyclopediaCrewroles
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/crewroles/", reqParam, &data)
	return data, err
}
