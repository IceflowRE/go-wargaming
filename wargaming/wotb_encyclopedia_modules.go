// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strings"
)

// EncyclopediaModules returns list of available modules, such as guns, engines, etc.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/modules
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotbService) EncyclopediaModules(ctx context.Context, realm Realm, options *wotb.EncyclopediaModulesOptions) (*wotb.EncyclopediaModules, error) {
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
		if options.ModuleId != nil {
			reqParam["module_id"] = internal.SliceIntToString(options.ModuleId, ",")
		}
		if options.Nation != nil {
			reqParam["nation"] = *options.Nation
		}
		if options.Type != nil {
			reqParam["type"] = *options.Type
		}
	}

	var data *wotb.EncyclopediaModules
	err := service.client.getRequest(ctx, sectionWotb, realm, "/encyclopedia/modules/", reqParam, &data, nil)
	return data, err
}
