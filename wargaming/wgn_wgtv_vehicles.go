package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/wargaming/wgn"
)

// WgtvVehicles Method returns list of vehicles sorted by games and covered by videos.
//
// https://developers.wargaming.net/reference/all/wgn/wgtv/vehicles
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wgnService) WgtvVehicles(ctx context.Context, realm Realm, options *wgn.WgtvVehiclesOptions) (*wgn.WgtvVehicles, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.ProjectId != nil {
			reqParam["project_id"] = internal.SliceIntToString(options.ProjectId, ",")
		}
		if options.ProgramId != nil {
			reqParam["program_id"] = internal.SliceIntToString(options.ProgramId, ",")
		}
		if options.CategoryId != nil {
			reqParam["category_id"] = internal.SliceIntToString(options.CategoryId, ",")
		}
	}

	var data *wgn.WgtvVehicles
	err := service.client.getRequest(ctx, sectionWgn, realm, "/wgtv/vehicles/", reqParam, &data)
	return data, err
}
