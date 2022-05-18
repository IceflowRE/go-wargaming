package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strings"
)

// WotbClanmessagesLikes Method returns message likes on clan message board.
//
// https://developers.wargaming.net/reference/all/wotb/clanmessages/likes
//
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// messageId:
//     Message ID. Maximum limit: 10.
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
func (client *Client) WotbClanmessagesLikes(realm Realm, accessToken string, messageId []int, fields []string, language string) (*wotb.ClanmessagesLikes, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"message_id": utils.SliceIntToString(messageId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wotb.ClanmessagesLikes
	err := client.doGetDataRequest(realm, "/wotb/clanmessages/likes/", reqParam, &result)
	return result, err
}
