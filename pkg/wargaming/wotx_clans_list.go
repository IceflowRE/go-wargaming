package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotx"
	"strconv"
	"strings"
)

// WotxClansList Method returns partial list of clans in alphabetical order by clan name or tag.
//
// https://developers.wargaming.net/reference/all/wotx/clans/list
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
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// page_no:
//     Page number. Default is 1. Min value is 1.
// search:
//     Part of name or tag for clan search. Minimum 2 characters
func (client *Client) WotxClansList(realm Realm, fields []string, language string, limit int, pageNo int, search string) ([]*wotx.ClansList, error) {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"search": search,
	}

	var result []*wotx.ClansList
	err := client.doGetDataRequest(realm, "/wotx/clans/list/", reqParam, &result)
	return result, err
}
