package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wowp"
	"strconv"
	"strings"
)

// RatingsTop Method returns list of top players by specified parameter.
//
// https://developers.wargaming.net/reference/all/wowp/ratings/top
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
// rankField:
//     Rating category
// type_:
//     Rating period. For valid values, check the Types of ratings method.
func (service *wowpService) RatingsTop(ctx context.Context, realm Realm, rankField string, type_ string, options *wowp.RatingsTopOptions) ([]*wowp.RatingsTop, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
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
	}

	var data []*wowp.RatingsTop
	err := service.client.getRequest(ctx, sectionWowp, realm, "/ratings/top/", reqParam, &data)
	return data, err
}
