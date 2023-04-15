// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgn"
	"strconv"
	"strings"
)

// AccountList returns partial list of players who have ever accessed in any Wargaming game. The list is filtered by name or by initial characters of user name and sorted alphabetically.
//
// https://developers.wargaming.net/reference/all/wgn/account/list
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// search:
//     Player name search string. Parameter "type" defines minimum length and type of search. Using the exact search type, you can enter several names, separated with commas. Maximum length: 24.
func (service *WgnService) AccountList(ctx context.Context, realm Realm, search string, options *wgn.AccountListOptions) ([]*wgn.AccountList, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"search": search,
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Game != nil {
			reqParam["game"] = strings.Join(options.Game, ",")
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

	var data []*wgn.AccountList
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWgn, realm, "/account/list/", reqParam, &data, &metaData)
	return data, metaData, err
}
