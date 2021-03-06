package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotx"
	"strconv"
	"strings"
)

// ClansList returns partial list of clans in alphabetical order by clan name or tag.
//
// https://developers.wargaming.net/reference/all/wotx/clans/list
//
// realm:
//     Valid realms: RealmWgcb
func (service *WotxService) ClansList(ctx context.Context, realm Realm, options *wotx.ClansListOptions) ([]*wotx.ClansList, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
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

	var data []*wotx.ClansList
	err := service.client.getRequest(ctx, sectionWotx, realm, "/clans/list/", reqParam, &data)
	return data, err
}
