package wargaming

import (
	"strings"
)

type WgnClansGlossary struct {
	// Available clan positions
	ClansRoles map[string]string `json:"clans_roles,omitempty"`
	// Game with clan entities
	Game string `json:"game,omitempty"`
}

// WgnClansGlossary Deprecated: Attention! The method is deprecated.
// Method returns information on clan entities in World of Tanks and World of Warplanes.This method will be removed. Use method Clan glossary (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wgn/clans/glossary
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// game:
//     Name of the game to perform the clan search in. If the parameter is not specified, search will be executed across World of Tanks. Default is "wot". Valid values:
//     
//     "wot" &mdash; World of Tanks (by default)
//     "wowp" &mdash; World of Warplanes
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
func (client *Client) WgnClansGlossary(realm Realm, fields []string, game string, language string) (*WgnClansGlossary, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"game": game,
		"language": language,
	}

	var result *WgnClansGlossary
	err := client.doGetDataRequest(realm, "/wgn/clans/glossary/", reqParam, &result)
	return result, err
}
