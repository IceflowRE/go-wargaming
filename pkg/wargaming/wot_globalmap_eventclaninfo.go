package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapEventclaninfo Method returns clan's statistics for a specific event.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventclaninfo
//
// clanId:
//     Clan ID. To get a clan ID, use the Clans method. Min value is 1.
//     Parameter is required.
// eventId:
//     Event ID. To get an event ID, use the Events method.
//     Parameter is required.
// frontId:
//     Front ID. To get a front ID, use the Fronts method. Maximum limit: 10.
//     Parameter is required.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WotGlobalmapEventclaninfo(realm Realm, clanId int, eventId string, frontId []string, fields []string) (*wot.GlobalmapEventclaninfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": strconv.Itoa(clanId),
		"event_id": eventId,
		"front_id": strings.Join(frontId, ","),
		"fields": strings.Join(fields, ","),
	}

	var result *wot.GlobalmapEventclaninfo
	err := client.doGetDataRequest(realm, "/wot/globalmap/eventclaninfo/", reqParam, &result)
	return result, err
}
