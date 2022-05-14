package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotEncyclopediaTankengines struct {
	// Chance of fire on impact
	FireStartingChance int `json:"fire_starting_chance,omitempty"`
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
	// Power
	Power int `json:"power,omitempty"`
	// Purchase cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Purchase cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// IDs of compatible vehicles
	Tanks []int `json:"tanks,omitempty"`
}

// WotEncyclopediaTankengines Deprecated: Attention! The method is deprecated.
// Method returns list of engines.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/tankengines
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
func (client *Client) WotEncyclopediaTankengines(realm Realm, fields []string, language string, moduleId []int, nation []string) (*WotEncyclopediaTankengines, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"module_id": utils.SliceIntToString(moduleId, ","),
		"nation": strings.Join(nation, ","),
	}

	var result *WotEncyclopediaTankengines
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/tankengines/", reqParam, &result)
	return result, err
}
