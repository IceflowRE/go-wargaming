package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wot"
	"strconv"
	"strings"
)

// EncyclopediaVehicleprofiles Method returns vehicle configuration characteristics.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/vehicleprofiles
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// tankId:
//     Vehicle ID
func (service *wotService) EncyclopediaVehicleprofiles(ctx context.Context, realm Realm, tankId int, options *wot.EncyclopediaVehicleprofilesOptions) (*wot.EncyclopediaVehicleprofiles, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tank_id": strconv.Itoa(tankId),
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

	var data *wot.EncyclopediaVehicleprofiles
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/vehicleprofiles/", reqParam, &data)
	return data, err
}
