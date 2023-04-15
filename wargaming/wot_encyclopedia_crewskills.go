// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// EncyclopediaCrewskills returns full description of all crew skills.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/crewskills
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) EncyclopediaCrewskills(ctx context.Context, realm Realm, options *wot.EncyclopediaCrewskillsOptions) (*wot.EncyclopediaCrewskills, *GenericMeta, error) {
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
			reqParam["role"] = *options.Role
		}
		if options.Skill != nil {
			reqParam["skill"] = strings.Join(options.Skill, ",")
		}
	}

	var data *wot.EncyclopediaCrewskills
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/crewskills/", reqParam, &data, &metaData)
	return data, metaData, err
}
