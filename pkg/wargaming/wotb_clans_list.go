package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strconv"
	"strings"
)

// WotbClansList Method searches through clans and sorts them in a specified order.
//
// https://developers.wargaming.net/reference/all/wotb/clans/list
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
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// pageNo:
//     Page number. Default is 1. Min value is 1.
// search:
//     Part of name or tag for clan search. Minimum 2 characters
func (client *Client) WotbClansList(realm Realm, fields []string, language string, limit int, pageNo int, search string) ([]*wotb.ClansList, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"search": search,
	}

	var result []*wotb.ClansList
	err := client.doGetDataRequest(realm, "/wotb/clans/list/", reqParam, &result)
	return result, err
}
