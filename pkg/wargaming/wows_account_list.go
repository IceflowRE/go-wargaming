package wargaming

import (
	"strconv"
	"strings"
)

type WowsAccountList struct {
	// Player ID
	AccountId int `json:"account_id,omitempty"`
	// Player name
	Nickname string `json:"nickname,omitempty"`
}

// WowsAccountList Method returns partial list of players. The list is filtered by initial characters of user name and sorted alphabetically.
//
// https://developers.wargaming.net/reference/all/wows/account/list
//
// search:
//     Player name search string. Parameter "type" defines minimum length and type of search. Using the exact search type, you can enter several names, separated with commas. Maximum length: 24.
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
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// type:
//     Search type. Default is "startswith". Valid values:
//     
//     "startswith" &mdash; Search by initial characters of player name. Minimum length: 3 characters. Maximum length: 24 characters. (by default)
//     "exact" &mdash; Search by exact match of player name. Case insensitive. You can enter several names, separated with commas (up to 100).
func (client *Client) WowsAccountList(realm Realm, search string, fields []string, language string, limit int, type_ string) ([]*WowsAccountList, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"search": search,
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"type": type_,
	}

	var result []*WowsAccountList
	err := client.doGetDataRequest(realm, "/wows/account/list/", reqParam, &result)
	return result, err
}
