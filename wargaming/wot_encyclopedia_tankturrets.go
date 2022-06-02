package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/wargaming/wot"
	"strings"
)

// EncyclopediaTankturrets Method returns list of turrets.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/tankturrets
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wotService) EncyclopediaTankturrets(ctx context.Context, realm Realm, options *wot.EncyclopediaTankturretsOptions) (*wot.EncyclopediaTankturrets, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
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

	var data *wot.EncyclopediaTankturrets
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/tankturrets/", reqParam, &data)
	return data, err
}
