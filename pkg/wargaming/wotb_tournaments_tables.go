package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strconv"
	"strings"
)

// WotbTournamentsTables Method returns tournament brackets.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/tables
//
// stage_id:
//     Stage ID that can be retrieved from the Tournaments Stages method.
// tournament_id:
//     Tournament ID that can be retrieved from the Tournaments list method.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// group_id:
//     Group ID. Maximum limit: 10.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Русский (by default)
// limit:
//     Number of returned entries. Default is 10. Min value is 1. Maximum value: 25.
// page_no:
//     Result page number. Default is 1. Min value is 1.
func (client *Client) WotbTournamentsTables(realm Realm, stageId int, tournamentId int, fields []string, groupId []int, language string, limit int, pageNo int) (*wotb.TournamentsTables, error) {
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
	}

	var result *wotb.TournamentsTables
	err := client.doGetDataRequest(realm, "/wotb/tournaments/tables/", reqParam, &result)
	return result, err
}
