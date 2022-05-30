package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapEventaccountinfo Method returns player's statistics for a specific event
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventaccountinfo
//
// accountId:
//     Account ID. Min value is 1.
//     Parameter is required.
// clanId:
//     Clan ID. To get a clan ID, use the Clans method. Min value is 1.
// eventId:
//     Event ID. To get an event ID, use the Events method.
//     Parameter is required.
// frontId:
//     Front ID. To get a front ID, use the Fronts method. Maximum limit: 10.
//     Parameter is required.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WotGlobalmapEventaccountinfo(realm Realm, accountId int, clanId int, eventId string, frontId []string, fields []string) (*wot.GlobalmapEventaccountinfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"clan_id": strconv.Itoa(clanId),
		"event_id": eventId,
		"front_id": strings.Join(frontId, ","),
		"fields": strings.Join(fields, ","),
	}

	var result *wot.GlobalmapEventaccountinfo
	err := client.doGetDataRequest(realm, "/wot/globalmap/eventaccountinfo/", reqParam, &result)
	return result, err
}