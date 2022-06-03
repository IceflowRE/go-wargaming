package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strings"
)

// EncyclopediaCrewskills Method returns full description of all crew skills.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/crewskills
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wotService) EncyclopediaCrewskills(ctx context.Context, realm Realm, options *wot.EncyclopediaCrewskillsOptions) (*wot.EncyclopediaCrewskills, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Skill != nil {
			reqParam["skill"] = strings.Join(options.Skill, ",")
		}
		if options.Role != nil {
			reqParam["role"] = *options.Role
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wot.EncyclopediaCrewskills
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/crewskills/", reqParam, &data)
	return data, err
}
