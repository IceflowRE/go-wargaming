// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strconv"
	"strings"
)

// ClansSeasonstats returns clan battles season player stats.
//
// https://developers.wargaming.net/reference/all/wows/clans/seasonstats
//
// realm:
//
//	Valid realms: RealmAsia, RealmEu, RealmNa
//
// accountId:
//
//	Player account ID. Min value is 1.
func (service *WowsService) ClansSeasonstats(ctx context.Context, realm Realm, accountId int, options *wows.ClansSeasonstatsOptions) (*wows.ClansSeasonstats, error) {
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
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wows.ClansSeasonstats
	err := service.client.getRequest(ctx, sectionWows, realm, "/clans/seasonstats/", reqParam, &data, nil)
	return data, err
}
