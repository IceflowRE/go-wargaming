package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wowp"
	"strconv"
	"strings"
)

// ClansList searches through clans and sorts them in a specified order.
//
// https://developers.wargaming.net/reference/all/wowp/clans/list
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
func (service *WowpService) ClansList(ctx context.Context, realm Realm, options *wowp.ClansListOptions) ([]*wowp.ClansList, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
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

	var data []*wowp.ClansList
	err := service.client.getRequest(ctx, sectionWowp, realm, "/clans/list/", reqParam, &data)
	return data, err
}
