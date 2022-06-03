package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wowp"
	"strings"
)

// AccountInfo2 Method returns player details.
//
// https://developers.wargaming.net/reference/all/wowp/account/info2
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
// accountId:
//     Player account ID. Maximum limit: 100.
func (service *WowpService) AccountInfo2(ctx context.Context, realm Realm, accountId []int, options *wowp.AccountInfo2Options) (*wowp.AccountInfo2, error) {
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

	var data *wowp.AccountInfo2
	err := service.client.getRequest(ctx, sectionWowp, realm, "/account/info2/", reqParam, &data)
	return data, err
}
