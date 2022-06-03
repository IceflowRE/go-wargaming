package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strings"
)

// RatingsTypes returns dictionary of rating periods and ratings details.
//
// https://developers.wargaming.net/reference/all/wot/ratings/types
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WotService) RatingsTypes(ctx context.Context, realm Realm, options *wot.RatingsTypesOptions) (*wot.RatingsTypes, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.BattleType != nil {
			reqParam["battle_type"] = *options.BattleType
		}
	}

	var data *wot.RatingsTypes
	err := service.client.getRequest(ctx, sectionWot, realm, "/ratings/types/", reqParam, &data)
	return data, err
}
