package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wowp"
	"strconv"
	"strings"
)

// WowpEncyclopediaPlanespecification Method returns information from Encyclopedia about aircraft specifications depending on installed modules.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planespecification
//
// bind_id:
//     ID of unique bind of slot and module. Maximum limit: 100.
// module_id:
//     Module ID. Maximum limit: 100.
// plane_id:
//     Aircraft ID
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
func (client *Client) WowpEncyclopediaPlanespecification(realm Realm, bindId []int, moduleId []int, planeId int, fields []string, language string) (*wowp.EncyclopediaPlanespecification, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"bind_id": utils.SliceIntToString(bindId, ","),
		"module_id": utils.SliceIntToString(moduleId, ","),
		"plane_id": strconv.Itoa(planeId),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wowp.EncyclopediaPlanespecification
	err := client.doGetDataRequest(realm, "/wowp/encyclopedia/planespecification/", reqParam, &result)
	return result, err
}
