// Auto generated file!

package wotb

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type TournamentsMatchesOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Group ID that can be retrieved from the Tournaments Stages method. Maximum limit: 10.
	GroupId []int `json:"group_id,omitempty"`
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
	// Tour number. Maximum limit: 10.
	RoundNumber []int `json:"round_number,omitempty"`
	// Team ID. Maximum limit: 10.
	TeamId []int `json:"team_id,omitempty"`
}

type TournamentsMatches struct {
	// Group ID
	GroupId *int `json:"group_id,omitempty"`
	// Match ID
	Id *string `json:"id,omitempty"`
	// ID of the following match for the loser
	NextMatchForLooser *string `json:"next_match_for_looser,omitempty"`
	// ID of the following match for the winner
	NextMatchForWinner *string `json:"next_match_for_winner,omitempty"`
	// Tour number of the current match
	Round *int `json:"round,omitempty"`
	// Stage ID
	StageId *int `json:"stage_id,omitempty"`
	// Match start time
	StartTime *wgnTime.UnixTime `json:"start_time,omitempty"`
	// Match state. Valid values:
	//
	// "waiting_results" - Match is still in progress and there are no final results
	// "got_results" - Match was finished and there are final results
	// "canceled" - Match was cancelled
	// "upcoming" - Already scheduled match state
	State *string `json:"state,omitempty"`
	// Team 1 ID
	Team1Id *int `json:"team_1_id,omitempty"`
	// Team 1 score
	Team1Score *int `json:"team_1_score,omitempty"`
	// Team 2 ID
	Team2Id *int `json:"team_2_id,omitempty"`
	// Team 2 score
	Team2Score *int `json:"team_2_score,omitempty"`
	// Tournament ID
	TournamentId *int `json:"tournament_id,omitempty"`
	// Winner Team ID
	WinnerTeamId *int `json:"winner_team_id,omitempty"`
}
