package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotEncyclopediaTankturrets struct {
	// Armor: sides
	ArmorBoard int `json:"armor_board,omitempty"`
	// Armor: rear
	ArmorFedd int `json:"armor_fedd,omitempty"`
	// Armor: front
	ArmorForehead int `json:"armor_forehead,omitempty"`
	// View range
	CircularVisionRadius int `json:"circular_vision_radius,omitempty"`
	// Tier
	Level int `json:"level,omitempty"`
	// Module ID
	ModuleId int `json:"module_id,omitempty"`
	// Vehicle name
	Name string `json:"name,omitempty"`
	// Localized name field
	NameI18N string `json:"name_i18n,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Localized nation field
	NationI18N string `json:"nation_i18n,omitempty"`
	// Purchase cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Purchase cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Traverse speed
	RotationSpeed int `json:"rotation_speed,omitempty"`
	// IDs of compatible vehicles
	Tanks []int `json:"tanks,omitempty"`
}

// WotEncyclopediaTankturrets Deprecated: Attention! The method is deprecated.
// Method returns list of turrets.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/tankturrets
//
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
// module_id:
//     Module ID. Maximum limit: 1000.
// nation:
//     Nation. Maximum limit: 100.
func (client *Client) WotEncyclopediaTankturrets(realm Realm, fields []string, language string, moduleId []int, nation []string) (*WotEncyclopediaTankturrets, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"module_id": utils.SliceIntToString(moduleId, ","),
		"nation": strings.Join(nation, ","),
	}

	var result *WotEncyclopediaTankturrets
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/tankturrets/", reqParam, &result)
	return result, err
}
