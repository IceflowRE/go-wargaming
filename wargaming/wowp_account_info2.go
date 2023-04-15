// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
	"strings"
)

// AccountInfo2 returns player details.
//
// https://developers.wargaming.net/reference/all/wowp/account/info2
//
// realm:
//     Valid realms: RealmEu, RealmNa
// accountId:
//     Player account ID. Maximum limit: 100.
func (service *WowpService) AccountInfo2(ctx context.Context, realm Realm, accountId []int, options *wowp.AccountInfo2Options) (*wowp.AccountInfo2, *GenericMeta, error) {
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

	var data *wowp.AccountInfo2
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/account/info2/", reqParam, &data, &metaData)
	return data, metaData, err
}
