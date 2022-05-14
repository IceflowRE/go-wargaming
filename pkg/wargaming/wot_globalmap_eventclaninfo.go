package wargaming

import (
	"strings"
	"strconv"
)

type WotGlobalmapEventclaninfo struct {
	// Clan info by events and Fronts
	Events struct {
		// Battle Fame Points
		BattleFamePoints int `json:"battle_fame_points,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Event ID
		EventId string `json:"event_id,omitempty"`
		// Total Fame Points
		FamePoints int `json:"fame_points,omitempty"`
		// Change of Fame Points since last turn calculation
		FamePointsSinceTurn int `json:"fame_points_since_turn,omitempty"`
		// Front ID
		FrontId string `json:"front_id,omitempty"`
		// Current rating
		Rank int `json:"rank,omitempty"`
		// Rating changes during previous turn
		RankDelta int `json:"rank_delta,omitempty"`
		// Fame Points for completing a clan task
		TaskFamePoints int `json:"task_fame_points,omitempty"`
		// Link to Front
		Url string `json:"url,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
	} `json:"events,omitempty"`
}

// WotGlobalmapEventclaninfo Method returns clan's statistics for a specific event.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventclaninfo
//
// clan_id:
//     Clan ID. To get a clan ID, use the Clans method. Min value is 1.
// event_id:
//     Event ID. To get an event ID, use the Events method.
// front_id:
//     Front ID. To get a front ID, use the Fronts method. Maximum limit: 10.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WotGlobalmapEventclaninfo(realm Realm, clanId int, eventId string, frontId []string, fields []string) (*WotGlobalmapEventclaninfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": strconv.Itoa(clanId),
		"event_id": eventId,
		"front_id": strings.Join(frontId, ","),
		"fields": strings.Join(fields, ","),
	}

	var result *WotGlobalmapEventclaninfo
	err := client.doGetDataRequest(realm, "/wot/globalmap/eventclaninfo/", reqParam, &result)
	return result, err
}
