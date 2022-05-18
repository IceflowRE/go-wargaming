package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotClanratingsTop Method returns the list of top clans by specified parameters.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/top
//
// limit:
//     Number of returned entries (fewer can be returned, but not more than 1000). If the limit sent exceeds 1000, a limit of 10 is applied (by default).
// page_no:
//     Page number. Default is 1. Min value is 1.
// rank_field:
//     Rating category
// date:
//     Ratings calculation date. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
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
func (client *Client) WotClanratingsTop(realm Realm, limit int, pageNo int, rankField string, date wgnTime.UnixTime, fields []string, language string) (*wot.ClanratingsTop, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"rank_field": rankField,
		"date": strconv.FormatInt(date.Unix(), 10),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wot.ClanratingsTop
	err := client.doGetDataRequest(realm, "/wot/clanratings/top/", reqParam, &result)
	return result, err
}
