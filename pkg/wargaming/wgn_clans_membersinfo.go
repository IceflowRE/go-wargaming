package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgn"
	"strings"
)

// WgnClansMembersinfo Method returns detailed clan member information and brief clan details. Information is available for World of Tanks and World of Warplanes clans.This method will be removed. Use method Clan member details (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wgn/clans/membersinfo
//
// Deprecated: Attention! The method is deprecated.
//
// accountId:
//     Account ID. Maximum limit: 100. Min value is 1.
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
func (client *Client) WgnClansMembersinfo(realm Realm, accountId []int, fields []string, game string, language string) (*wgn.ClansMembersinfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"fields": strings.Join(fields, ","),
		"game": game,
		"language": language,
	}

	var result *wgn.ClansMembersinfo
	err := client.doGetDataRequest(realm, "/wgn/clans/membersinfo/", reqParam, &result)
	return result, err
}
