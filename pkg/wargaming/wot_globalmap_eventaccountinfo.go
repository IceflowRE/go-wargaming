package wargaming

import (
	"strings"
	"strconv"
)

type WotGlobalmapEventaccountinfo struct {
	// Account information by events and Fronts
	Events struct {
		// Account ID
		AccountId int `json:"account_id,omitempty"`
		// Award level
		AwardLevel string `json:"award_level,omitempty"`
		// Battles fought for current clan
		Battles int `json:"battles,omitempty"`
		// Battles to fight in a current clan to get clan award for the season
		BattlesToAward int `json:"battles_to_award,omitempty"`
		// Clan ID
		ClanId int `json:"clan_id,omitempty"`
		// Clan rating
		ClanRank int `json:"clan_rank,omitempty"`
		// Event ID
		EventId string `json:"event_id,omitempty"`
		// Total Fame Points
		FamePoints int `json:"fame_points,omitempty"`
		// Change of Fame Points since last turn calculation
		FamePointsSinceTurn int `json:"fame_points_since_turn,omitempty"`
		// Amount of Fame Points needed to improve personal award
		FamePointsToImproveAward int `json:"fame_points_to_improve_award,omitempty"`
		// Front ID
		FrontId string `json:"front_id,omitempty"`
		// Current rating
		Rank int `json:"rank,omitempty"`
		// Rating changes during previous turn
		RankDelta int `json:"rank_delta,omitempty"`
		// Date when account data were updated
		UpdatedAt UnixTime `json:"updated_at,omitempty"`
		// Link to Front
		Url string `json:"url,omitempty"`
	} `json:"events,omitempty"`
}

// WotGlobalmapEventaccountinfo Method returns player's statistics for a specific event
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventaccountinfo
//
// account_id:
//     Account ID. Min value is 1.
// clan_id:
//     Clan ID. To get a clan ID, use the Clans method. Min value is 1.
// event_id:
//     Event ID. To get an event ID, use the Events method.
// front_id:
//     Front ID. To get a front ID, use the Fronts method. Maximum limit: 10.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WotGlobalmapEventaccountinfo(realm Realm, accountId int, clanId int, eventId string, frontId []string, fields []string) (*WotGlobalmapEventaccountinfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"clan_id": strconv.Itoa(clanId),
		"event_id": eventId,
		"front_id": strings.Join(frontId, ","),
		"fields": strings.Join(fields, ","),
	}

	var result *WotGlobalmapEventaccountinfo
	err := client.doGetDataRequest(realm, "/wot/globalmap/eventaccountinfo/", reqParam, &result)
	return result, err
}
