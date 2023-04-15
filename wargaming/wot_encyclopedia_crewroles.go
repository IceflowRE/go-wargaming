// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// EncyclopediaCrewroles returns full description of all crew qualifications.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/crewroles
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) EncyclopediaCrewroles(ctx context.Context, realm Realm, options *wot.EncyclopediaCrewrolesOptions) (*wot.EncyclopediaCrewroles, *GenericMeta, error) {
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
		if options.Role != nil {
			reqParam["role"] = strings.Join(options.Role, ",")
		}
	}

	var data *wot.EncyclopediaCrewroles
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/crewroles/", reqParam, &data, &metaData)
	return data, metaData, err
}
