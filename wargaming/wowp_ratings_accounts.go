package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wowp"
	"strconv"
	"strings"
)

// RatingsAccounts returns player ratings by specified IDs.
//
// https://developers.wargaming.net/reference/all/wowp/ratings/accounts
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
// accountId:
//     Player account ID. Maximum limit: 100.
// type_:
//     Rating period. For valid values, check the Types of ratings method.
func (service *WowpService) RatingsAccounts(ctx context.Context, realm Realm, accountId []int, type_ string, options *wowp.RatingsAccountsOptions) (*wowp.RatingsAccounts, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": internal.SliceIntToString(accountId, ","),
		"type":       type_,
	}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Date != nil {
			reqParam["date"] = strconv.FormatInt(options.Date.Unix(), 10)
		}
	}

	var data *wowp.RatingsAccounts
	err := service.client.getRequest(ctx, sectionWowp, realm, "/ratings/accounts/", reqParam, &data)
	return data, err
}
