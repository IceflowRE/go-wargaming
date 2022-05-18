package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strings"
)

// WotRatingsDates Method returns dates with available rating data.
//
// https://developers.wargaming.net/reference/all/wot/ratings/dates
//
// Deprecated: Attention! The method is deprecated.
//
// accountId:
//     Player account ID. Maximum limit: 100.
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
func (client *Client) WotRatingsDates(realm Realm, accountId []int, type_ string, battleType string, fields []string, language string) (*wot.RatingsDates, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"type": type_,
		"battle_type": battleType,
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wot.RatingsDates
	err := client.doGetDataRequest(realm, "/wot/ratings/dates/", reqParam, &result)
	return result, err
}
