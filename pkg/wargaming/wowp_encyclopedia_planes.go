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
//     "en" - English 
//     "ru" - Русский (by default)
//     "pl" - Polski 
//     "de" - Deutsch 
//     "fr" - Français 
//     "es" - Español 
//     "zh-cn" - 简体中文 
//     "tr" - Türkçe 
//     "cs" - Čeština 
//     "th" - ไทย 
//     "vi" - Tiếng Việt 
//     "ko" - 한국어
// nation:
//     Nation. Maximum limit: 100. Valid values:
//     
//     "ussr" - U.S.S.R. 
//     "germany" - Germany 
//     "usa" - U.S.A. 
//     "france" - France 
//     "uk" - U.K. 
//     "china" - China 
//     "japan" - Japan
// type_:
//     Type. Maximum limit: 100. Valid values:
//     
//     "fighter" - Fighter 
//     "fighterHeavy" - Heavy Fighter 
//     "assault" - Attack Aircraft 
//     "navy" - Multirole Fighter 
//     "bomber" - Bomber
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
