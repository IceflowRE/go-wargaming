package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wowp"
	"strconv"
	"strings"
)

// WowpAccountList Method returns partial list of players. The list is filtered by initial characters of user name and sorted alphabetically.
//
// https://developers.wargaming.net/reference/all/wowp/account/list
//
// search:
//     Player name search string. Parameter "type" defines minimum length and type of search. Using the exact search type, you can enter several names, separated with commas. Maximum length: 24.
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
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of None is applied (by default).
// type_:
//     Search type. Default is "startswith". Valid values:
//     
//     "startswith" &mdash; Search by initial characters of player name. Minimum length: 3 characters. Maximum length: 24 characters. (by default)
//     "exact" &mdash; Search by exact match of player name. Case insensitive. You can enter several names, separated with commas (up to 100).
func (client *Client) WowpAccountList(realm Realm, search string, fields []string, language string, limit int, type_ string) ([]*wowp.AccountList, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"search": search,
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"type": type_,
	}

	var result []*wowp.AccountList
	err := client.doGetDataRequest(realm, "/wowp/account/list/", reqParam, &result)
	return result, err
}
