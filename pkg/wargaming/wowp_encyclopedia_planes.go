package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wowp"
	"strings"
)

// WowpEncyclopediaPlanes Method returns list of all aircrafts from Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planes
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
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
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
// nation:
//     Nation. Maximum limit: 100. Valid values:
//     
//     "ussr" &mdash; U.S.S.R. 
//     "germany" &mdash; Germany 
//     "usa" &mdash; U.S.A. 
//     "france" &mdash; France 
//     "uk" &mdash; U.K. 
//     "china" &mdash; China 
//     "japan" &mdash; Japan
// type:
//     Type. Maximum limit: 100. Valid values:
//     
//     "fighter" &mdash; Fighter 
//     "fighterHeavy" &mdash; Heavy Fighter 
//     "assault" &mdash; Attack Aircraft 
//     "navy" &mdash; Multirole Fighter 
//     "bomber" &mdash; Bomber
func (client *Client) WowpEncyclopediaPlanes(realm Realm, fields []string, language string, nation []string, type_ []string) ([]*wowp.EncyclopediaPlanes, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"nation": strings.Join(nation, ","),
		"type": strings.Join(type_, ","),
	}

	var result []*wowp.EncyclopediaPlanes
	err := client.doGetDataRequest(realm, "/wowp/encyclopedia/planes/", reqParam, &result)
	return result, err
}
