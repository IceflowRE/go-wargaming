package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotb"
	"strings"
)

// EncyclopediaProvisions Method returns list of available equipment and consumables.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/provisions
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wotbService) EncyclopediaProvisions(ctx context.Context, realm Realm, options *wotb.EncyclopediaProvisionsOptions) (*wotb.EncyclopediaProvisions, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Type_ != nil {
			reqParam["type"] = *options.Type_
		}
		if options.TankId != nil {
			reqParam["tank_id"] = internal.SliceIntToString(options.TankId, ",")
		}
		if options.ProvisionId != nil {
			reqParam["provision_id"] = internal.SliceIntToString(options.ProvisionId, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wotb.EncyclopediaProvisions
	err := service.client.getRequest(ctx, sectionWotb, realm, "/encyclopedia/provisions/", reqParam, &data)
	return data, err
}
