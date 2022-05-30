package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strings"
)

// WotClansMessageboard Method returns messages of clan message board.This method will be removed. Use method Message board (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wot/clans/messageboard
//
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
//     Parameter is required.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WotClansMessageboard(realm Realm, accessToken string, fields []string) (*wot.ClansMessageboard, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
	}

	var result *wot.ClansMessageboard
	err := client.doGetDataRequest(realm, "/wot/clans/messageboard/", reqParam, &result)
	return result, err
}
