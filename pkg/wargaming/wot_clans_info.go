package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strings"
)

// WotClansInfo Method returns detailed clan information.
//
// https://developers.wargaming.net/reference/all/wot/clans/info
//
// clanId:
//     Clan ID. Maximum limit: 100.
//     Parameter is required.
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "private.online_members"
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
// membersKey:
//     This parameter changes members field type. Valid values:
//     
//     "id" - Members field will contain associative array with account_id indexing in response
func (client *Client) WotClansInfo(realm Realm, clanId []int, accessToken string, extra []string, fields []string, language string, membersKey string) (*wot.ClansInfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": utils.SliceIntToString(clanId, ","),
		"access_token": accessToken,
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
		"members_key": membersKey,
	}

	var result *wot.ClansInfo
	err := client.doGetDataRequest(realm, "/wot/clans/info/", reqParam, &result)
	return result, err
}
