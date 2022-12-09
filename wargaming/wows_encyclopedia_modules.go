// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wows"
	"strconv"
	"strings"
)

// EncyclopediaModules returns list of available modules that can be mounted on a ship (hull, engines, etc.).
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/modules
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaModules(ctx context.Context, realm Realm, options *wows.EncyclopediaModulesOptions) (*wows.EncyclopediaModules, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
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
		if options.ModuleId != nil {
			reqParam["module_id"] = internal.SliceIntToString(options.ModuleId, ",")
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.Type != nil {
			reqParam["type"] = *options.Type
		}
	}

	var data *wows.EncyclopediaModules
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/modules/", reqParam, &data, nil)
	return data, err
}
