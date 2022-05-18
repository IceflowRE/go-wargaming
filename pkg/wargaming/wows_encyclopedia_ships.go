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
//     "cs" &mdash; Čeština 
//     "de" &mdash; Deutsch 
//     "en" &mdash; English 
//     "es" &mdash; Español 
//     "fr" &mdash; Français 
//     "ja" &mdash; 日本語 
//     "pl" &mdash; Polski 
//     "ru" &mdash; Русский (by default)
//     "th" &mdash; ไทย 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "zh-cn" &mdash; 中文 
//     "pt-br" &mdash; Português do Brasil 
//     "es-mx" &mdash; Español (México)
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// nation:
//     Nation. Maximum limit: 100.
// page_no:
//     Result page number. Default is 1. Min value is 1.
// ship_id:
//     Ship ID. Maximum limit: 100.
// type:
//     Ship type. Maximum limit: 100. Valid values:
//     
//     "AirCarrier" &mdash; Aircraft carrier 
//     "Battleship" &mdash; Battleship 
//     "Destroyer" &mdash; Destroyer 
//     "Cruiser" &mdash; Cruiser
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
