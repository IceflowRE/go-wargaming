// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgn"
)

// WgtvVehicles returns list of vehicles sorted by games and covered by videos.
//
// https://developers.wargaming.net/reference/all/wgn/wgtv/vehicles
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WgnService) WgtvVehicles(ctx context.Context, realm Realm, options *wgn.WgtvVehiclesOptions) (*wgn.WgtvVehicles, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.CategoryId != nil {
			reqParam["category_id"] = internal.SliceIntToString(options.CategoryId, ",")
		}
		if options.ProgramId != nil {
			reqParam["program_id"] = internal.SliceIntToString(options.ProgramId, ",")
		}
		if options.ProjectId != nil {
			reqParam["project_id"] = internal.SliceIntToString(options.ProjectId, ",")
		}
	}

	var data *wgn.WgtvVehicles
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWgn, realm, "/wgtv/vehicles/", reqParam, &data, &metaData)
	return data, metaData, err
}
