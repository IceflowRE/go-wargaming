// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotx"
	"strconv"
	"strings"
)

// EncyclopediaVehicles returns list of available vehicles.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/vehicles
//
// realm:
//     Valid realms: RealmWgcb
func (service *WotxService) EncyclopediaVehicles(ctx context.Context, realm Realm, options *wotx.EncyclopediaVehiclesOptions) (*wotx.EncyclopediaVehicles, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
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
	}

	var data *wotx.EncyclopediaVehicles
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotx, realm, "/encyclopedia/vehicles/", reqParam, &data, &metaData)
	return data, metaData, err
}
