package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wows"
	"strconv"
	"strings"
)

// WowsEncyclopediaConsumables Method returns information about consumables: camouflages, flags, and upgrades.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/consumables
//
// consumableId:
//     Consumable ID. Maximum limit: 100.
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
// pageNo:
//     Page limit. Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default). Default is 1.
// type_:
//     Consumable type. Valid values:
//     
//     "Camouflage" &mdash; Camouflages 
//     "Flags" &mdash; Flags 
//     "Permoflage" &mdash; Permanent camouflages 
//     "Modernization" &mdash; Upgrades 
//     "Skin" &mdash; Ship camouflages
func (client *Client) WowsEncyclopediaConsumables(realm Realm, consumableId []int, fields []string, language string, limit int, pageNo int, type_ string) (*wows.EncyclopediaConsumables, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"consumable_id": utils.SliceIntToString(consumableId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"type": type_,
	}

	var result *wows.EncyclopediaConsumables
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/consumables/", reqParam, &result)
	return result, err
}
