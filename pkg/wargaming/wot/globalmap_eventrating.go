package wot

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type GlobalmapEventrating struct {
	// Date of rating calculation
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Total Fame Points
	TotalFamePoints int `json:"total_fame_points,omitempty"`
	// Fame Points for completing a clan task
	TaskFamePoints int `json:"task_fame_points,omitempty"`
	// Clan tag
	Tag string `json:"tag,omitempty"`
	// Rating changes during previous turn
	RankDelta int `json:"rank_delta,omitempty"`
	// Current rating
	Rank int `json:"rank,omitempty"`
	// Clan name
	Name string `json:"name,omitempty"`
	// Amount of Fame Points needed to improve personal award
	FamePointsToImproveAward int `json:"fame_points_to_improve_award,omitempty"`
	// Clan color
	Color string `json:"color,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Battle Fame Points
	BattleFamePoints int `json:"battle_fame_points,omitempty"`
	// Award level
	AwardLevel string `json:"award_level,omitempty"`
} 
