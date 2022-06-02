package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wot"
	"strconv"
	"strings"
)

// RatingsNeighbors Method returns list of adjacent positions in specified rating.
//
// https://developers.wargaming.net/reference/all/wot/ratings/neighbors
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accountId:
//     Player account ID
// rankField:
//     Rating category
// type_:
//     Rating period. For valid values, check the Types of ratings method.
func (service *wotService) RatingsNeighbors(ctx context.Context, realm Realm, accountId int, rankField string, type_ string, options *wot.RatingsNeighborsOptions) (*wot.RatingsNeighbors, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"rank_field": rankField,
		"type":       type_,
	}
	if options != nil {
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

	var data *wot.RatingsNeighbors
	err := service.client.getRequest(ctx, sectionWot, realm, "/ratings/neighbors/", reqParam, &data)
	return data, err
}
