package wargaming

import (
	"strconv"
	"strings"
)

type WotClansList struct {
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Clan color in HEX #RRGGBB
	Color string `json:"color,omitempty"`
	// Clan creation date
	CreatedAt UnixTime `json:"created_at,omitempty"`
	// Number of clan members
	MembersCount int `json:"members_count,omitempty"`
	// Clan name
	Name string `json:"name,omitempty"`
	// Clan tag
	Tag string `json:"tag,omitempty"`
	// Information on clan emblems in games and on clan portal
	Emblems struct {
		// List of links to 195x195 px emblems
		X195 map[string]string `json:"x195,omitempty"`
		// List of links to 24x24 px emblems
		X24 map[string]string `json:"x24,omitempty"`
		// List of links to 256x256 px emblems
		X256 map[string]string `json:"x256,omitempty"`
		// List of links to 32x32 px emblems
		X32 map[string]string `json:"x32,omitempty"`
		// List of links to 64x64 px emblems
		X64 map[string]string `json:"x64,omitempty"`
	} `json:"emblems,omitempty"`
}

// WotClansList Method searches through clans and sorts them in a specified order.
//
// https://developers.wargaming.net/reference/all/wot/clans/list
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
// page_no:
//     Page number. Default is 1. Min value is 1.
// search:
//     Part of name or tag for clan search. Minimum 2 characters
func (client *Client) WotClansList(realm Realm, fields []string, language string, limit int, pageNo int, search string) ([]*WotClansList, error) {
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

	var result []*WotClansList
	err := client.doGetDataRequest(realm, "/wot/clans/list/", reqParam, &result)
	return result, err
}
