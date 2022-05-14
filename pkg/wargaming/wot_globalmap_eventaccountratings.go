package wargaming

import (
	"strconv"
	"strings"
)

type WotGlobalmapEventaccountratings struct {
	// Player account ID
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
}

// WotGlobalmapEventaccountratings Method returns account event rating.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventaccountratings
//
// event_id:
//     Event ID. To get an event ID, use the Events method.
// front_id:
//     Front ID. To get a front ID, use the Fronts method.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// in_rating:
//     Get data only for accounts with rating. Default is 0. Valid values:
//     
//     "1" &mdash; Only with rating 
//     "0" &mdash; All (by default)
// limit:
//     Accounts limit. Default is 20. Min value is 10. Maximum value: 100.
// page_no:
//     Page number. Default is 1. Min value is 1.
func (client *Client) WotGlobalmapEventaccountratings(realm Realm, eventId string, frontId string, fields []string, inRating int, limit int, pageNo int) (*WotGlobalmapEventaccountratings, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
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

	var result *WotGlobalmapEventaccountratings
	err := client.doGetDataRequest(realm, "/wot/globalmap/eventaccountratings/", reqParam, &result)
	return result, err
}
