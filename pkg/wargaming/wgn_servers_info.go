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
//     "wotb" &mdash; World of Tanks Blitz 
//     "wot" &mdash; World of Tanks 
//     "wows" &mdash; World of Warships
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "en" &mdash; English 
//     "ru" &mdash; Русский (by default)
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "zh-cn" &mdash; 简体中文 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
func (client *Client) WgnServersInfo(realm Realm, fields []string, game []string, language string) (*wgn.ServersInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
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
