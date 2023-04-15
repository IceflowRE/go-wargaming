// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
	"strconv"
	"strings"
)

// ClansList searches through clans and sorts them in a specified order.
//
// https://developers.wargaming.net/reference/all/wowp/clans/list
//
// realm:
//     Valid realms: RealmEu, RealmNa
func (service *WowpService) ClansList(ctx context.Context, realm Realm, options *wowp.ClansListOptions) ([]*wowp.ClansList, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
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

	var data []*wowp.ClansList
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/clans/list/", reqParam, &data, &metaData)
	return data, metaData, err
}
