// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapProvinces returns information about the Global Map provinces.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/provinces
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// frontId:
//     Front ID. To get a front ID, use the Fronts method.
func (service *WotService) GlobalmapProvinces(ctx context.Context, realm Realm, frontId string, options *wot.GlobalmapProvincesOptions) ([]*wot.GlobalmapProvinces, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"front_id": frontId,
	}

	if options != nil {
		if options.ArenaId != nil {
			reqParam["arena_id"] = *options.ArenaId
		}
		if options.DailyRevenueGte != nil {
			reqParam["daily_revenue_gte"] = strconv.Itoa(*options.DailyRevenueGte)
		}
		if options.DailyRevenueLte != nil {
			reqParam["daily_revenue_lte"] = strconv.Itoa(*options.DailyRevenueLte)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.LandingType != nil {
			reqParam["landing_type"] = *options.LandingType
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.OrderBy != nil {
			reqParam["order_by"] = *options.OrderBy
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.PrimeHour != nil {
			reqParam["prime_hour"] = strconv.Itoa(*options.PrimeHour)
		}
		if options.ProvinceId != nil {
			reqParam["province_id"] = strings.Join(options.ProvinceId, ",")
		}
	}

	var data []*wot.GlobalmapProvinces
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/provinces/", reqParam, &data, &metaData)
	return data, metaData, err
}
