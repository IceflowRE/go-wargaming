// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wowp"
	"strconv"
	"strings"
)

// RatingsNeighbors returns list of adjacent positions in specified rating.
//
// https://developers.wargaming.net/reference/all/wowp/ratings/neighbors
//
// realm:
//     Valid realms: RealmEu, RealmNa
// accountId:
//     Player account ID
// rankField:
//     Rating category
// typ:
//     Rating period. For valid values, check the Types of ratings method.
func (service *WowpService) RatingsNeighbors(ctx context.Context, realm Realm, accountId int, rankField string, typ string, options *wowp.RatingsNeighborsOptions) ([]*wowp.RatingsNeighbors, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"rank_field": rankField,
		"type":       typ,
	}

	if options != nil {
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

	var data []*wowp.RatingsNeighbors
	err := service.client.getRequest(ctx, sectionWowp, realm, "/ratings/neighbors/", reqParam, &data, nil)
	return data, err
}
