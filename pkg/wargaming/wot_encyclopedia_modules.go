package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotEncyclopediaModules Method returns list of available modules that can be installed on vehicles, such as engines, turrets, etc. At least one input filter parameter (module ID, type) is required to be indicated.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/modules
//
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "default_profile"
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
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// module_id:
//     Module ID. Maximum limit: 100.
// nation:
//     Nation. Maximum limit: 100.
// page_no:
//     Result page number
// type:
//     Module type. Maximum limit: 100. Valid values:
//     
//     "vehicleRadio" &mdash; Radio 
//     "vehicleEngine" &mdash; Engines 
//     "vehicleGun" &mdash; Gun 
//     "vehicleChassis" &mdash; Suspension 
//     "vehicleTurret" &mdash; Turret
func (client *Client) WotEncyclopediaModules(realm Realm, extra []string, fields []string, language string, limit int, moduleId []int, nation []string, pageNo int, type_ []string) (*wot.EncyclopediaModules, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"module_id": utils.SliceIntToString(moduleId, ","),
		"nation": strings.Join(nation, ","),
		"page_no": strconv.Itoa(pageNo),
		"type": strings.Join(type_, ","),
	}

	var result *wot.EncyclopediaModules
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/modules/", reqParam, &result)
	return result, err
}
