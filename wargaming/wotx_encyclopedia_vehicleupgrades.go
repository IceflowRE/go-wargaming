// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wotx"
	"strings"
)

// EncyclopediaVehicleupgrades returns list of vehicle equipment and consumables.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/vehicleupgrades
//
// realm:
//     Valid realms: RealmWgcb
// tankId:
//     Vehicle ID. Maximum limit: 100.
func (service *WotxService) EncyclopediaVehicleupgrades(ctx context.Context, realm Realm, tankId []int, options *wotx.EncyclopediaVehicleupgradesOptions) (*wotx.EncyclopediaVehicleupgrades, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tank_id": internal.SliceIntToString(tankId, ","),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wotx.EncyclopediaVehicleupgrades
	err := service.client.getRequest(ctx, sectionWotx, realm, "/encyclopedia/vehicleupgrades/", reqParam, &data, nil)
	return data, err
}
