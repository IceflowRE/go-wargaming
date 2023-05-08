// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strings"
)

// AccountInfo returns player details. Players may hide their game profiles, use field hidden_profile for determination.
//
// https://developers.wargaming.net/reference/all/wows/account/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Player account ID. Maximum limit: 100. Min value is 1.
func (service *WowsService) AccountInfo(ctx context.Context, realm Realm, accountId []int, options *wows.AccountInfoOptions) (map[int]*wows.AccountInfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"account_id": internal.SliceIntToString(accountId, ","),
	}

	if options != nil {
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
		if options.Extra != nil {
			reqParam["extra"] = strings.Join(options.Extra, ",")
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data map[int]*wows.AccountInfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWows, realm, "/account/info/", reqParam, &data, &metaData)
	return data, metaData, err
}
