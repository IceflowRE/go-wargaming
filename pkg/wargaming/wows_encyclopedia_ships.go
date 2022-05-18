package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wows"
	"strconv"
	"strings"
)

// WowsEncyclopediaShips Method returns list of ships available.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/ships
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "cs" - Čeština 
//     "de" - Deutsch 
//     "en" - English 
//     "es" - Español 
//     "fr" - Français 
//     "ja" - 日本語 
//     "pl" - Polski 
//     "ru" - Русский (by default)
//     "th" - ไทย 
//     "zh-tw" - 繁體中文 
//     "tr" - Türkçe 
//     "zh-cn" - 中文 
//     "pt-br" - Português do Brasil 
//     "es-mx" - Español (México)
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// nation:
//     Nation. Maximum limit: 100.
// pageNo:
//     Result page number. Default is 1. Min value is 1.
// shipId:
//     Ship ID. Maximum limit: 100.
// type_:
//     Ship type. Maximum limit: 100. Valid values:
//     
//     "AirCarrier" - Aircraft carrier 
//     "Battleship" - Battleship 
//     "Destroyer" - Destroyer 
//     "Cruiser" - Cruiser
func (client *Client) WowsEncyclopediaShips(realm Realm, fields []string, language string, limit int, nation []string, pageNo int, shipId []int, type_ []string) (*wows.EncyclopediaShips, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"nation": strings.Join(nation, ","),
		"page_no": strconv.Itoa(pageNo),
		"ship_id": utils.SliceIntToString(shipId, ","),
		"type": strings.Join(type_, ","),
	}

	var result *wows.EncyclopediaShips
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/ships/", reqParam, &result)
	return result, err
}
