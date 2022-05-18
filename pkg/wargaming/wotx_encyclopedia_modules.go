package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotx"
	"strconv"
	"strings"
)

// WotxEncyclopediaModules Method returns list of available modules that can be installed on vehicles, such as engines, turrets, etc. At least one input filter parameter (module ID, type) is required to be indicated.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/modules
//
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "default_profile"
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "en". Valid values:
//     
//     "en" - English (by default)
//     "ru" - Русский 
//     "pl" - Polski 
//     "de" - Deutsch 
//     "fr" - Français 
//     "es" - Español 
//     "tr" - Türkçe
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// moduleId:
//     Module ID. Maximum limit: 100.
// nation:
//     Nation. Maximum limit: 100.
// pageNo:
//     Result page number
// type_:
//     Module type. Maximum limit: 100. Valid values:
//     
//     "vehicleRadio" - Radio 
//     "vehicleEngine" - Engines 
//     "vehicleGun" - Gun 
//     "vehicleChassis" - Suspension 
//     "vehicleTurret" - Turret
func (client *Client) WotxEncyclopediaModules(realm Realm, extra []string, fields []string, language string, limit int, moduleId []int, nation []string, pageNo int, type_ []string) (*wotx.EncyclopediaModules, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
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

	var result *wotx.EncyclopediaModules
	err := client.doGetDataRequest(realm, "/wotx/encyclopedia/modules/", reqParam, &result)
	return result, err
}
