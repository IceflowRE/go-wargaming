package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wowp"
	"strconv"
	"strings"
)

// RatingsNeighbors Method returns list of adjacent positions in specified rating.
//
// https://developers.wargaming.net/reference/all/wowp/ratings/neighbors
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
// accountId:
//     Player account ID
// rankField:
//     Rating category
// type_:
//     Rating period. For valid values, check the Types of ratings method.
func (service *wowpService) RatingsNeighbors(ctx context.Context, realm Realm, accountId int, rankField string, type_ string, options *wowp.RatingsNeighborsOptions) ([]*wowp.RatingsNeighbors, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
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
	}

	var data []*wowp.RatingsNeighbors
	err := service.client.getRequest(ctx, sectionWowp, realm, "/ratings/neighbors/", reqParam, &data)
	return data, err
}
