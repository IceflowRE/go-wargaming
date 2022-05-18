package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strings"
)

// WotbEncyclopediaModules Method returns list of available modules, such as guns, engines, etc.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/modules
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
//     "zh-tw" - 繁體中文 
//     "tr" - Türkçe 
//     "cs" - Čeština 
//     "th" - ไทย 
//     "vi" - Tiếng Việt 
//     "ko" - 한국어
// moduleId:
//     Module ID. Maximum limit: 100.
// nation:
//     Nation
// type_:
//     Module type. Valid values:
//     
//     "vehicleEngine" - Engines 
//     "vehicleGun" - Gun 
//     "vehicleChassis" - Suspension 
//     "vehicleTurret" - Turret
func (client *Client) WotbEncyclopediaModules(realm Realm, fields []string, language string, moduleId []int, nation string, type_ string) (*wotb.EncyclopediaModules, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"module_id": utils.SliceIntToString(moduleId, ","),
		"nation": nation,
		"type": type_,
	}

	var result *wotb.EncyclopediaModules
	err := client.doGetDataRequest(realm, "/wotb/encyclopedia/modules/", reqParam, &result)
	return result, err
}
