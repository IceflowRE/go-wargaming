package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strings"
)

// WotbEncyclopediaCrewskills Method returns information about crew skills.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/crewskills
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
// skillId:
//     Skills IDs. Maximum limit: 100.
// vehicleType:
//     Vehicle types. Maximum limit: 100.
func (client *Client) WotbEncyclopediaCrewskills(realm Realm, fields []string, language string, skillId []string, vehicleType []string) (*wotb.EncyclopediaCrewskills, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"skill_id": strings.Join(skillId, ","),
		"vehicle_type": strings.Join(vehicleType, ","),
	}

	var result *wotb.EncyclopediaCrewskills
	err := client.doGetDataRequest(realm, "/wotb/encyclopedia/crewskills/", reqParam, &result)
	return result, err
}
