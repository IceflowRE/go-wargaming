package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wows"
	"strconv"
	"strings"
)

// WowsEncyclopediaModules Method returns list of available modules that can be mounted on a ship (hull, engines, etc.).
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/modules
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
// module_id:
//     Module ID. Maximum limit: 100.
// page_no:
//     Result page number. Default is 1. Min value is 1.
// type:
//     Module type. Valid values:
//     
//     "Artillery" &mdash; Main battery 
//     "Torpedoes" &mdash; Torpedo tubes 
//     "Suo" &mdash; Gun Fire Control System 
//     "FlightControl" &mdash; Flight Control 
//     "Hull" &mdash; Hull 
//     "Engine" &mdash; Engine 
//     "Fighter" &mdash; Fighters 
//     "TorpedoBomber" &mdash; Torpedo Bombers 
//     "DiveBomber" &mdash; Dive bombers
func (client *Client) WowsEncyclopediaModules(realm Realm, fields []string, language string, limit int, moduleId []int, pageNo int, type_ string) (*wows.EncyclopediaModules, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"module_id": utils.SliceIntToString(moduleId, ","),
		"page_no": strconv.Itoa(pageNo),
		"type": type_,
	}

	var result *wows.EncyclopediaModules
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/modules/", reqParam, &result)
	return result, err
}
