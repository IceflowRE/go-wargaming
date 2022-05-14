package wargaming

import (
	"strings"
	"strconv"
)

type WotGlobalmapEventratingneighbors struct {
	// Award level
	AwardLevel string `json:"award_level,omitempty"`
	// Battle Fame Points
	BattleFamePoints int `json:"battle_fame_points,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Clan color
	Color string `json:"color,omitempty"`
	// Amount of Fame Points needed to improve personal award
	FamePointsToImproveAward int `json:"fame_points_to_improve_award,omitempty"`
	// Clan name
	Name string `json:"name,omitempty"`
	// Current rating
	Rank int `json:"rank,omitempty"`
	// Rating changes during previous turn
	RankDelta int `json:"rank_delta,omitempty"`
	// Clan tag
	Tag string `json:"tag,omitempty"`
	// Fame Points for completing a clan task
	TaskFamePoints int `json:"task_fame_points,omitempty"`
	// Total Fame Points
	TotalFamePoints int `json:"total_fame_points,omitempty"`
	// Date of rating calculation
	UpdatedAt UnixTime `json:"updated_at,omitempty"`
}

// WotGlobalmapEventratingneighbors Method returns list of adjacent positions in event clan rating
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventratingneighbors
//
// clan_id:
//     Clan ID. To get a clan ID, use the Clans method. Min value is 1.
// event_id:
//     Event ID. To get an event ID, use the Events method.
// front_id:
//     Front ID. To get a front ID, use the Fronts method.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// limit:
//     Neighbors limit. Default is 10. Min value is 1. Maximum value: 99.
func (client *Client) WotGlobalmapEventratingneighbors(realm Realm, clanId int, eventId string, frontId string, fields []string, limit int) (*WotGlobalmapEventratingneighbors, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": strconv.Itoa(clanId),
		"event_id": eventId,
		"front_id": frontId,
		"fields": strings.Join(fields, ","),
		"limit": strconv.Itoa(limit),
	}

	var result *WotGlobalmapEventratingneighbors
	err := client.doGetDataRequest(realm, "/wot/globalmap/eventratingneighbors/", reqParam, &result)
	return result, err
}
