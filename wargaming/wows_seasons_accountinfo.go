// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strings"
)

// SeasonsAccountinfo returns players' statistics in Ranked Battles seasons. Accounts with hidden game profiles are excluded from response. Hidden profiles are listed in the field meta.hidden.
//
// https://developers.wargaming.net/reference/all/wows/seasons/accountinfo
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Player account ID. Maximum limit: 100. Min value is 1.
func (service *WowsService) SeasonsAccountinfo(ctx context.Context, realm Realm, accountId []int, options *wows.SeasonsAccountinfoOptions) (*wows.SeasonsAccountinfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
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
		if options.SeasonId != nil {
			reqParam["season_id"] = internal.SliceIntToString(options.SeasonId, ",")
		}
	}

	var data *wows.SeasonsAccountinfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWows, realm, "/seasons/accountinfo/", reqParam, &data, &metaData)
	return data, metaData, err
}
