package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strings"
)

// RatingsDates Method returns dates with available rating data.
//
// https://developers.wargaming.net/reference/all/wot/ratings/dates
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// type_:
//     Rating period. For valid values, check the Types of ratings method.
func (service *WotService) RatingsDates(ctx context.Context, realm Realm, type_ string, options *wot.RatingsDatesOptions) (*wot.RatingsDates, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"type": type_,
	}
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
		if options.AccountId != nil {
			reqParam["account_id"] = internal.SliceIntToString(options.AccountId, ",")
		}
	}

	var data *wot.RatingsDates
	err := service.client.getRequest(ctx, sectionWot, realm, "/ratings/dates/", reqParam, &data)
	return data, err
}
