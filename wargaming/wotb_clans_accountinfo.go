package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/wargaming/wotb"
	"strings"
)

// ClansAccountinfo Method returns player clan data. Player clan data exist only for accounts, that were participating in clan activities: sent join requests, were clan members etc.
//
// https://developers.wargaming.net/reference/all/wotb/clans/accountinfo
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accountId:
//     Account ID. Maximum limit: 100. Min value is 1.
func (service *wotbService) ClansAccountinfo(ctx context.Context, realm Realm, accountId []int, options *wotb.ClansAccountinfoOptions) (*wotb.ClansAccountinfo, error) {
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
	}

	var data *wotb.ClansAccountinfo
	err := service.client.getRequest(ctx, sectionWotb, realm, "/clans/accountinfo/", reqParam, &data)
	return data, err
}