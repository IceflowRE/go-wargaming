// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type GlobalmapEventaccountratingsOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Get data only for accounts with rating. Default is 0. Valid values:
	//
	// "1" - Only with rating
	// "0" - All (by default)
	InRating *int `json:"in_rating,omitempty"`
	// Accounts limit. Default is 20. Min value is 10. Maximum value: 100.
	Limit *int `json:"limit,omitempty"`
	// Page number. Default is 1. Min value is 1.
	PageNo *int `json:"page_no,omitempty"`
}

type GlobalmapEventaccountratings struct {
	// Player account ID
	AccountId *int `json:"account_id,omitempty"`
	// Award level
	AwardLevel *string `json:"award_level,omitempty"`
	// Battles fought for current clan
	Battles *int `json:"battles,omitempty"`
	// Battles to fight in a current clan to get clan award for the season
	BattlesToAward *int `json:"battles_to_award,omitempty"`
	// Clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Clan rating
	ClanRank *int `json:"clan_rank,omitempty"`
	// Event ID
	EventId *string `json:"event_id,omitempty"`
	// Total Fame Points
	FamePoints *int `json:"fame_points,omitempty"`
	// Amount of Fame Points needed to improve personal award
	FamePointsToImproveAward *int `json:"fame_points_to_improve_award,omitempty"`
	// Front ID
	FrontId *string `json:"front_id,omitempty"`
	// Current rating
	Rank *int `json:"rank,omitempty"`
	// Rating changes during previous turn
	RankDelta *int `json:"rank_delta,omitempty"`
	// Date when account data were updated
	UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Link to Front
	Url *string `json:"url,omitempty"`
}
