package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotb"
	"strings"
)

// EncyclopediaVehicleprofiles Method returns vehicle configuration characteristics.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/vehicleprofiles
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// tankId:
//     Vehicle ID. Maximum limit: 25.
func (service *WotbService) EncyclopediaVehicleprofiles(ctx context.Context, realm Realm, tankId []int, options *wotb.EncyclopediaVehicleprofilesOptions) (*wotb.EncyclopediaVehicleprofiles, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tank_id": internal.SliceIntToString(tankId, ","),
	}
	if options != nil {
		if options.OrderBy != nil {
			reqParam["order_by"] = *options.OrderBy
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wotb.EncyclopediaVehicleprofiles
	err := service.client.getRequest(ctx, sectionWotb, realm, "/encyclopedia/vehicleprofiles/", reqParam, &data)
	return data, err
}
