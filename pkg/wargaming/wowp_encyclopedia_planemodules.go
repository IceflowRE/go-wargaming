package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wowp"
	"strings"
)

// WowpEncyclopediaPlanemodules Method returns information from Encyclopedia about modules available for specified aircrafts.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planemodules
//
// planeId:
//     Aircraft ID. Maximum limit: 1000.
//     Parameter is required.
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
// type_:
//     Configuration. Default is "engine, bomb, rocket, turret, turretfront, turretupper, turretlower, gun, construction". Maximum limit: 100. Valid values:
//     
//     "engine" - engine 
//     "bomb" - bomb 
//     "rocket" - rocket 
//     "turret" - rear gun 
//     "turretfront" - Передняя турель 
//     "turretupper" - Верхняя турель 
//     "turretlower" - Нижняя турель 
//     "gun" - autocannon 
//     "construction" - airframe
func (client *Client) WowpEncyclopediaPlanemodules(realm Realm, planeId []int, fields []string, language string, type_ []string) (*wowp.EncyclopediaPlanemodules, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"plane_id": utils.SliceIntToString(planeId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
		"type": strings.Join(type_, ","),
	}

	var result *wowp.EncyclopediaPlanemodules
	err := client.doGetDataRequest(realm, "/wowp/encyclopedia/planemodules/", reqParam, &result)
	return result, err
}
