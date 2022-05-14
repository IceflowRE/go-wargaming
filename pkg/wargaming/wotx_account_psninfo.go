package wargaming

import (
	"strings"
)

type WotxAccountPsninfo struct {
	// Account ID
	AccountId int `json:"account_id,omitempty"`
	// Play Station UID
	Psnid string `json:"psnid,omitempty"`
}

// WotxAccountPsninfo Method returns player data based on Play Station UID
//
// https://developers.wargaming.net/reference/all/wotx/account/psninfo
//
// psnid:
//     Play Station UID. Maximum limit: 100.
func (client *Client) WotxAccountPsninfo(realm Realm, psnid []string) (*WotxAccountPsninfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"psnid": strings.Join(psnid, ","),
	}

	var result *WotxAccountPsninfo
	err := client.doGetDataRequest(realm, "/wotx/account/psninfo/", reqParam, &result)
	return result, err
}
