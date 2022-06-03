package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotb"
	"strconv"
	"strings"
)

// TournamentsStandings returns tournament results of each team.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/standings
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// tournamentId:
//     Tournament ID that can be retrieved from the Tournaments list method.
func (service *WotbService) TournamentsStandings(ctx context.Context, realm Realm, tournamentId int, options *wotb.TournamentsStandingsOptions) ([]*wotb.TournamentsStandings, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tournament_id": strconv.Itoa(tournamentId),
	}
	if options != nil {
		if options.ToPosition != nil {
			reqParam["to_position"] = strconv.Itoa(*options.ToPosition)
		}
		if options.TeamId != nil {
			reqParam["team_id"] = internal.SliceIntToString(options.TeamId, ",")
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.FromPosition != nil {
			reqParam["from_position"] = strconv.Itoa(*options.FromPosition)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data []*wotb.TournamentsStandings
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tournaments/standings/", reqParam, &data)
	return data, err
}
