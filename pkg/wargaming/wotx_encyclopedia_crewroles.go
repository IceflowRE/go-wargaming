package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotx"
	"strings"
)

// WotxEncyclopediaCrewroles Method returns information about crews.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/crewroles
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "en". Valid values:
//     
//     "en" - English (by default)
//     "ru" - Русский 
//     "pl" - Polski 
//     "de" - Deutsch 
//     "fr" - Français 
//     "es" - Español 
//     "tr" - Türkçe
func (client *Client) WotxEncyclopediaCrewroles(realm Realm, fields []string, language string) (*wotx.EncyclopediaCrewroles, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wotx.EncyclopediaCrewroles
	err := client.doGetDataRequest(realm, "/wotx/encyclopedia/crewroles/", reqParam, &result)
	return result, err
}
