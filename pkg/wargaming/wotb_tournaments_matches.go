package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strconv"
	"strings"
)

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
func (client *Client) WotbTournamentsMatches(realm Realm, stageId int, tournamentId int, fields []string, groupId []int, language string, limit int, pageNo int, roundNumber []int, teamId []int) (*wotb.TournamentsMatches, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
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

	var result *wotb.TournamentsMatches
	err := client.doGetDataRequest(realm, "/wotb/tournaments/matches/", reqParam, &result)
	return result, err
}
