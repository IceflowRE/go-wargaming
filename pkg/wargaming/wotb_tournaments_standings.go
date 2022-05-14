package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
	"strconv"
)

type WotbTournamentsStandings struct {
	// Number of battles played by a team
	BattlePlayed int `json:"battle_played,omitempty"`
	// Number of battles drawn by a team
	Draws int `json:"draws,omitempty"`
	// ID of a team's group
	GroupId int `json:"group_id,omitempty"`
	// Number of battles lost by a team
	Losses int `json:"losses,omitempty"`
	// Number of points earned by a team
	Points int `json:"points,omitempty"`
	// Team's place
	Position int `json:"position,omitempty"`
	// Stage ID
	StageId int `json:"stage_id,omitempty"`
	// Team ID
	TeamId int `json:"team_id,omitempty"`
	// Number of battles won by a team
	Wins int `json:"wins,omitempty"`
}

// WotbTournamentsStandings Method returns tournament results of each team.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/standings
//
// tournament_id:
//     Tournament ID that can be retrieved from the Tournaments list method.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// from_position:
//     Allows to get all team standings starting from a specific place, including this place
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Русский (by default)
// limit:
//     Number of returned entries. Default is 10. Min value is 1. Maximum value: 25.
// page_no:
//     Result page number. Default is 1. Min value is 1.
// team_id:
//     Team ID. Maximum limit: 10.
// to_position:
//     Allows to get all team standings up to a specific place, including this place
func (client *Client) WotbTournamentsStandings(realm Realm, tournamentId int, fields []string, fromPosition int, language string, limit int, pageNo int, teamId []int, toPosition int) (*WotbTournamentsStandings, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tournament_id": strconv.Itoa(tournamentId),
		"fields": strings.Join(fields, ","),
		"from_position": strconv.Itoa(fromPosition),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"team_id": utils.SliceIntToString(teamId, ","),
		"to_position": strconv.Itoa(toPosition),
	}

	var result *WotbTournamentsStandings
	err := client.doGetDataRequest(realm, "/wotb/tournaments/standings/", reqParam, &result)
	return result, err
}
