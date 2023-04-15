// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// EncyclopediaVehicleprofile returns vehicle configuration characteristics based on the specified module IDs.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/vehicleprofile
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// tankId:
//     Vehicle ID
func (service *WotService) EncyclopediaVehicleprofile(ctx context.Context, realm Realm, tankId int, options *wot.EncyclopediaVehicleprofileOptions) (*wot.EncyclopediaVehicleprofile, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"tank_id": strconv.Itoa(tankId),
	}

	if options != nil {
		if options.EngineId != nil {
			reqParam["engine_id"] = strconv.Itoa(*options.EngineId)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.GunId != nil {
			reqParam["gun_id"] = strconv.Itoa(*options.GunId)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.ProfileId != nil {
			reqParam["profile_id"] = *options.ProfileId
		}
		if options.RadioId != nil {
			reqParam["radio_id"] = strconv.Itoa(*options.RadioId)
		}
		if options.SuspensionId != nil {
			reqParam["suspension_id"] = strconv.Itoa(*options.SuspensionId)
		}
		if options.TurretId != nil {
			reqParam["turret_id"] = strconv.Itoa(*options.TurretId)
		}
	}

	var data *wot.EncyclopediaVehicleprofile
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/vehicleprofile/", reqParam, &data, &metaData)
	return data, metaData, err
}
