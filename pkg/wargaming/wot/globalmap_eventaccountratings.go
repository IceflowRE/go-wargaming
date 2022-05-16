package wot

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type GlobalmapEventaccountratings struct {
	// Link to Front
	Url string `json:"url,omitempty"`
	// Date when account data were updated
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Rating changes during previous turn
	RankDelta int `json:"rank_delta,omitempty"`
	// Current rating
	Rank int `json:"rank,omitempty"`
	// Front ID
	FrontId string `json:"front_id,omitempty"`
	// Amount of Fame Points needed to improve personal award
	FamePointsToImproveAward int `json:"fame_points_to_improve_award,omitempty"`
	// Total Fame Points
	FamePoints int `json:"fame_points,omitempty"`
	// Event ID
	EventId string `json:"event_id,omitempty"`
	// Clan rating
	ClanRank int `json:"clan_rank,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Battles to fight in a current clan to get clan award for the season
	BattlesToAward int `json:"battles_to_award,omitempty"`
	// Battles fought for current clan
	Battles int `json:"battles,omitempty"`
	// Award level
	AwardLevel string `json:"award_level,omitempty"`
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
} 
