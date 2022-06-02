package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wotb"
	"strconv"
	"strings"
)

// TournamentsStages Method returns list of stages in specified tournament.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/stages
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// tournamentId:
//     Tournament ID that can be retrieved from the Tournaments list method.
func (service *wotbService) TournamentsStages(ctx context.Context, realm Realm, tournamentId int, options *wotb.TournamentsStagesOptions) ([]*wotb.TournamentsStages, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
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
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data []*wotb.TournamentsStages
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tournaments/stages/", reqParam, &data)
	return data, err
}
