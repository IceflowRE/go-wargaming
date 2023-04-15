// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotx"
)

// AccountXuidinfo returns player details based on the player Microsoft XUID.
//
// https://developers.wargaming.net/reference/all/wotx/account/xuidinfo
//
// realm:
//     Valid realms: RealmWgcb
// xuid:
//     Player Microsoft XUID. Maximum limit: 100.
func (service *WotxService) AccountXuidinfo(ctx context.Context, realm Realm, xuid []int) ([]*wotx.AccountXuidinfo, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"xuid": internal.SliceIntToString(xuid, ","),
	}

	var data []*wotx.AccountXuidinfo
	err := service.client.getRequest(ctx, sectionWotx, realm, "/account/xuidinfo/", reqParam, &data, nil)
	return data, err
}
