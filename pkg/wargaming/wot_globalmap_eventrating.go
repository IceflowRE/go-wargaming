package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapEventrating Method returns event clan rating
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventrating
//
// event_id:
//     Event ID. To get an event ID, use the Events method.
// front_id:
//     Front ID. To get a front ID, use the Fronts method.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// limit:
//     Clans limit. Default is 10. Min value is 1. Maximum value: 100.
// page_no:
//     Page number. Default is 1. Min value is 1.
func (client *Client) WotGlobalmapEventrating(realm Realm, eventId string, frontId string, fields []string, limit int, pageNo int) (*wot.GlobalmapEventrating, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"event_id": eventId,
		"front_id": frontId,
		"fields": strings.Join(fields, ","),
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
	}

	var result *wot.GlobalmapEventrating
	err := client.doGetDataRequest(realm, "/wot/globalmap/eventrating/", reqParam, &result)
	return result, err
}
