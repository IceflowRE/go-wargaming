package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wowp"
)

// WowpEncyclopediaInfo Method returns information about Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/info
func (client *Client) WowpEncyclopediaInfo(realm Realm) (*wowp.EncyclopediaInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{

	}

	var result *wowp.EncyclopediaInfo
	err := client.doGetDataRequest(realm, "/wowp/encyclopedia/info/", reqParam, &result)
	return result, err
}
