package wargaming

import (
	"strings"
)

type WotEncyclopediaCrewskills struct {
	// Skill description
	Description string `json:"description,omitempty"`
	// Indicates if it is a perk
	IsPerk bool `json:"is_perk,omitempty"`
	// Skill name
	Name string `json:"name,omitempty"`
	// Skill ID
	Skill string `json:"skill,omitempty"`
	// URL to skill icon
	ImageUrl struct {
		// URL to large image
		BigIcon string `json:"big_icon,omitempty"`
		// URL to small image
		SmallIcon string `json:"small_icon,omitempty"`
	} `json:"image_url,omitempty"`
}

// WotEncyclopediaCrewskills Method returns full description of all crew skills.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/crewskills
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
// role:
//     Сrew qualification ID
// skill:
//     Skill ID. Maximum limit: 100.
func (client *Client) WotEncyclopediaCrewskills(realm Realm, fields []string, language string, role string, skill []string) (*WotEncyclopediaCrewskills, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"role": role,
		"skill": strings.Join(skill, ","),
	}

	var result *WotEncyclopediaCrewskills
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/crewskills/", reqParam, &result)
	return result, err
}
