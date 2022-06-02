package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/wargaming/wotx"
	"strings"
)

// EncyclopediaVehicles Method returns list of available vehicles.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/vehicles
//
// realm:
//     Valid realms: RealmWgcb
func (service *wotxService) EncyclopediaVehicles(ctx context.Context, realm Realm, options *wotx.EncyclopediaVehiclesOptions) (*wotx.EncyclopediaVehicles, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Tier != nil {
			reqParam["tier"] = internal.SliceIntToString(options.Tier, ",")
		}
		if options.TankId != nil {
			reqParam["tank_id"] = internal.SliceIntToString(options.TankId, ",")
		}
		if options.Nation != nil {
			reqParam["nation"] = strings.Join(options.Nation, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wotx.EncyclopediaVehicles
	err := service.client.getRequest(ctx, sectionWotx, realm, "/encyclopedia/vehicles/", reqParam, &data)
	return data, err
}
