package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wotx"
	"strconv"
	"strings"
)

// AccountList Method returns partial list of players. The list is filtered by initial characters of user name and sorted alphabetically.
//
// https://developers.wargaming.net/reference/all/wotx/account/list
//
// realm:
//     Valid realms: RealmWgcb
// search:
//     Player name search string. Parameter "type" defines minimum length and type of search. Using the exact search type, you can enter several names, separated with commas. Maximum length: 24.
func (service *wotxService) AccountList(ctx context.Context, realm Realm, search string, options *wotx.AccountListOptions) ([]*wotx.AccountList, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"search": search,
	}
	if options != nil {
		if options.Type_ != nil {
			reqParam["type"] = *options.Type_
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

	var data []*wotx.AccountList
	err := service.client.getRequest(ctx, sectionWotx, realm, "/account/list/", reqParam, &data)
	return data, err
}
