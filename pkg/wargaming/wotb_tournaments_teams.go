package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
	"strconv"
)

type WotbTournamentsTeams struct {
	// Team clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Team status
	Status string `json:"status,omitempty"`
	// Team ID
	TeamId int `json:"team_id,omitempty"`
	// Team name
	Title string `json:"title,omitempty"`
	// Tournament ID
	TournamentId int `json:"tournament_id,omitempty"`
	// Information on team players
	Players struct {
		// Player account ID
		AccountId int `json:"account_id,omitempty"`
		// Link to player image
		Image string `json:"image,omitempty"`
		// Player name
		Name string `json:"name,omitempty"`
		// Technical position name
		Role string `json:"role,omitempty"`
	} `json:"players,omitempty"`
}

// WotbTournamentsTeams Method returns list of teams in tournament.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/teams
//
// account_id:
//     ID of the account that belongs to the team. Maximum limit: 100.
// clan_id:
//     ID of the clan that owns the team. Maximum limit: 100.
// tournament_id:
//     Tournament ID that can be retrieved from the Tournaments list method.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Русский (by default)
// limit:
//     Number of returned entries. Default is 10. Min value is 1. Maximum value: 100.
// page_no:
//     Result page number. Default is 1. Min value is 1.
// search:
//     First letters in team name for search. Minimum length: 2 characters. Maximum length: 50 characters.
// status:
//     Team status. Maximum limit: 100. Valid values:
//     
//     "forming" &mdash; team roster is not yet confirmed 
//     "confirmed" &mdash; team roster is confirmed 
//     "disqualified" &mdash; team is disqualified
// team_id:
//     Team ID. Maximum limit: 25.
func (client *Client) WotbTournamentsTeams(realm Realm, accountId []int, clanId []int, tournamentId int, fields []string, language string, limit int, pageNo int, search string, status []string, teamId []int) (*WotbTournamentsTeams, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"clan_id": utils.SliceIntToString(clanId, ","),
		"tournament_id": strconv.Itoa(tournamentId),
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"search": search,
		"status": strings.Join(status, ","),
		"team_id": utils.SliceIntToString(teamId, ","),
	}

	var result *WotbTournamentsTeams
	err := client.doGetDataRequest(realm, "/wotb/tournaments/teams/", reqParam, &result)
	return result, err
}
