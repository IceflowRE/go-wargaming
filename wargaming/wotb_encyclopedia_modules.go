package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotb"
	"strings"
)

// EncyclopediaModules returns list of available modules, such as guns, engines, etc.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/modules
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WotbService) EncyclopediaModules(ctx context.Context, realm Realm, options *wotb.EncyclopediaModulesOptions) (*wotb.EncyclopediaModules, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Type_ != nil {
			reqParam["type"] = *options.Type_
		}
		if options.Nation != nil {
			reqParam["nation"] = *options.Nation
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

	var data *wotb.EncyclopediaModules
	err := service.client.getRequest(ctx, sectionWotb, realm, "/encyclopedia/modules/", reqParam, &data)
	return data, err
}
