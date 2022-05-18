package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strings"
)

// WotStrongholdClanreserves Method returns information about available Reserves and their current status.
//
// https://developers.wargaming.net/reference/all/wot/stronghold/clanreserves
//
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" - Russian (by default)
//     "be" - Belarusian 
//     "uk" - Ukrainian 
//     "kk" - Kazakh
func (client *Client) WotStrongholdClanreserves(realm Realm, accessToken string, fields []string, language string) (*wot.StrongholdClanreserves, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wot.StrongholdClanreserves
	err := client.doGetDataRequest(realm, "/wot/stronghold/clanreserves/", reqParam, &result)
	return result, err
}
