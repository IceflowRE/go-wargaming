package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strings"
)

// WotEncyclopediaTankinfo Method returns vehicle details from Tankopedia.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/tankinfo
//
// Deprecated: Attention! The method is deprecated.
//
// tankId:
//     Vehicle ID. Maximum limit: 1000.
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
//     "zh-tw" - 繁體中文 
//     "tr" - Türkçe 
//     "cs" - Čeština 
//     "th" - ไทย 
//     "vi" - Tiếng Việt 
//     "ko" - 한국어
func (client *Client) WotEncyclopediaTankinfo(realm Realm, tankId []int, fields []string, language string) (*wot.EncyclopediaTankinfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tank_id": utils.SliceIntToString(tankId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wot.EncyclopediaTankinfo
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/tankinfo/", reqParam, &result)
	return result, err
}
