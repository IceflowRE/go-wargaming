// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strconv"
	"strings"
)

// ClansList searches through clans and sorts them in a specified order.
//
// https://developers.wargaming.net/reference/all/wotb/clans/list
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotbService) ClansList(ctx context.Context, realm Realm, options *wotb.ClansListOptions) ([]*wotb.ClansList, *GenericMeta, error) {
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
	}

	var data []*wotb.ClansList
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotb, realm, "/clans/list/", reqParam, &data, &metaData)
	return data, metaData, err
}
