package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strconv"
	"strings"
)

// WotbTournamentsList Method returns list of tournaments.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/list
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Русский (by default)
// limit:
//     Page limit. Number of tournaments in one page. Default is 10. Min value is 1. Maximum value: 25.
// page_no:
//     Result page number. Default is 1. Min value is 1.
// search:
//     First letters in tournament name for search. Minimum length: 2 characters. Maximum length: 50 characters.
// status:
//     Tournament status. Maximum limit: 100. Valid values:
//     
//     "upcoming" &mdash; Tournament is planned 
//     "registration_started" &mdash; Players can apply to join the tournament 
//     "registration_finished" &mdash; Registration has finished 
//     "running" &mdash; The first match has started 
//     "finished" &mdash; The last match among all stages has been played 
//     "complete" &mdash; Tournament has been completed
func (client *Client) WotbTournamentsList(realm Realm, fields []string, language string, limit int, pageNo int, search string, status []string) ([]*wotb.TournamentsList, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"search": search,
		"status": strings.Join(status, ","),
	}

	var result []*wotb.TournamentsList
	err := client.doGetDataRequest(realm, "/wotb/tournaments/list/", reqParam, &result)
	return result, err
}
