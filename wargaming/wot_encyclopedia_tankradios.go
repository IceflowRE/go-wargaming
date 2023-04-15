// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// EncyclopediaTankradios returns list of radios.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/tankradios
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) EncyclopediaTankradios(ctx context.Context, realm Realm, options *wot.EncyclopediaTankradiosOptions) (*wot.EncyclopediaTankradios, error) {
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
			reqParam["nation"] = strings.Join(options.Nation, ",")
		}
	}

	var data *wot.EncyclopediaTankradios
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/tankradios/", reqParam, &data, nil)
	return data, err
}
