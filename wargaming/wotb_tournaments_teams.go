package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotb"
	"strconv"
	"strings"
)

// TournamentsTeams Method returns list of teams in tournament.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/teams
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// tournamentId:
//     Tournament ID that can be retrieved from the Tournaments list method.
func (service *wotbService) TournamentsTeams(ctx context.Context, realm Realm, tournamentId int, options *wotb.TournamentsTeamsOptions) ([]*wotb.TournamentsTeams, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tournament_id": strconv.Itoa(tournamentId),
	}
	if options != nil {
		if options.TeamId != nil {
			reqParam["team_id"] = internal.SliceIntToString(options.TeamId, ",")
		}
		if options.Status != nil {
			reqParam["status"] = strings.Join(options.Status, ",")
		}
		if options.Search != nil {
			reqParam["search"] = *options.Search
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
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.ClanId != nil {
			reqParam["clan_id"] = internal.SliceIntToString(options.ClanId, ",")
		}
		if options.AccountId != nil {
			reqParam["account_id"] = internal.SliceIntToString(options.AccountId, ",")
		}
	}

	var data []*wotb.TournamentsTeams
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tournaments/teams/", reqParam, &data)
	return data, err
}
