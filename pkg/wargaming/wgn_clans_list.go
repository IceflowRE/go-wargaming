package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgn"
	"strconv"
	"strings"
)

// WgnClansList Method searches and sorts clans by the following logic:
// 
// exact match of clan tag placed first
// exact match of clan name placed second
// name or tag matches are placed next
// all comparisons are performed case insensitive
// expression [wg] is converted to wg and considered as the clan tag
// search for expression "[wg] clan" is performed by exact match of clan name and tag
// 
// Disbanded, NPC and technically frozen clans are excluded from response. Search is executed across World of Tanks and World of Warplanes.This method will be removed. Use method Clans (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wgn/clans/list
//
// Deprecated: Attention! The method is deprecated.
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// game:
//     Name of the game to perform the clan search in. If the parameter is not specified, search will be executed across World of Tanks. Default is "wot". Maximum limit: 100. Valid values:
//     
//     "wot" - World of Tanks 
//     "wowp" - World of Warplanes
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "en" - English 
//     "ru" - Русский (by default)
//     "pl" - Polski 
//     "de" - Deutsch 
//     "fr" - Français 
//     "es" - Español 
//     "zh-cn" - 简体中文 
//     "zh-tw" - 繁體中文 
//     "tr" - Türkçe 
//     "cs" - Čeština 
//     "th" - ไทย 
//     "vi" - Tiếng Việt 
//     "ko" - 한국어
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// pageNo:
//     Page number. Default is 1. Min value is 1.
// search:
//     Part of name or tag for clan search. Minimum 2 characters
func (client *Client) WgnClansList(realm Realm, fields []string, game []string, language string, limit int, pageNo int, search string) ([]*wgn.ClansList, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"game": strings.Join(game, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"search": search,
	}

	var result []*wgn.ClansList
	err := client.doGetDataRequest(realm, "/wgn/clans/list/", reqParam, &result)
	return result, err
}
