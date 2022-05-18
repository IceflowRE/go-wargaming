package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotRatingsNeighbors Method returns list of adjacent positions in specified rating.
//
// https://developers.wargaming.net/reference/all/wot/ratings/neighbors
//
// Deprecated: Attention! The method is deprecated.
//
// accountId:
//     Player account ID
// limit:
//     Number of returned entries. Default is 5. Min value is 1. Maximum value: 50.
// rankField:
//     Rating category
// date:
//     Ratings calculation date. Up to 7 days before the current date; default value: yesterday. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
// type_:
//     Rating period. For valid values, check the Types of ratings method.
// battleType:
//     Battle types. Default is "default". Valid values:
//     
//     "company" - Tank Company Battles 
//     "random" - Random Battles 
//     "team" - Team Battles 
//     "default" - any battle type (by default)
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
func (client *Client) WotRatingsNeighbors(realm Realm, accountId int, limit int, rankField string, date wgnTime.UnixTime, type_ string, battleType string, fields []string, language string) (*wot.RatingsNeighbors, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"limit": strconv.Itoa(limit),
		"rank_field": rankField,
		"date": strconv.FormatInt(date.Unix(), 10),
		"type": type_,
		"battle_type": battleType,
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wot.RatingsNeighbors
	err := client.doGetDataRequest(realm, "/wot/ratings/neighbors/", reqParam, &result)
	return result, err
}
