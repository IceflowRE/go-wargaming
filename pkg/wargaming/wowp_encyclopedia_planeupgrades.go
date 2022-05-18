package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wowp"
	"strings"
)

// WowpEncyclopediaPlaneupgrades Method returns information from Encyclopedia about slots of aircrafts and lists of modules which are compatible with specified slots.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planeupgrades
//
// plane_id:
//     Aircraft ID. Maximum limit: 100.
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
func (client *Client) WowpEncyclopediaPlaneupgrades(realm Realm, planeId []int, fields []string, language string) (*wowp.EncyclopediaPlaneupgrades, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"plane_id": utils.SliceIntToString(planeId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wowp.EncyclopediaPlaneupgrades
	err := client.doGetDataRequest(realm, "/wowp/encyclopedia/planeupgrades/", reqParam, &result)
	return result, err
}
