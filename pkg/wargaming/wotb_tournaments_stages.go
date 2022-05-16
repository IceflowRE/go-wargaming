package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strconv"
	"strings"
)

// WotbTournamentsStages Method returns list of stages in specified tournament.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/stages
//
// tournament_id:
//     Tournament ID that can be retrieved from the Tournaments list method.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Русский (by default)
// limit:
//     Number of returned entries. Default is 10. Min value is 1. Maximum value: 25.
// page_no:
//     Result page number. Default is 1. Min value is 1.
func (client *Client) WotbTournamentsStages(realm Realm, tournamentId int, fields []string, language string, limit int, pageNo int) (*wotb.TournamentsStages, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tournament_id": strconv.Itoa(tournamentId),
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
	}

	var result *wotb.TournamentsStages
	err := client.doGetDataRequest(realm, "/wotb/tournaments/stages/", reqParam, &result)
	return result, err
}
