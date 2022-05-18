package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wows"
	"strings"
)

// WowsEncyclopediaCrews Method returns information about Commanders.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/crews
//
// commanderId:
//     Commander ID. Maximum limit: 100.
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
func (client *Client) WowsEncyclopediaCrews(realm Realm, commanderId []int, fields []string, language string) (*wows.EncyclopediaCrews, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"commander_id": utils.SliceIntToString(commanderId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wows.EncyclopediaCrews
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/crews/", reqParam, &result)
	return result, err
}
