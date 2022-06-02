package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wotb"
	"strconv"
	"strings"
)

// ClansList Method searches through clans and sorts them in a specified order.
//
// https://developers.wargaming.net/reference/all/wotb/clans/list
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wotbService) ClansList(ctx context.Context, realm Realm, options *wotb.ClansListOptions) ([]*wotb.ClansList, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
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

	var data []*wotb.ClansList
	err := service.client.getRequest(ctx, sectionWotb, realm, "/clans/list/", reqParam, &data)
	return data, err
}
