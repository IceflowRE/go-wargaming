package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strconv"
	"strings"
)

// WotbTournamentsTeams Method returns list of teams in tournament.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/teams
//
// accountId:
//     ID of the account that belongs to the team. Maximum limit: 100.
// clanId:
//     ID of the clan that owns the team. Maximum limit: 100.
// tournamentId:
//     Tournament ID that can be retrieved from the Tournaments list method.
//     Parameter is required.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" - Русский (by default)
// limit:
//     Number of returned entries. Default is 10. Min value is 1. Maximum value: 100.
// pageNo:
//     Result page number. Default is 1. Min value is 1.
// search:
//     First letters in team name for search. Minimum length: 2 characters. Maximum length: 50 characters.
// status:
//     Team status. Maximum limit: 100. Valid values:
//     
//     "forming" - team roster is not yet confirmed 
//     "confirmed" - team roster is confirmed 
//     "disqualified" - team is disqualified
// teamId:
//     Team ID. Maximum limit: 25.
func (client *Client) WotbTournamentsTeams(realm Realm, accountId []int, clanId []int, tournamentId int, fields []string, language string, limit int, pageNo int, search string, status []string, teamId []int) (*wotb.TournamentsTeams, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
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

	var result *wotb.TournamentsTeams
	err := client.doGetDataRequest(realm, "/wotb/tournaments/teams/", reqParam, &result)
	return result, err
}