package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
	"strconv"
)

type WotbTournamentsMatches struct {
	// Group ID
	GroupId int `json:"group_id,omitempty"`
	// Match ID
	Id string `json:"id,omitempty"`
	// ID of the following match for the loser
	NextMatchForLooser string `json:"next_match_for_looser,omitempty"`
	// ID of the following match for the winner
	NextMatchForWinner string `json:"next_match_for_winner,omitempty"`
	// Tour number of the current match
	Round int `json:"round,omitempty"`
	// Stage ID
	StageId int `json:"stage_id,omitempty"`
	// Match start time
	StartTime UnixTime `json:"start_time,omitempty"`
	// Match state. Valid values:
	// 
	// "waiting_results" &mdash; Match is still in progress and there are no final results 
	// "got_results" &mdash; Match was finished and there are final results 
	// "canceled" &mdash; Match was cancelled 
	// "upcoming" &mdash; Already scheduled match state
	State string `json:"state,omitempty"`
	// Team 1 ID
	Team1Id int `json:"team_1_id,omitempty"`
	// Team 1 score
	Team1Score int `json:"team_1_score,omitempty"`
	// Team 2 ID
	Team2Id int `json:"team_2_id,omitempty"`
	// Team 2 score
	Team2Score int `json:"team_2_score,omitempty"`
	// Tournament ID
	TournamentId int `json:"tournament_id,omitempty"`
	// Winner Team ID
	WinnerTeamId int `json:"winner_team_id,omitempty"`
}

// WotbTournamentsMatches Method returns information about matches of a particular tournament and stage.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/matches
//
// stage_id:
//     Stage ID that can be retrieved from the Tournaments Stages method.
// tournament_id:
//     Tournament ID that can be retrieved from the Tournaments list method.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// group_id:
//     Group ID that can be retrieved from the Tournaments Stages method. Maximum limit: 10.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Русский (by default)
// limit:
//     Number of returned entries. Default is 10. Min value is 1. Maximum value: 25.
// page_no:
//     Result page number. Default is 1. Min value is 1.
// round_number:
//     Tour number. Maximum limit: 10.
// team_id:
//     Team ID. Maximum limit: 10.
func (client *Client) WotbTournamentsMatches(realm Realm, stageId int, tournamentId int, fields []string, groupId []int, language string, limit int, pageNo int, roundNumber []int, teamId []int) (*WotbTournamentsMatches, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"stage_id": strconv.Itoa(stageId),
		"tournament_id": strconv.Itoa(tournamentId),
		"fields": strings.Join(fields, ","),
		"group_id": utils.SliceIntToString(groupId, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"round_number": utils.SliceIntToString(roundNumber, ","),
		"team_id": utils.SliceIntToString(teamId, ","),
	}

	var result *WotbTournamentsMatches
	err := client.doGetDataRequest(realm, "/wotb/tournaments/matches/", reqParam, &result)
	return result, err
}
