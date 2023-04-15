// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// RatingsNeighbors returns list of adjacent positions in specified rating.
//
// https://developers.wargaming.net/reference/all/wot/ratings/neighbors
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Player account ID
// rankField:
//     Rating category
// typ:
//     Rating period. For valid values, check the Types of ratings method.
func (service *WotService) RatingsNeighbors(ctx context.Context, realm Realm, accountId int, rankField string, typ string, options *wot.RatingsNeighborsOptions) (*wot.RatingsNeighbors, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
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
	}

	var data *wot.RatingsNeighbors
	err := service.client.getRequest(ctx, sectionWot, realm, "/ratings/neighbors/", reqParam, &data, nil)
	return data, err
}
