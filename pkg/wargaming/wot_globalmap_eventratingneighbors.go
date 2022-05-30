package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapEventratingneighbors Method returns list of adjacent positions in event clan rating
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventratingneighbors
//
// clanId:
//     Clan ID. To get a clan ID, use the Clans method. Min value is 1.
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
//     Neighbors limit. Default is 10. Min value is 1. Maximum value: 99.
func (client *Client) WotGlobalmapEventratingneighbors(realm Realm, clanId int, eventId string, frontId string, fields []string, limit int) (*wot.GlobalmapEventratingneighbors, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": strconv.Itoa(clanId),
		"event_id": eventId,
		"front_id": frontId,
		"fields": strings.Join(fields, ","),
		"limit": strconv.Itoa(limit),
	}

	var result *wot.GlobalmapEventratingneighbors
	err := client.doGetDataRequest(realm, "/wot/globalmap/eventratingneighbors/", reqParam, &result)
	return result, err
}
