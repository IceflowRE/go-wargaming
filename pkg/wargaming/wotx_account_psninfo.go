package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotx"
	"strings"
)

// WotxAccountPsninfo Method returns player data based on Play Station UID
//
// https://developers.wargaming.net/reference/all/wotx/account/psninfo
//
// psnid:
//     Play Station UID. Maximum limit: 100.
//     Parameter is required.
func (client *Client) WotxAccountPsninfo(realm Realm, psnid []string) (*wotx.AccountPsninfo, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"psnid": strings.Join(psnid, ","),
	}

	var result *wotx.AccountPsninfo
	err := client.doGetDataRequest(realm, "/wotx/account/psninfo/", reqParam, &result)
	return result, err
}
