// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strings"
)

// EncyclopediaCrewskills returns information about crew skills.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/crewskills
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotbService) EncyclopediaCrewskills(ctx context.Context, realm Realm, options *wotb.EncyclopediaCrewskillsOptions) (*wotb.EncyclopediaCrewskills, error) {
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
			reqParam["skill_id"] = strings.Join(options.SkillId, ",")
		}
		if options.VehicleType != nil {
			reqParam["vehicle_type"] = strings.Join(options.VehicleType, ",")
		}
	}

	var data *wotb.EncyclopediaCrewskills
	err := service.client.getRequest(ctx, sectionWotb, realm, "/encyclopedia/crewskills/", reqParam, &data, nil)
	return data, err
}
