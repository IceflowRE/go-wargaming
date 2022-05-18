package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strings"
)

// WotEncyclopediaPersonalmissions Method returns details on Personal Missions on the basis of specified campaign IDs, operation IDs, mission branch and tag IDs.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/personalmissions
//
// campaignId:
//     Campaign ID. Maximum limit: 100.
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
// operationId:
//     Operation ID. Maximum limit: 100.
// setId:
//     Mission branch ID. Maximum limit: 100.
// tag:
//     Mission tag. Maximum limit: 100.
func (client *Client) WotEncyclopediaPersonalmissions(realm Realm, campaignId []int, fields []string, language string, operationId []int, setId []int, tag []string) (*wot.EncyclopediaPersonalmissions, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"campaign_id": utils.SliceIntToString(campaignId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
		"operation_id": utils.SliceIntToString(operationId, ","),
		"set_id": utils.SliceIntToString(setId, ","),
		"tag": strings.Join(tag, ","),
	}

	var result *wot.EncyclopediaPersonalmissions
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/personalmissions/", reqParam, &result)
	return result, err
}
