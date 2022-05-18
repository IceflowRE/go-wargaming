package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgn"
)

// WgnWgtvVehicles Method returns list of vehicles sorted by games and covered by videos.
//
// https://developers.wargaming.net/reference/all/wgn/wgtv/vehicles
//
// category_id:
//     Content category ID. Maximum limit: 100.
// program_id:
//     Program ID. Maximum limit: 100.
// project_id:
//     Game project ID. Maximum limit: 100.
func (client *Client) WgnWgtvVehicles(realm Realm, categoryId []int, programId []int, projectId []int) (*wgn.WgtvVehicles, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"category_id": utils.SliceIntToString(categoryId, ","),
		"program_id": utils.SliceIntToString(programId, ","),
		"project_id": utils.SliceIntToString(projectId, ","),
	}

	var result *wgn.WgtvVehicles
	err := client.doGetDataRequest(realm, "/wgn/wgtv/vehicles/", reqParam, &result)
	return result, err
}
