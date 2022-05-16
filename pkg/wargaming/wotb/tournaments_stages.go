package wotb

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type TournamentsStages struct {
	// Number of victories to win the match
	VictoryLimit int `json:"victory_limit,omitempty"`
	// Bracket type of the stage. Valid values:
	// 
	// 
	// RR — round robin
	// 
	// 
	// SE — single elimination
	// 
	// 
	// DE — double elimination
	Type_ string `json:"type,omitempty"`
	// Tournament ID
	TournamentId int `json:"tournament_id,omitempty"`
	// Stage name
	Title string `json:"title,omitempty"`
	// Stage state. Valid values:
	// 
	// 
	// draft — stage created as draft
	// 
	// 
	// groups_ready —  stage has completed groups
	// 
	// 
	// schedule_ready — stage has a finalized schedule
	// 
	// 
	// complete — stage completed
	State string `json:"state,omitempty"`
	// Stage start day and time
	StartAt wgnTime.UnixTime `json:"start_at,omitempty"`
	// Stage ID
	StageId int `json:"stage_id,omitempty"`
	// Number of tours in the stage
	RoundsCount int `json:"rounds_count,omitempty"`
	// List of stage tours (numbers of tours)
	Rounds []int `json:"rounds,omitempty"`
	// The allowed minimum vehicle tier for each player in the team
	MinTier int `json:"min_tier,omitempty"`
	// The allowed maximum vehicle tier for each player in the team
	MaxTier int `json:"max_tier,omitempty"`
	// Number of groups in the stage
	GroupsCount int `json:"groups_count,omitempty"`
	// Groups info for current stage
	Groups struct {
		// Group order number
		GroupOrder int `json:"group_order,omitempty"`
		// Group ID
		GroupId int `json:"group_id,omitempty"`
	} `json:"groups,omitempty"`
	// Stage end day and time
	EndAt wgnTime.UnixTime `json:"end_at,omitempty"`
	// Stage description
	Description string `json:"description,omitempty"`
	// Number of battles in match
	BattleLimit int `json:"battle_limit,omitempty"`
} 
