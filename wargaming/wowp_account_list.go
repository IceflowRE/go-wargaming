// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wowp"
	"strconv"
	"strings"
)

// AccountList returns partial list of players. The list is filtered by initial characters of user name and sorted alphabetically.
//
// https://developers.wargaming.net/reference/all/wowp/account/list
//
// realm:
//     Valid realms: RealmEu, RealmNa
// search:
//     Player name search string. Parameter "type" defines minimum length and type of search. Using the exact search type, you can enter several names, separated with commas. Maximum length: 24.
func (service *WowpService) AccountList(ctx context.Context, realm Realm, search string, options *wowp.AccountListOptions) ([]*wowp.AccountList, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, err
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

	var data []*wowp.AccountList
	err := service.client.getRequest(ctx, sectionWowp, realm, "/account/list/", reqParam, &data, nil)
	return data, err
}
