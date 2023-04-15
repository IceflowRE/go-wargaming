// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strconv"
	"strings"
)

// TournamentsList returns list of tournaments.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/list
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotbService) TournamentsList(ctx context.Context, realm Realm, options *wotb.TournamentsListOptions) ([]*wotb.TournamentsList, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

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
		if options.Search != nil {
			reqParam["search"] = *options.Search
		}
		if options.Status != nil {
			reqParam["status"] = strings.Join(options.Status, ",")
		}
	}

	var data []*wotb.TournamentsList
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tournaments/list/", reqParam, &data, &metaData)
	return data, metaData, err
}
