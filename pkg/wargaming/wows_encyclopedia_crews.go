package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WowsEncyclopediaCrews struct {
	// Basic training cost
	BaseTrainingHirePrice int `json:"base_training_hire_price,omitempty"`
	// Basic training level
	BaseTrainingLevel int `json:"base_training_level,omitempty"`
	// Commanders' first names
	FirstNames []string `json:"first_names,omitempty"`
	// Retraining cost in doubloons
	GoldRetrainingPrice int `json:"gold_retraining_price,omitempty"`
	// Training cost in gold
	GoldTrainingHirePrice int `json:"gold_training_hire_price,omitempty"`
	// Commander training level purchased for doubloons
	GoldTrainingLevel int `json:"gold_training_level,omitempty"`
	// The list of the Commander images:
	// 
	// 1—URL to the image of the Commander with 1–7 skill points;
	// 8—URL to the image of the Commander with 8–13 skill points;
	// 14—URL to the image of the Commander with 14–20 skill points;.
	// 
	// If only the value for the key 1 is specified, the Commander has not earned skill points yet.
	Icons map[string]string `json:"icons,omitempty"`
	// Indicates if the skill is preserved after retraining
	IsRetrainable bool `json:"is_retrainable,omitempty"`
	// Commanders' last names
	LastNames []string `json:"last_names,omitempty"`
	// Retraining cost in credits
	MoneyRetrainingPrice int `json:"money_retraining_price,omitempty"`
	// Training cost in credits
	MoneyTrainingHirePrice int `json:"money_training_hire_price,omitempty"`
	// Commander training level purchased for credits
	MoneyTrainingLevel int `json:"money_training_level,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Subnation index
	SubnationIndex int `json:"subnation_index,omitempty"`
}

// WowsEncyclopediaCrews Method returns information about Commanders.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/crews
//
// commander_id:
//     Commander ID. Maximum limit: 100.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "cs" &mdash; Čeština 
//     "de" &mdash; Deutsch 
//     "en" &mdash; English 
//     "es" &mdash; Español 
//     "fr" &mdash; Français 
//     "ja" &mdash; 日本語 
//     "pl" &mdash; Polski 
//     "ru" &mdash; Русский (by default)
//     "th" &mdash; ไทย 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "zh-cn" &mdash; 中文 
//     "pt-br" &mdash; Português do Brasil 
//     "es-mx" &mdash; Español (México)
func (client *Client) WowsEncyclopediaCrews(realm Realm, commanderId []int, fields []string, language string) (*WowsEncyclopediaCrews, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"commander_id": utils.SliceIntToString(commanderId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowsEncyclopediaCrews
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/crews/", reqParam, &result)
	return result, err
}
