package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgn"
	"strconv"
	"strings"
)

// WgnAccountList Method returns partial list of players who have ever accessed in any Wargaming game. The list is filtered by name or by initial characters of user name and sorted alphabetically.
//
// https://developers.wargaming.net/reference/all/wgn/account/list
//
// search:
//     Player name search string. Parameter "type" defines minimum length and type of search. Using the exact search type, you can enter several names, separated with commas. Maximum length: 24.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// game:
//     Name of the game to player search. If the parameter is not specified, search will be executed across known games. Maximum limit: 10. Valid values:
//     
//     "wotb" - World of Tanks Blitz 
//     "wot" - World of Tanks 
//     "wows" - World of Warships 
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
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of None is applied (by default).
// type_:
//     Search type. Default is "startswith". Valid values:
//     
//     "startswith" - Search by initial characters of player name. Minimum length: 3 characters. Maximum length: 24 characters. (by default)
//     "exact" - Search by exact match of player name. Case insensitive. You can enter several names, separated with commas (up to 100).
func (client *Client) WgnAccountList(realm Realm, search string, fields []string, game []string, language string, limit int, type_ string) ([]*wgn.AccountList, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"search": search,
		"fields": strings.Join(fields, ","),
		"game": strings.Join(game, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"type": type_,
	}

	var result []*wgn.AccountList
	err := client.doGetDataRequest(realm, "/wgn/account/list/", reqParam, &result)
	return result, err
}
