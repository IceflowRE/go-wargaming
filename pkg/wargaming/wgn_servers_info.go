package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgn"
	"strings"
)

// WgnServersInfo Method returns the number of online players on the servers.
//
// https://developers.wargaming.net/reference/all/wgn/servers/info
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// game:
//     Game ID. Maximum limit: 100. Valid values:
//     
//     "wotb" - World of Tanks Blitz 
//     "wot" - World of Tanks 
//     "wows" - World of Warships
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "en" - English 
//     "ru" - Русский (by default)
//     "pl" - Polski 
//     "de" - Deutsch 
//     "fr" - Français 
//     "es" - Español 
//     "zh-cn" - 简体中文 
//     "zh-tw" - 繁體中文 
//     "tr" - Türkçe 
//     "cs" - Čeština 
//     "th" - ไทย 
//     "vi" - Tiếng Việt 
//     "ko" - 한국어
func (client *Client) WgnServersInfo(realm Realm, fields []string, game []string, language string) (*wgn.ServersInfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"game": strings.Join(game, ","),
		"language": language,
	}

	var result *wgn.ServersInfo
	err := client.doGetDataRequest(realm, "/wgn/servers/info/", reqParam, &result)
	return result, err
}
