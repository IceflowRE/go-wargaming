package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wows"
	"strings"
)

// WowsEncyclopediaCrewranks Method returns information about Commanders' ranks.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/crewranks
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
// nation:
//     Nation
func (client *Client) WowsEncyclopediaCrewranks(realm Realm, fields []string, language string, nation string) (*wows.EncyclopediaCrewranks, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"nation": nation,
	}

	var result *wows.EncyclopediaCrewranks
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/crewranks/", reqParam, &result)
	return result, err
}
