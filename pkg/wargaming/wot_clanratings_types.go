package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
)

// WotClanratingsTypes Method returns details on ratings types and categories.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/types
func (client *Client) WotClanratingsTypes(realm Realm) (*wot.ClanratingsTypes, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{

	}

	var result *wot.ClanratingsTypes
	err := client.doGetDataRequest(realm, "/wot/clanratings/types/", reqParam, &result)
	return result, err
}
