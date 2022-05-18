package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotClanratingsNeighbors Method returns list of adjacent positions in specified clan rating.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/neighbors
//
// clanId:
//     Clan ID
// rankField:
//     Rating category
// date:
//     Ratings calculation date. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
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
//     Number of returned entries (fewer can be returned, but not more than 50). If the limit sent exceeds 50, a limit of 5 is applied (by default).
func (client *Client) WotClanratingsNeighbors(realm Realm, clanId int, rankField string, date wgnTime.UnixTime, fields []string, language string, limit int) (*wot.ClanratingsNeighbors, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": strconv.Itoa(clanId),
		"rank_field": rankField,
		"date": strconv.FormatInt(date.Unix(), 10),
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
	}

	var result *wot.ClanratingsNeighbors
	err := client.doGetDataRequest(realm, "/wot/clanratings/neighbors/", reqParam, &result)
	return result, err
}
