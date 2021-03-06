package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgn"
	"strings"
)

// AccountInfo returns Wargaming account details.
//
// https://developers.wargaming.net/reference/all/wgn/account/info
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
// accountId:
//     Player ID. Maximum limit: 100.
func (service *WgnService) AccountInfo(ctx context.Context, realm Realm, accountId []int, options *wgn.AccountInfoOptions) (*wgn.AccountInfo, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": internal.SliceIntToString(accountId, ","),
	}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
	}

	var data *wgn.AccountInfo
	err := service.client.getRequest(ctx, sectionWgn, realm, "/account/info/", reqParam, &data)
	return data, err
}
