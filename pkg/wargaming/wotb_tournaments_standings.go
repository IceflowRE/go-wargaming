package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strconv"
	"strings"
)

// WotbTournamentsStandings Method returns tournament results of each team.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/standings
//
// tournamentId:
//     Tournament ID that can be retrieved from the Tournaments list method.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// fromPosition:
//     Allows to get all team standings starting from a specific place, including this place
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Русский (by default)
// limit:
//     Number of returned entries. Default is 10. Min value is 1. Maximum value: 25.
// pageNo:
//     Result page number. Default is 1. Min value is 1.
// teamId:
//     Team ID. Maximum limit: 10.
// toPosition:
//     Allows to get all team standings up to a specific place, including this place
func (client *Client) WotbTournamentsStandings(realm Realm, tournamentId int, fields []string, fromPosition int, language string, limit int, pageNo int, teamId []int, toPosition int) (*wotb.TournamentsStandings, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
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

	var result *wotb.TournamentsStandings
	err := client.doGetDataRequest(realm, "/wotb/tournaments/standings/", reqParam, &result)
	return result, err
}
