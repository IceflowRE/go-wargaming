// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
	"strconv"
	"strings"
)

// RatingsAccounts returns player ratings by specified IDs.
//
// https://developers.wargaming.net/reference/all/wowp/ratings/accounts
//
// realm:
//     Valid realms: RealmEu, RealmNa
// accountId:
//     Player account ID. Maximum limit: 100.
// typ:
//     Rating period. For valid values, check the Types of ratings method.
func (service *WowpService) RatingsAccounts(ctx context.Context, realm Realm, accountId []int, typ string, options *wowp.RatingsAccountsOptions) (*wowp.RatingsAccounts, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"account_id": internal.SliceIntToString(accountId, ","),
		"type":       typ,
	}

	if options != nil {
		if options.Date != nil {
			reqParam["date"] = strconv.FormatInt(options.Date.Unix(), 10)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wowp.RatingsAccounts
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/ratings/accounts/", reqParam, &data, &metaData)
	return data, metaData, err
}
