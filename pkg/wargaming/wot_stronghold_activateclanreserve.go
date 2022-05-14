package wargaming

import (
	"strings"
	"strconv"
)

type WotStrongholdActivateclanreserve struct {
	// Activation time of a clan Reserve
	ActivatedAt UnixTime `json:"activated_at,omitempty"`
}

// WotStrongholdActivateclanreserve This method activates an available clan Reserve. A clan Reserve can be activated only by a clan member with the required permission.
//
// https://developers.wargaming.net/reference/all/wot/stronghold/activateclanreserve
//
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// reserve_level:
//     Level of clan Reserve to be activated
// reserve_type:
//     Type of clan Reserve to be activated
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Russian (by default)
//     "be" &mdash; Belarusian 
//     "uk" &mdash; Ukrainian 
//     "kk" &mdash; Kazakh
func (client *Client) WotStrongholdActivateclanreserve(realm Realm, accessToken string, reserveLevel int, reserveType string, fields []string, language string) (*WotStrongholdActivateclanreserve, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"reserve_level": strconv.Itoa(reserveLevel),
		"reserve_type": reserveType,
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotStrongholdActivateclanreserve
	err := client.doPostDataRequest(realm, "/wot/stronghold/activateclanreserve/", reqParam, &result)
	return result, err
}
