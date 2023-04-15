// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotx"
	"strconv"
	"strings"
)

// ClansList returns partial list of clans in alphabetical order by clan name or tag.
//
// https://developers.wargaming.net/reference/all/wotx/clans/list
//
// realm:
//     Valid realms: RealmWgcb
func (service *WotxService) ClansList(ctx context.Context, realm Realm, options *wotx.ClansListOptions) ([]*wotx.ClansList, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
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

	var data []*wotx.ClansList
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotx, realm, "/clans/list/", reqParam, &data, &metaData)
	return data, metaData, err
}
