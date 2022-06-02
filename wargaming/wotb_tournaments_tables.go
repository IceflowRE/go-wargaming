package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/wargaming/wotb"
	"strconv"
	"strings"
)

// TournamentsTables Method returns tournament brackets.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/tables
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// stageId:
//     Stage ID that can be retrieved from the Tournaments Stages method.
// tournamentId:
//     Tournament ID that can be retrieved from the Tournaments list method.
func (service *wotbService) TournamentsTables(ctx context.Context, realm Realm, stageId int, tournamentId int, options *wotb.TournamentsTablesOptions) ([]*wotb.TournamentsTables, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"stage_id":      strconv.Itoa(stageId),
		"tournament_id": strconv.Itoa(tournamentId),
	}
	if options != nil {
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.GroupId != nil {
			reqParam["group_id"] = internal.SliceIntToString(options.GroupId, ",")
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data []*wotb.TournamentsTables
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tournaments/tables/", reqParam, &data)
	return data, err
}
