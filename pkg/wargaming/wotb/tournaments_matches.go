package wotb

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type TournamentsMatches struct {
	// Winner Team ID
	WinnerTeamId int `json:"winner_team_id,omitempty"`
	// Tournament ID
	TournamentId int `json:"tournament_id,omitempty"`
	// Team 2 score
	Team2Score int `json:"team_2_score,omitempty"`
	// Team 2 ID
	Team2Id int `json:"team_2_id,omitempty"`
	// Team 1 score
	Team1Score int `json:"team_1_score,omitempty"`
	// Team 1 ID
	Team1Id int `json:"team_1_id,omitempty"`
	// Match state. Valid values:
	// 
	// "waiting_results" - Match is still in progress and there are no final results 
	// "got_results" - Match was finished and there are final results 
	// "canceled" - Match was cancelled 
	// "upcoming" - Already scheduled match state
	State string `json:"state,omitempty"`
	// Match start time
	StartTime wgnTime.UnixTime `json:"start_time,omitempty"`
	// Stage ID
	StageId int `json:"stage_id,omitempty"`
	// Tour number of the current match
	Round int `json:"round,omitempty"`
	// ID of the following match for the winner
	NextMatchForWinner string `json:"next_match_for_winner,omitempty"`
	// ID of the following match for the loser
	NextMatchForLooser string `json:"next_match_for_looser,omitempty"`
	// Match ID
	Id string `json:"id,omitempty"`
	// Group ID
	GroupId int `json:"group_id,omitempty"`
} 
