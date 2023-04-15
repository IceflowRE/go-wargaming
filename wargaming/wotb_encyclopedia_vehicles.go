// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strings"
)

// EncyclopediaVehicles returns list of available vehicles.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/vehicles
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotbService) EncyclopediaVehicles(ctx context.Context, realm Realm, options *wotb.EncyclopediaVehiclesOptions) (*wotb.EncyclopediaVehicles, *GenericMeta, error) {
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
		if options.Nation != nil {
			reqParam["nation"] = strings.Join(options.Nation, ",")
		}
		if options.TankId != nil {
			reqParam["tank_id"] = internal.SliceIntToString(options.TankId, ",")
		}
	}

	var data *wotb.EncyclopediaVehicles
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotb, realm, "/encyclopedia/vehicles/", reqParam, &data, &metaData)
	return data, metaData, err
}
