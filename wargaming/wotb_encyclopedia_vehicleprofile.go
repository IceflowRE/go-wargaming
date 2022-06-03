package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotb"
	"strconv"
	"strings"
)

// EncyclopediaVehicleprofile returns vehicle configuration characteristics based on the specified module IDs.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/vehicleprofile
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// tankId:
//     Vehicle ID
func (service *WotbService) EncyclopediaVehicleprofile(ctx context.Context, realm Realm, tankId int, options *wotb.EncyclopediaVehicleprofileOptions) (*wotb.EncyclopediaVehicleprofile, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
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

	var data *wotb.EncyclopediaVehicleprofile
	err := service.client.getRequest(ctx, sectionWotb, realm, "/encyclopedia/vehicleprofile/", reqParam, &data)
	return data, err
}
