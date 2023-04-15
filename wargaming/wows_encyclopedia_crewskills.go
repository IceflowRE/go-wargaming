// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strings"
)

// EncyclopediaCrewskills returns information about Commanders' skills.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/crewskills
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaCrewskills(ctx context.Context, realm Realm, options *wows.EncyclopediaCrewskillsOptions) (*wows.EncyclopediaCrewskills, error) {
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
		if options.SkillId != nil {
			reqParam["skill_id"] = internal.SliceIntToString(options.SkillId, ",")
		}
	}

	var data *wows.EncyclopediaCrewskills
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/crewskills/", reqParam, &data, nil)
	return data, err
}
