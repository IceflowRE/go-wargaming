// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// RatingsTypes returns dictionary of rating periods and ratings details.
//
// https://developers.wargaming.net/reference/all/wot/ratings/types
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) RatingsTypes(ctx context.Context, realm Realm, options *wot.RatingsTypesOptions) (*wot.RatingsTypes, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.BattleType != nil {
			reqParam["battle_type"] = *options.BattleType
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wot.RatingsTypes
	err := service.client.getRequest(ctx, sectionWot, realm, "/ratings/types/", reqParam, &data, nil)
	return data, err
}
