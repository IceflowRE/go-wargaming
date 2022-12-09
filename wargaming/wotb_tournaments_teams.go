// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wotb"
	"strconv"
	"strings"
)

// TournamentsTeams returns list of teams in tournament.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/teams
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// tournamentId:
//     Tournament ID that can be retrieved from the Tournaments list method.
func (service *WotbService) TournamentsTeams(ctx context.Context, realm Realm, tournamentId int, options *wotb.TournamentsTeamsOptions) ([]*wotb.TournamentsTeams, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tournament_id": strconv.Itoa(tournamentId),
	}

	if options != nil {
		if options.AccountId != nil {
			reqParam["account_id"] = internal.SliceIntToString(options.AccountId, ",")
		}
		if options.ClanId != nil {
			reqParam["clan_id"] = internal.SliceIntToString(options.ClanId, ",")
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
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
		if options.Search != nil {
			reqParam["search"] = *options.Search
		}
		if options.Status != nil {
			reqParam["status"] = strings.Join(options.Status, ",")
		}
		if options.TeamId != nil {
			reqParam["team_id"] = internal.SliceIntToString(options.TeamId, ",")
		}
	}

	var data []*wotb.TournamentsTeams
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tournaments/teams/", reqParam, &data, nil)
	return data, err
}
