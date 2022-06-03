package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotx"
	"strconv"
	"strings"
)

// EncyclopediaVehicleprofile Method returns vehicle configuration characteristics based on the specified module IDs.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/vehicleprofile
//
// realm:
//     Valid realms: RealmWgcb
// tankId:
//     Vehicle ID
func (service *WotxService) EncyclopediaVehicleprofile(ctx context.Context, realm Realm, tankId int, options *wotx.EncyclopediaVehicleprofileOptions) (*wotx.EncyclopediaVehicleprofile, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tank_id": strconv.Itoa(tankId),
	}
	if options != nil {
		if options.TurretId != nil {
			reqParam["turret_id"] = strconv.Itoa(*options.TurretId)
		}
		if options.SuspensionId != nil {
			reqParam["suspension_id"] = strconv.Itoa(*options.SuspensionId)
		}
		if options.RadioId != nil {
			reqParam["radio_id"] = strconv.Itoa(*options.RadioId)
		}
		if options.ProfileId != nil {
			reqParam["profile_id"] = *options.ProfileId
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.GunId != nil {
			reqParam["gun_id"] = strconv.Itoa(*options.GunId)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.EngineId != nil {
			reqParam["engine_id"] = strconv.Itoa(*options.EngineId)
		}
	}

	var data *wotx.EncyclopediaVehicleprofile
	err := service.client.getRequest(ctx, sectionWotx, realm, "/encyclopedia/vehicleprofile/", reqParam, &data)
	return data, err
}
