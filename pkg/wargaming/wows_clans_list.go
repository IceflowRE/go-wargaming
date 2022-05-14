package wargaming

import (
	"strconv"
	"strings"
)

type WowsClansList struct {
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Clan creation date
	CreatedAt UnixTime `json:"created_at,omitempty"`
	// Number of clan members
	MembersCount int `json:"members_count,omitempty"`
	// Clan name
	Name string `json:"name,omitempty"`
	// Clan tag
	Tag string `json:"tag,omitempty"`
}

// WowsClansList Method searches through clans and sorts them in a specified order.
//
// https://developers.wargaming.net/reference/all/wows/clans/list
//
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
// page_no:
//     Page number. Default is 1. Min value is 1.
// search:
//     Part of name or tag for clan search. Minimum 2 characters
func (client *Client) WowsClansList(realm Realm, fields []string, language string, limit int, pageNo int, search string) ([]*WowsClansList, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"search": search,
	}

	var result []*WowsClansList
	err := client.doGetDataRequest(realm, "/wows/clans/list/", reqParam, &result)
	return result, err
}
