package wargaming

import (
	"strconv"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotEncyclopediaProvisions struct {
	// Achievement description
	Description string `json:"description,omitempty"`
	// Image link
	Image string `json:"image,omitempty"`
	// Vehicle name
	Name string `json:"name,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Equipment or consumables ID
	ProvisionId int `json:"provision_id,omitempty"`
	// Technical name
	Tag string `json:"tag,omitempty"`
	// Type: consumable or equipment
	Type_ string `json:"type,omitempty"`
	// Weight in kilos. Used for equipment only.
	Weight int `json:"weight,omitempty"`
}

// WotEncyclopediaProvisions Method returns a list of available equipment and consumables.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/provisions
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
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// page_no:
//     Result page number
// provision_id:
//     Equipment or consumables ID. Maximum limit: 100.
// type:
//     Type. Default is "equipment, optionalDevice". Maximum limit: 100. Valid values:
//     
//     "equipment" &mdash; Consumables 
//     "optionalDevice" &mdash; Equipment
func (client *Client) WotEncyclopediaProvisions(realm Realm, fields []string, language string, limit int, pageNo int, provisionId []int, type_ []string) (*WotEncyclopediaProvisions, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"provision_id": utils.SliceIntToString(provisionId, ","),
		"type": strings.Join(type_, ","),
	}

	var result *WotEncyclopediaProvisions
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/provisions/", reqParam, &result)
	return result, err
}
