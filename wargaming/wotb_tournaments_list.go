package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotb"
	"strconv"
	"strings"
)

// TournamentsList Method returns list of tournaments.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/list
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wotbService) TournamentsList(ctx context.Context, realm Realm, options *wotb.TournamentsListOptions) ([]*wotb.TournamentsList, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
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
	}

	var data []*wotb.TournamentsList
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tournaments/list/", reqParam, &data)
	return data, err
}
