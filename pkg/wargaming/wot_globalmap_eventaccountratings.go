package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapEventaccountratings Method returns account event rating.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventaccountratings
//
// eventId:
//     Event ID. To get an event ID, use the Events method.
//     Parameter is required.
// frontId:
//     Front ID. To get a front ID, use the Fronts method.
//     Parameter is required.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// inRating:
//     Get data only for accounts with rating. Default is 0. Valid values:
//     
//     "1" - Only with rating 
//     "0" - All (by default)
// limit:
//     Accounts limit. Default is 20. Min value is 10. Maximum value: 100.
// pageNo:
//     Page number. Default is 1. Min value is 1.
func (client *Client) WotGlobalmapEventaccountratings(realm Realm, eventId string, frontId string, fields []string, inRating int, limit int, pageNo int) (*wot.GlobalmapEventaccountratings, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"event_id": eventId,
		"front_id": frontId,
		"fields": strings.Join(fields, ","),
		"in_rating": strconv.Itoa(inRating),
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
	}

	var result *wot.GlobalmapEventaccountratings
	err := client.doGetDataRequest(realm, "/wot/globalmap/eventaccountratings/", reqParam, &result)
	return result, err
}