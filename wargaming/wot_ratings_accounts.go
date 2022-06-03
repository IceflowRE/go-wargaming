package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strconv"
	"strings"
)

// RatingsAccounts Method returns player ratings by specified IDs.
//
// https://developers.wargaming.net/reference/all/wot/ratings/accounts
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accountId:
//     IDs of player accounts. Maximum limit: 100.
// type_:
//     Rating period. For valid values, check the Types of ratings method.
func (service *WotService) RatingsAccounts(ctx context.Context, realm Realm, accountId []int, type_ string, options *wot.RatingsAccountsOptions) (*wot.RatingsAccounts, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
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
		if options.BattleType != nil {
			reqParam["battle_type"] = *options.BattleType
		}
	}

	var data *wot.RatingsAccounts
	err := service.client.getRequest(ctx, sectionWot, realm, "/ratings/accounts/", reqParam, &data)
	return data, err
}
