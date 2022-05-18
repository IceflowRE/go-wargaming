package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wowp"
	"strings"
)

// WowpEncyclopediaPlaneinfo Method returns aircraft details from Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planeinfo
//
// planeId:
//     Aircraft ID. Maximum limit: 1000.
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
func (client *Client) WowpEncyclopediaPlaneinfo(realm Realm, planeId []int, fields []string, language string) (*wowp.EncyclopediaPlaneinfo, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"plane_id": utils.SliceIntToString(planeId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wowp.EncyclopediaPlaneinfo
	err := client.doGetDataRequest(realm, "/wowp/encyclopedia/planeinfo/", reqParam, &result)
	return result, err
}
