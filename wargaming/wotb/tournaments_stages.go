// Auto generated file!

package wotb

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type TournamentsStagesOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "cs" - Čeština
	// "de" - Deutsch
	// "es" - Español
	// "fr" - Français
	// "pl" - Polski
	// "tr" - Türkçe
	Language *string `json:"language,omitempty"`
	// Number of returned entries. Default is 10. Min value is 1. Maximum value: 25.
	Limit *int `json:"limit,omitempty"`
	// Result page number. Default is 1. Min value is 1.
	PageNo *int `json:"page_no,omitempty"`
}

type TournamentsStages struct {
	// Number of battles in match
	BattleLimit *int `json:"battle_limit,omitempty"`
	// Stage description
	Description *string `json:"description,omitempty"`
	// Stage end day and time
	EndAt *wgnTime.UnixTime `json:"end_at,omitempty"`
	// Groups info for current stage
	Groups []*struct {
		// Group ID
		GroupId *int `json:"group_id,omitempty"`
		// Group order number
		GroupOrder *int `json:"group_order,omitempty"`
	} `json:"groups,omitempty"`
	// Number of groups in the stage
	GroupsCount *int `json:"groups_count,omitempty"`
	// The allowed maximum vehicle tier for each player in the team
	MaxTier *int `json:"max_tier,omitempty"`
	// The allowed minimum vehicle tier for each player in the team
	MinTier *int `json:"min_tier,omitempty"`
	// List of stage tours (numbers of tours)
	Rounds []int `json:"rounds,omitempty"`
	// Number of tours in the stage
	RoundsCount *int `json:"rounds_count,omitempty"`
	// Stage ID
	StageId *int `json:"stage_id,omitempty"`
	// Stage start day and time
	StartAt *wgnTime.UnixTime `json:"start_at,omitempty"`
	// Stage state. Valid values:
	//
	//
	// draft - stage created as draft
	//
	//
	// groups_ready -  stage has completed groups
	//
	//
	// schedule_ready - stage has a finalized schedule
	//
	//
	// complete - stage completed
	//
	State *string `json:"state,omitempty"`
	// Stage name
	Title *string `json:"title,omitempty"`
	// Tournament ID
	TournamentId *int `json:"tournament_id,omitempty"`
	// Bracket type of the stage. Valid values:
	//
	//
	// RR - round robin
	//
	//
	// SE - single elimination
	//
	//
	// DE - double elimination
	//
	Type *string `json:"type,omitempty"`
	// Number of victories to win the match
	VictoryLimit *int `json:"victory_limit,omitempty"`
}
