// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wows"
	"strconv"
	"strings"
)

// AccountStatsbydate returns statistics slices by dates in specified time span.
//
// https://developers.wargaming.net/reference/all/wows/account/statsbydate
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Player account ID
func (service *WowsService) AccountStatsbydate(ctx context.Context, realm Realm, accountId int, options *wows.AccountStatsbydateOptions) (*wows.AccountStatsbydate, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
	}

	if options != nil {
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
		if options.Dates != nil {
			reqParam["dates"] = strings.Join(options.Dates, ",")
		}
		if options.Extra != nil {
			reqParam["extra"] = strings.Join(options.Extra, ",")
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wows.AccountStatsbydate
	err := service.client.getRequest(ctx, sectionWows, realm, "/account/statsbydate/", reqParam, &data, nil)
	return data, err
}
