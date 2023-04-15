// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotx"
	"strconv"
	"strings"
)

// AccountList returns partial list of players. The list is filtered by initial characters of user name and sorted alphabetically.
//
// https://developers.wargaming.net/reference/all/wotx/account/list
//
// realm:
//     Valid realms: RealmWgcb
// search:
//     Player name search string. Parameter "type" defines minimum length and type of search. Using the exact search type, you can enter several names, separated with commas. Maximum length: 24.
func (service *WotxService) AccountList(ctx context.Context, realm Realm, search string, options *wotx.AccountListOptions) ([]*wotx.AccountList, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"search": search,
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
		if options.Type != nil {
			reqParam["type"] = *options.Type
		}
	}

	var data []*wotx.AccountList
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotx, realm, "/account/list/", reqParam, &data, &metaData)
	return data, metaData, err
}
