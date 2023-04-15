// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
	"strconv"
	"strings"
)

// RatingsTop returns list of top players by specified parameter.
//
// https://developers.wargaming.net/reference/all/wowp/ratings/top
//
// realm:
//     Valid realms: RealmEu, RealmNa
// rankField:
//     Rating category
// typ:
//     Rating period. For valid values, check the Types of ratings method.
func (service *WowpService) RatingsTop(ctx context.Context, realm Realm, rankField string, typ string, options *wowp.RatingsTopOptions) ([]*wowp.RatingsTop, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
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
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
	}

	var data []*wowp.RatingsTop
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/ratings/top/", reqParam, &data, &metaData)
	return data, metaData, err
}
