package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strconv"
	"strings"
)

// EncyclopediaModules Method returns list of available modules that can be mounted on a ship (hull, engines, etc.).
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/modules
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WowsService) EncyclopediaModules(ctx context.Context, realm Realm, options *wows.EncyclopediaModulesOptions) (*wows.EncyclopediaModules, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Type_ != nil {
			reqParam["type"] = *options.Type_
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.ModuleId != nil {
			reqParam["module_id"] = internal.SliceIntToString(options.ModuleId, ",")
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wows.EncyclopediaModules
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/modules/", reqParam, &data)
	return data, err
}
