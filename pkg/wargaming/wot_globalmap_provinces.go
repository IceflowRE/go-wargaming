package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapProvinces Method returns information about the Global Map provinces.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/provinces
//
// arena_id:
//     Map ID
// daily_revenue_gte:
//     Search for provinces with daily income equal to or more than the value
// daily_revenue_lte:
//     Search for provinces with daily income equal to or less than the value
// front_id:
//     Front ID. To get a front ID, use the Fronts method.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// landing_type:
//     Search for provinces by landing type. Valid values:
//     
//     "null" &mdash; auctions disabled 
//     "auction" &mdash; auction 
//     "tournament" &mdash; landing tournament
// language:
//     Language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Russian (by default)
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// order_by:
//     Sorting. Valid values:
//     
//     "province_id" &mdash; by province name 
//     "-province_id" &mdash; by province name in reverse order 
//     "daily_revenue" &mdash; by province income 
//     "-daily_revenue" &mdash; by province income in reverse order 
//     "prime_hour" &mdash; by Prime Time 
//     "-prime_hour" &mdash; by Prime Time in reverse order
// page_no:
//     Page number. Default is 1. Min value is 1.
// prime_hour:
//     Search for provinces with the value of Prime Time start hour. Values available: from 0 to 23. Maximum value: 23.
// province_id:
//     Filter by the list of province IDs. Maximum limit: 100.
func (client *Client) WotGlobalmapProvinces(realm Realm, arenaId string, dailyRevenueGte int, dailyRevenueLte int, frontId string, fields []string, landingType string, language string, limit int, orderBy string, pageNo int, primeHour int, provinceId []string) (*wot.GlobalmapProvinces, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"arena_id": arenaId,
		"daily_revenue_gte": strconv.Itoa(dailyRevenueGte),
		"daily_revenue_lte": strconv.Itoa(dailyRevenueLte),
		"front_id": frontId,
		"fields": strings.Join(fields, ","),
		"landing_type": landingType,
		"language": language,
		"limit": strconv.Itoa(limit),
		"order_by": orderBy,
		"page_no": strconv.Itoa(pageNo),
		"prime_hour": strconv.Itoa(primeHour),
		"province_id": strings.Join(provinceId, ","),
	}

	var result *wot.GlobalmapProvinces
	err := client.doGetDataRequest(realm, "/wot/globalmap/provinces/", reqParam, &result)
	return result, err
}
