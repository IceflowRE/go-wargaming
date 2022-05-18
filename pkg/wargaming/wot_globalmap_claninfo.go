package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strings"
)

// WotGlobalmapClaninfo Method returns clan data on the Global Map.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/claninfo
//
// clan_id:
//     Clan ID. To get a clan ID, use the Clans method. Maximum limit: 10.
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WotGlobalmapClaninfo(realm Realm, clanId []int, accessToken string, fields []string) (*wot.GlobalmapClaninfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": utils.SliceIntToString(clanId, ","),
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
	}

	var result *wot.GlobalmapClaninfo
	err := client.doGetDataRequest(realm, "/wot/globalmap/claninfo/", reqParam, &result)
	return result, err
}
