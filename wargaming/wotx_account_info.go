// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wotx"
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
func (service *WotxService) AccountInfo(ctx context.Context, realm Realm, accountId []int, options *wotx.AccountInfoOptions) (*wotx.AccountInfo, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
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
	err := service.client.getRequest(ctx, sectionWotx, realm, "/account/info/", reqParam, &data, nil)
	return data, err
}
