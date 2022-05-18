package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgn"
	"strings"
)

// WgnClansMessageboard Method returns messages of clan message board.This method will be removed. Use method Message board (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wgn/clans/messageboard
//
// Deprecated: Attention! The method is deprecated.
//
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// game:
//     Name of the game to perform the clan search in. If the parameter is not specified, search will be executed across World of Tanks. Default is "wot". Valid values:
//     
//     "wot" &mdash; World of Tanks (by default)
//     "wowp" &mdash; World of Warplanes
func (client *Client) WgnClansMessageboard(realm Realm, accessToken string, fields []string, game string) (*wgn.ClansMessageboard, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
		"game": game,
	}

	var result *wgn.ClansMessageboard
	err := client.doGetDataRequest(realm, "/wgn/clans/messageboard/", reqParam, &result)
	return result, err
}
