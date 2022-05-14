package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
)

type WotxAccountXuidinfo struct {
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
	// Player Microsoft XUID
	Xuid int `json:"xuid,omitempty"`
}

// WotxAccountXuidinfo Method returns player details based on the player Microsoft XUID.
//
// https://developers.wargaming.net/reference/all/wotx/account/xuidinfo
//
// xuid:
//     Player Microsoft XUID. Maximum limit: 100.
func (client *Client) WotxAccountXuidinfo(realm Realm, xuid []int) (*WotxAccountXuidinfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"xuid": utils.SliceIntToString(xuid, ","),
	}

	var result *WotxAccountXuidinfo
	err := client.doGetDataRequest(realm, "/wotx/account/xuidinfo/", reqParam, &result)
	return result, err
}
