// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
	"strings"
)

// ClansAccountinfo returns player clan data. Player clan data exist only for accounts, that were participating in clan activities: sent join requests, were clan members etc.
//
// https://developers.wargaming.net/reference/all/wowp/clans/accountinfo
//
// realm:
//     Valid realms: RealmEu, RealmNa
// accountId:
//     Account ID. Maximum limit: 100. Min value is 1.
func (service *WowpService) ClansAccountinfo(ctx context.Context, realm Realm, accountId []int, options *wowp.ClansAccountinfoOptions) (*wowp.ClansAccountinfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"account_id": internal.SliceIntToString(accountId, ","),
	}

	if options != nil {
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

	var data *wowp.ClansAccountinfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/clans/accountinfo/", reqParam, &data, &metaData)
	return data, metaData, err
}
