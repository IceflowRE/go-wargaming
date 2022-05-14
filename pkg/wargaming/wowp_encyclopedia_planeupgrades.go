package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WowpEncyclopediaPlaneupgrades struct {
	// Slot name
	SlotName string `json:"slot_name,omitempty"`
	// Localized slot_name field
	SlotNameI18N string `json:"slot_name_i18n,omitempty"`
	// Modules available
	Modules struct {
		// ID of unique bind of slot and module
		BindId int `json:"bind_id,omitempty"`
		// Incompatible modules
		Conflicts []string `json:"conflicts,omitempty"`
		// Number of installed modules
		Count int `json:"count,omitempty"`
		// List of IDs of incompatible modules
		IncompatibleModules []int `json:"incompatible_modules,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Module name
		Name string `json:"name,omitempty"`
		// Aircraft ID to be opened with module research
		NextPlaneId int `json:"next_plane_id,omitempty"`
		// Name of parent module in Tech Tree
		Parent string `json:"parent,omitempty"`
		// Parent module ID in Tech Tree
		ParentId int `json:"parent_id,omitempty"`
	} `json:"modules,omitempty"`
}

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
func (client *Client) WowpEncyclopediaPlaneupgrades(realm Realm, planeId []int, fields []string, language string) (*WowpEncyclopediaPlaneupgrades, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"plane_id": utils.SliceIntToString(planeId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowpEncyclopediaPlaneupgrades
	err := client.doGetDataRequest(realm, "/wowp/encyclopedia/planeupgrades/", reqParam, &result)
	return result, err
}
