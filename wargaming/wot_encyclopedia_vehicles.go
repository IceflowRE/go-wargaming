// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wot"
	"strconv"
	"strings"
)

// EncyclopediaVehicles returns list of available vehicles.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/vehicles
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) EncyclopediaVehicles(ctx context.Context, realm Realm, options *wot.EncyclopediaVehiclesOptions) (*wot.EncyclopediaVehicles, error) {
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
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Nation != nil {
			reqParam["nation"] = strings.Join(options.Nation, ",")
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.TankId != nil {
			reqParam["tank_id"] = internal.SliceIntToString(options.TankId, ",")
		}
		if options.Tier != nil {
			reqParam["tier"] = internal.SliceIntToString(options.Tier, ",")
		}
		if options.Type != nil {
			reqParam["type"] = strings.Join(options.Type, ",")
		}
	}

	var data *wot.EncyclopediaVehicles
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/vehicles/", reqParam, &data, nil)
	return data, err
}
