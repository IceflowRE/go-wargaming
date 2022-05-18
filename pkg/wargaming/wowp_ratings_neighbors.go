package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wowp"
	"strconv"
	"strings"
)

// WowpRatingsNeighbors Method returns list of adjacent positions in specified rating.
//
// https://developers.wargaming.net/reference/all/wowp/ratings/neighbors
//
// accountId:
//     Player account ID
// date:
//     Ratings calculation date. Up to 7 days before the current date; default value: yesterday. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
// limit:
//     Maximum number of adjacent positions in rating. Default is 5. Min value is 1. Maximum value: 50.
// rankField:
//     Rating category
// type_:
//     Rating period. For valid values, check the Types of ratings method.
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
func (client *Client) WowpRatingsNeighbors(realm Realm, accountId int, date wgnTime.UnixTime, limit int, rankField string, type_ string, fields []string, language string) (*wowp.RatingsNeighbors, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"date": strconv.FormatInt(date.Unix(), 10),
		"limit": strconv.Itoa(limit),
		"rank_field": rankField,
		"type": type_,
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wowp.RatingsNeighbors
	err := client.doGetDataRequest(realm, "/wowp/ratings/neighbors/", reqParam, &result)
	return result, err
}
