package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotx"
	"strings"
)

// ClansAccountinfo returns player clan data. Player clan data exist only for accounts, that were clan members at least once.
//
// https://developers.wargaming.net/reference/all/wotx/clans/accountinfo
//
// realm:
//     Valid realms: RealmWgcb
// accountId:
//     Account ID. Maximum limit: 100. Min value is 1.
func (service *WotxService) ClansAccountinfo(ctx context.Context, realm Realm, accountId []int, options *wotx.ClansAccountinfoOptions) (*wotx.ClansAccountinfo, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
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
	}

	var data *wotx.ClansAccountinfo
	err := service.client.getRequest(ctx, sectionWotx, realm, "/clans/accountinfo/", reqParam, &data)
	return data, err
}
