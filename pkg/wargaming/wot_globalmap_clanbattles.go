package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapClanbattles Method returns list of clan's battles on the Global Map.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/clanbattles
//
// clanId:
//     Clan ID. To get a clan ID, use the Clans method.
//     Parameter is required.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Language. Default is "ru". Valid values:
//     
//     "ru" - Russian (by default)
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// pageNo:
//     Page number. Default is 1. Min value is 1.
func (client *Client) WotGlobalmapClanbattles(realm Realm, clanId int, fields []string, language string, limit int, pageNo int) (*wot.GlobalmapClanbattles, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": strconv.Itoa(clanId),
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
	}

	var result *wot.GlobalmapClanbattles
	err := client.doGetDataRequest(realm, "/wot/globalmap/clanbattles/", reqParam, &result)
	return result, err
}
