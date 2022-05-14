package wargaming

import (
	"strings"
)

type WotxEncyclopediaCrewroles struct {
	// Crew member qualification
	Name string `json:"name,omitempty"`
	// Crew skills and perks
	Skills struct {
		// Description of skill or perk
		Description string `json:"description,omitempty"`
		// Indicates whether a skill or a perk is dealt with
		IsPerk bool `json:"is_perk,omitempty"`
		// Name of skill or perk
		Name string `json:"name,omitempty"`
		// Images of skills and perks
		Images struct {
			// Link to large image
			Large string `json:"large,omitempty"`
		} `json:"images,omitempty"`
	} `json:"skills,omitempty"`
}

// WotxEncyclopediaCrewroles Method returns information about crews.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/crewroles
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "en". Valid values:
//     
//     "en" &mdash; English (by default)
//     "ru" &mdash; Русский 
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "tr" &mdash; Türkçe
func (client *Client) WotxEncyclopediaCrewroles(realm Realm, fields []string, language string) (*WotxEncyclopediaCrewroles, error) {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotxEncyclopediaCrewroles
	err := client.doGetDataRequest(realm, "/wotx/encyclopedia/crewroles/", reqParam, &result)
	return result, err
}
