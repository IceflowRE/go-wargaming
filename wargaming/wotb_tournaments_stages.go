// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wotb"
	"strconv"
	"strings"
)

// TournamentsStages returns list of stages in specified tournament.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/stages
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// tournamentId:
//     Tournament ID that can be retrieved from the Tournaments list method.
func (service *WotbService) TournamentsStages(ctx context.Context, realm Realm, tournamentId int, options *wotb.TournamentsStagesOptions) ([]*wotb.TournamentsStages, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tournament_id": strconv.Itoa(tournamentId),
	}

	if options != nil {
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
	}

	var data []*wotb.TournamentsStages
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tournaments/stages/", reqParam, &data, nil)
	return data, err
}
