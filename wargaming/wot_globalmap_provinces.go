package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapProvinces Method returns information about the Global Map provinces.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/provinces
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// frontId:
//     Front ID. To get a front ID, use the Fronts method.
func (service *wotService) GlobalmapProvinces(ctx context.Context, realm Realm, frontId string, options *wot.GlobalmapProvincesOptions) ([]*wot.GlobalmapProvinces, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"front_id": frontId,
	}
	if options != nil {
		if options.ProvinceId != nil {
			reqParam["province_id"] = strings.Join(options.ProvinceId, ",")
		}
		if options.PrimeHour != nil {
			reqParam["prime_hour"] = strconv.Itoa(*options.PrimeHour)
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.OrderBy != nil {
			reqParam["order_by"] = *options.OrderBy
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.LandingType != nil {
			reqParam["landing_type"] = *options.LandingType
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.DailyRevenueLte != nil {
			reqParam["daily_revenue_lte"] = strconv.Itoa(*options.DailyRevenueLte)
		}
		if options.DailyRevenueGte != nil {
			reqParam["daily_revenue_gte"] = strconv.Itoa(*options.DailyRevenueGte)
		}
		if options.ArenaId != nil {
			reqParam["arena_id"] = *options.ArenaId
		}
	}

	var data []*wot.GlobalmapProvinces
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/provinces/", reqParam, &data)
	return data, err
}
