// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// EncyclopediaVehicleprofiles returns vehicle configuration characteristics.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/vehicleprofiles
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// tankId:
//     Vehicle ID
func (service *WotService) EncyclopediaVehicleprofiles(ctx context.Context, realm Realm, tankId int, options *wot.EncyclopediaVehicleprofilesOptions) (*wot.EncyclopediaVehicleprofiles, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"tank_id": strconv.Itoa(tankId),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.OrderBy != nil {
			reqParam["order_by"] = *options.OrderBy
		}
	}

	var data *wot.EncyclopediaVehicleprofiles
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/vehicleprofiles/", reqParam, &data, &metaData)
	return data, metaData, err
}
