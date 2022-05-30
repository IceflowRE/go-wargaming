package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapEventaccountratingneighbors Method returns adjacent position in account event rating.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventaccountratingneighbors
//
// accountId:
//     Account ID. Min value is 1.
//     Parameter is required.
// eventId:
//     Event ID. To get an event ID, use the Events method.
//     Parameter is required.
// frontId:
//     Front ID. To get a front ID, use the Fronts method.
//     Parameter is required.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// limit:
//     Clans limit. Default is 20. Min value is 10. Maximum value: 100.
// neighboursCount:
//     How many neighbors to show next to the account. Default is 3. Min value is 1. Maximum value: 99.
// pageNo:
//     Page number. Default is 1. Min value is 1.
func (client *Client) WotGlobalmapEventaccountratingneighbors(realm Realm, accountId int, eventId string, frontId string, fields []string, limit int, neighboursCount int, pageNo int) (*wot.GlobalmapEventaccountratingneighbors, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"event_id": eventId,
		"front_id": frontId,
		"fields": strings.Join(fields, ","),
		"limit": strconv.Itoa(limit),
		"neighbours_count": strconv.Itoa(neighboursCount),
		"page_no": strconv.Itoa(pageNo),
	}

	var result *wot.GlobalmapEventaccountratingneighbors
	err := client.doGetDataRequest(realm, "/wot/globalmap/eventaccountratingneighbors/", reqParam, &result)
	return result, err
}
