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
// campaign_id:
//     Campaign ID. Maximum limit: 100.
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
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
// operation_id:
//     Operation ID. Maximum limit: 100.
// set_id:
//     Mission branch ID. Maximum limit: 100.
// tag:
//     Mission tag. Maximum limit: 100.
func (client *Client) WotEncyclopediaPersonalmissions(realm Realm, campaignId []int, fields []string, language string, operationId []int, setId []int, tag []string) (*wot.EncyclopediaPersonalmissions, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
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
