// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strconv"
	"strings"
)

// TournamentsMatches returns information about matches of a particular tournament and stage.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/matches
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// stageId:
//     Stage ID that can be retrieved from the Tournaments Stages method.
// tournamentId:
//     Tournament ID that can be retrieved from the Tournaments list method.
func (service *WotbService) TournamentsMatches(ctx context.Context, realm Realm, stageId int, tournamentId int, options *wotb.TournamentsMatchesOptions) (*wotb.TournamentsMatches, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"stage_id":      strconv.Itoa(stageId),
		"tournament_id": strconv.Itoa(tournamentId),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.GroupId != nil {
			reqParam["group_id"] = internal.SliceIntToString(options.GroupId, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.RoundNumber != nil {
			reqParam["round_number"] = internal.SliceIntToString(options.RoundNumber, ",")
		}
		if options.TeamId != nil {
			reqParam["team_id"] = internal.SliceIntToString(options.TeamId, ",")
		}
	}

	var data *wotb.TournamentsMatches
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tournaments/matches/", reqParam, &data, nil)
	return data, err
}
