package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotb"
	"strings"
)

// EncyclopediaCrewskills Method returns information about crew skills.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/crewskills
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wotbService) EncyclopediaCrewskills(ctx context.Context, realm Realm, options *wotb.EncyclopediaCrewskillsOptions) (*wotb.EncyclopediaCrewskills, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.VehicleType != nil {
			reqParam["vehicle_type"] = strings.Join(options.VehicleType, ",")
		}
		if options.SkillId != nil {
			reqParam["skill_id"] = strings.Join(options.SkillId, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wotb.EncyclopediaCrewskills
	err := service.client.getRequest(ctx, sectionWotb, realm, "/encyclopedia/crewskills/", reqParam, &data)
	return data, err
}
