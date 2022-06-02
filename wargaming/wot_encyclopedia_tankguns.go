package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/wargaming/wot"
	"strconv"
	"strings"
)

// EncyclopediaTankguns Method returns a tanks' gun list.
// The logic of this method and some field values may vary according to optional parameters passed.
// Changeable fields:
//
// damage
// piercing_power
// rate
// price_credit
// price_gold
//
// Optional input parameters work as follows:
//
// correct turret_id passed — tank guns are filtered by whether they are placed on the turret and the abovementioned values change according to the turret;
// correct turret_id and module_id passed — the method returns details on each module with the abovementioned values changed according to the turret, or returns null if the module is not compatible with the turret;
// correct tank_id passed — if tank type matches one of AT-SPG, SPG, mediumTank, tank guns are filtered by whether they belong to the tank, the abovementioned values change according to the tank; otherwise, returns an error and requests turret_id. If module_id is also passed, the method returns details on each module with the abovementioned values changed according to the tank, or returns null if the module is not compatible with the tank;
// compatible turret_id and tank_id passed — tank guns are filtered by whether they belong to the tank and are placed on the turret, the abovementioned values changed according to the turret.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/tankguns
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wotService) EncyclopediaTankguns(ctx context.Context, realm Realm, options *wot.EncyclopediaTankgunsOptions) (*wot.EncyclopediaTankguns, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.TurretId != nil {
			reqParam["turret_id"] = strconv.Itoa(*options.TurretId)
		}
		if options.TankId != nil {
			reqParam["tank_id"] = strconv.Itoa(*options.TankId)
		}
		if options.Nation != nil {
			reqParam["nation"] = strings.Join(options.Nation, ",")
		}
		if options.ModuleId != nil {
			reqParam["module_id"] = internal.SliceIntToString(options.ModuleId, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wot.EncyclopediaTankguns
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/tankguns/", reqParam, &data)
	return data, err
}