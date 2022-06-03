package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strings"
)

// EncyclopediaCrewskills Method returns information about Commanders' skills.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/crewskills
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wowsService) EncyclopediaCrewskills(ctx context.Context, realm Realm, options *wows.EncyclopediaCrewskillsOptions) (*wows.EncyclopediaCrewskills, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.SkillId != nil {
			reqParam["skill_id"] = internal.SliceIntToString(options.SkillId, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wows.EncyclopediaCrewskills
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/crewskills/", reqParam, &data)
	return data, err
}
