// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// RatingsTop returns list of top players by specified parameter.
//
// https://developers.wargaming.net/reference/all/wot/ratings/top
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// rankField:
//     Rating category
// typ:
//     Rating period. For valid values, check the Types of ratings method.
func (service *WotService) RatingsTop(ctx context.Context, realm Realm, rankField string, typ string, options *wot.RatingsTopOptions) (*wot.RatingsTop, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"rank_field": rankField,
		"type":       typ,
	}

	if options != nil {
		if options.BattleType != nil {
			reqParam["battle_type"] = *options.BattleType
		}
		if options.Date != nil {
			reqParam["date"] = strconv.FormatInt(options.Date.Unix(), 10)
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
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
	}

	var data *wot.RatingsTop
	err := service.client.getRequest(ctx, sectionWot, realm, "/ratings/top/", reqParam, &data, nil)
	return data, err
}
