package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WowsEncyclopediaCrewskills struct {
	// URL to skill icon
	Icon string `json:"icon,omitempty"`
	// Name
	Name string `json:"name,omitempty"`
	// Tier
	Tier int `json:"tier,omitempty"`
	// Skill type ID
	TypeId int `json:"type_id,omitempty"`
	// Skill type name
	TypeName string `json:"type_name,omitempty"`
	// Skills
	Perks struct {
		// Description
		Description string `json:"description,omitempty"`
		// Skill ID
		PerkId int `json:"perk_id,omitempty"`
	} `json:"perks,omitempty"`
}

// WowsEncyclopediaCrewskills Method returns information about Commanders' skills.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/crewskills
//
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
// skill_id:
//     Skill ID. Maximum limit: 100.
func (client *Client) WowsEncyclopediaCrewskills(realm Realm, fields []string, language string, skillId []int) (*WowsEncyclopediaCrewskills, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"skill_id": utils.SliceIntToString(skillId, ","),
	}

	var result *WowsEncyclopediaCrewskills
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/crewskills/", reqParam, &result)
	return result, err
}
