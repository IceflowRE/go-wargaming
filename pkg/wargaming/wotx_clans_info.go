package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotx"
	"strings"
)

// WotxClansInfo Method returns detailed clan information.
//
// https://developers.wargaming.net/reference/all/wotx/clans/info
//
// clanId:
//     Clan ID. Maximum limit: 100. Min value is 1.
//     Parameter is required.
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "members"
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "en". Valid values:
//     
//     "en" - English (by default)
//     "ru" - Русский 
//     "pl" - Polski 
//     "de" - Deutsch 
//     "fr" - Français 
//     "es" - Español 
//     "tr" - Türkçe
func (client *Client) WotxClansInfo(realm Realm, clanId []int, extra []string, fields []string, language string) (*wotx.ClansInfo, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": utils.SliceIntToString(clanId, ","),
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wotx.ClansInfo
	err := client.doGetDataRequest(realm, "/wotx/clans/info/", reqParam, &result)
	return result, err
}
