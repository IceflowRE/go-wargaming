package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotbEncyclopediaProvisions struct {
	// Localized description
	DescriptionI18N string `json:"description_i18n,omitempty"`
	// Localized name
	NameI18N string `json:"name_i18n,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Purchase cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Equipment or consumable ID
	ProvisionId int `json:"provision_id,omitempty"`
	// List of compatible vehicle IDs
	Tanks []int `json:"tanks,omitempty"`
	// Type
	Type_ string `json:"type,omitempty"`
}

// WotbEncyclopediaProvisions Method returns list of available equipment and consumables.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/provisions
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
// provision_id:
//     Equipment or consumable ID. Maximum limit: 100.
// tank_id:
//     Vehicle ID. Maximum limit: 100.
// type:
//     Type. Valid values:
//     
//     "equipment" &mdash; Consumables 
//     "optionalDevice" &mdash; Equipment
func (client *Client) WotbEncyclopediaProvisions(realm Realm, fields []string, language string, provisionId []int, tankId []int, type_ string) (*WotbEncyclopediaProvisions, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"provision_id": utils.SliceIntToString(provisionId, ","),
		"tank_id": utils.SliceIntToString(tankId, ","),
		"type": type_,
	}

	var result *WotbEncyclopediaProvisions
	err := client.doGetDataRequest(realm, "/wotb/encyclopedia/provisions/", reqParam, &result)
	return result, err
}
