// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// RatingsDates returns dates with available rating data.
//
// https://developers.wargaming.net/reference/all/wot/ratings/dates
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// typ:
//     Rating period. For valid values, check the Types of ratings method.
func (service *WotService) RatingsDates(ctx context.Context, realm Realm, typ string, options *wot.RatingsDatesOptions) (*wot.RatingsDates, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"type": typ,
	}

	if options != nil {
		if options.AccountId != nil {
			reqParam["account_id"] = internal.SliceIntToString(options.AccountId, ",")
		}
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

	var data *wot.RatingsDates
	err := service.client.getRequest(ctx, sectionWot, realm, "/ratings/dates/", reqParam, &data, nil)
	return data, err
}
