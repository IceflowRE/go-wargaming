package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strings"
)

// WotStrongholdClaninfo Method returns general information and the battle statistics of clans in the Stronghold mode. Please note that information about the number of battles fought as well as the number of defeats and victories is updated once every 24 hours.
//
// https://developers.wargaming.net/reference/all/wot/stronghold/claninfo
//
// clanId:
//     Clan ID. To get a clan ID, use the Clans method. Maximum limit: 10.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Russian (by default)
//     "be" &mdash; Belarusian 
//     "uk" &mdash; Ukrainian 
//     "kk" &mdash; Kazakh
func (client *Client) WotStrongholdClaninfo(realm Realm, clanId []int, fields []string, language string) (*wot.StrongholdClaninfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": utils.SliceIntToString(clanId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wot.StrongholdClaninfo
	err := client.doGetDataRequest(realm, "/wot/stronghold/claninfo/", reqParam, &result)
	return result, err
}
