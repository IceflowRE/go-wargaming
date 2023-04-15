// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotx"
	"strings"
)

// AccountInfo returns player details.
//
// https://developers.wargaming.net/reference/all/wotx/account/info
//
// realm:
//     Valid realms: RealmWgcb
// accountId:
//     Player account ID. Maximum limit: 100.
func (service *WotxService) AccountInfo(ctx context.Context, realm Realm, accountId []int, options *wotx.AccountInfoOptions) (*wotx.AccountInfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"account_id": internal.SliceIntToString(accountId, ","),
	}

	if options != nil {
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wotx.AccountInfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotx, realm, "/account/info/", reqParam, &data, &metaData)
	return data, metaData, err
}
