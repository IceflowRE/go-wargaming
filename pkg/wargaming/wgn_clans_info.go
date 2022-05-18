package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgn"
	"strings"
)

// WgnClansInfo Method returns detailed clan information. Information is available for World of Tanks and World of Warplanes clans.This method will be removed. Use method Clan details (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wgn/clans/info
//
// Deprecated: Attention! The method is deprecated.
//
// clanId:
//     Clan ID. Maximum limit: 100.
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "private.online_members"
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// game:
//     Name of the game to perform the clan search in. If the parameter is not specified, search will be executed across World of Tanks. Default is "wot". Valid values:
//     
//     "wot" - World of Tanks (by default)
//     "wowp" - World of Warplanes
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
func (client *Client) WgnClansInfo(realm Realm, clanId []int, accessToken string, extra []string, fields []string, game string, language string, membersKey string) (*wgn.ClansInfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": utils.SliceIntToString(clanId, ","),
		"access_token": accessToken,
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"game": game,
		"language": language,
		"members_key": membersKey,
	}

	var result *wgn.ClansInfo
	err := client.doGetDataRequest(realm, "/wgn/clans/info/", reqParam, &result)
	return result, err
}
