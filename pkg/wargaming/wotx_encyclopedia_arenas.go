package wargaming

import (
	"strings"
)

type WotxEncyclopediaArenas struct {
	// Map ID
	ArenaId string `json:"arena_id,omitempty"`
	// Map type
	CamouflageType string `json:"camouflage_type,omitempty"`
	// Short map description
	Description string `json:"description,omitempty"`
	// Localized map name
	NameI18N string `json:"name_i18n,omitempty"`
}

// WotxEncyclopediaArenas Method returns information about maps.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/arenas
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
func (client *Client) WotxEncyclopediaArenas(realm Realm, fields []string, language string) (*WotxEncyclopediaArenas, error) {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotxEncyclopediaArenas
	err := client.doGetDataRequest(realm, "/wotx/encyclopedia/arenas/", reqParam, &result)
	return result, err
}
