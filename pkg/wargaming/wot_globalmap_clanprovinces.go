package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strings"
)

// WotGlobalmapClanprovinces Method returns lists of clans provinces.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/clanprovinces
//
// clan_id:
//     List of clan IDs. To get a clan ID, use the Clans method. Maximum limit: 10.
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Russian (by default)
func (client *Client) WotGlobalmapClanprovinces(realm Realm, clanId []int, accessToken string, fields []string, language string) (*wot.GlobalmapClanprovinces, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": utils.SliceIntToString(clanId, ","),
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wot.GlobalmapClanprovinces
	err := client.doGetDataRequest(realm, "/wot/globalmap/clanprovinces/", reqParam, &result)
	return result, err
}
