package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strconv"
	"strings"
)

// RatingsTop Method returns list of top players by specified parameter.
//
// https://developers.wargaming.net/reference/all/wot/ratings/top
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// rankField:
//     Rating category
// type_:
//     Rating period. For valid values, check the Types of ratings method.
func (service *WotService) RatingsTop(ctx context.Context, realm Realm, rankField string, type_ string, options *wot.RatingsTopOptions) (*wot.RatingsTop, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"rank_field": rankField,
		"type":       type_,
	}
	if options != nil {
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Date != nil {
			reqParam["date"] = strconv.FormatInt(options.Date.Unix(), 10)
		}
		if options.BattleType != nil {
			reqParam["battle_type"] = *options.BattleType
		}
	}

	var data *wot.RatingsTop
	err := service.client.getRequest(ctx, sectionWot, realm, "/ratings/top/", reqParam, &data)
	return data, err
}
