package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strings"
)

// WotbTournamentsInfo Method returns detailed information on the specified tournament.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/info
//
// tournament_id:
//     Tournament ID that can be retrieved from the Tournaments list method. Maximum limit: 25.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Русский (by default)
func (client *Client) WotbTournamentsInfo(realm Realm, tournamentId []int, fields []string, language string) (*wotb.TournamentsInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tournament_id": utils.SliceIntToString(tournamentId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wotb.TournamentsInfo
	err := client.doGetDataRequest(realm, "/wotb/tournaments/info/", reqParam, &result)
	return result, err
}
