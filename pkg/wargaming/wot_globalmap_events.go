package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapEvents Method returns events information.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/events
//
// eventId:
//     Event ID
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Russian (by default)
// limit:
//     Page limit. Default is 5. Min value is 1. Maximum value: 20.
// pageNo:
//     Page number. Default is 1. Min value is 1.
// status:
//     Response with events filtered by status. Valid values:
//     
//     "PLANNED" &mdash; Upcoming event 
//     "ACTIVE" &mdash; Current event 
//     "FINISHED" &mdash; Event is over
func (client *Client) WotGlobalmapEvents(realm Realm, eventId string, fields []string, language string, limit int, pageNo int, status string) (*wot.GlobalmapEvents, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"event_id": eventId,
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"status": status,
	}

	var result *wot.GlobalmapEvents
	err := client.doGetDataRequest(realm, "/wot/globalmap/events/", reqParam, &result)
	return result, err
}
