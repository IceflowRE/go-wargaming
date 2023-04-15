// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotx"
	"strconv"
	"strings"
)

// EncyclopediaModules returns list of available modules that can be installed on vehicles, such as engines, turrets, etc. At least one input filter parameter (module ID, type) is required to be indicated.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/modules
//
// realm:
//     Valid realms: RealmWgcb
func (service *WotxService) EncyclopediaModules(ctx context.Context, realm Realm, options *wotx.EncyclopediaModulesOptions) (*wotx.EncyclopediaModules, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.Extra != nil {
			reqParam["extra"] = strings.Join(options.Extra, ",")
		}
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
		if options.Nation != nil {
			reqParam["nation"] = strings.Join(options.Nation, ",")
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.Type != nil {
			reqParam["type"] = strings.Join(options.Type, ",")
		}
	}

	var data *wotx.EncyclopediaModules
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotx, realm, "/encyclopedia/modules/", reqParam, &data, &metaData)
	return data, metaData, err
}
