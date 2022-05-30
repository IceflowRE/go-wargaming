package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotx"
)

// WotxAccountXuidinfo Method returns player details based on the player Microsoft XUID.
//
// https://developers.wargaming.net/reference/all/wotx/account/xuidinfo
//
// xuid:
//     Player Microsoft XUID. Maximum limit: 100.
//     Parameter is required.
func (client *Client) WotxAccountXuidinfo(realm Realm, xuid []int) (*wotx.AccountXuidinfo, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"xuid": utils.SliceIntToString(xuid, ","),
	}

	var result *wotx.AccountXuidinfo
	err := client.doGetDataRequest(realm, "/wotx/account/xuidinfo/", reqParam, &result)
	return result, err
}
