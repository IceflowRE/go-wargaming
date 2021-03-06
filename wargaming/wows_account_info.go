package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strings"
)

// AccountInfo returns player details. Players may hide their game profiles, use field hidden_profile for determination.
//
// https://developers.wargaming.net/reference/all/wows/account/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accountId:
//     Player account ID. Maximum limit: 100. Min value is 1.
func (service *WowsService) AccountInfo(ctx context.Context, realm Realm, accountId []int, options *wows.AccountInfoOptions) (*wows.AccountInfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
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
		if options.Extra != nil {
			reqParam["extra"] = strings.Join(options.Extra, ",")
		}
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
	}

	var data *wows.AccountInfo
	err := service.client.getRequest(ctx, sectionWows, realm, "/account/info/", reqParam, &data)
	return data, err
}
