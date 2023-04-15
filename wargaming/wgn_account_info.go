// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgn"
	"strings"
)

// AccountInfo returns Wargaming account details.
//
// https://developers.wargaming.net/reference/all/wgn/account/info
//
// realm:
//     Valid realms: RealmEu, RealmNa
// accountId:
//     Player ID. Maximum limit: 100.
func (service *WgnService) AccountInfo(ctx context.Context, realm Realm, accountId []int, options *wgn.AccountInfoOptions) (*wgn.AccountInfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
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

	var data *wgn.AccountInfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWgn, realm, "/account/info/", reqParam, &data, &metaData)
	return data, metaData, err
}
