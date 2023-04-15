// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strings"
)

// EncyclopediaProvisions returns list of available equipment and consumables.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/provisions
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotbService) EncyclopediaProvisions(ctx context.Context, realm Realm, options *wotb.EncyclopediaProvisionsOptions) (*wotb.EncyclopediaProvisions, *GenericMeta, error) {
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
		if options.ProvisionId != nil {
			reqParam["provision_id"] = internal.SliceIntToString(options.ProvisionId, ",")
		}
		if options.TankId != nil {
			reqParam["tank_id"] = internal.SliceIntToString(options.TankId, ",")
		}
		if options.Type != nil {
			reqParam["type"] = *options.Type
		}
	}

	var data *wotb.EncyclopediaProvisions
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotb, realm, "/encyclopedia/provisions/", reqParam, &data, &metaData)
	return data, metaData, err
}
