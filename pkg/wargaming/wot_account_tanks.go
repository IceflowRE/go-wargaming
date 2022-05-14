package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotAccountTanks struct {
	// Mastery Badges:
	// 
	// 0 — None
	// 1 — 3rd Class 
	// 2 — 2nd Class
	// 3 — 1st Class
	// 4 — Ace Tanker
	MarkOfMastery int `json:"mark_of_mastery,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Vehicle statistics
	Statistics struct {
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
	} `json:"statistics,omitempty"`
}

// WotAccountTanks Method returns details on player's vehicles.
//
// https://developers.wargaming.net/reference/all/wot/account/tanks
//
// account_id:
//     Player account ID. Maximum limit: 100.
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
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
// tank_id:
//     Player's vehicle ID. Maximum limit: 100.
func (client *Client) WotAccountTanks(realm Realm, accountId []int, accessToken string, fields []string, language string, tankId []int) (*WotAccountTanks, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
		"language": language,
		"tank_id": utils.SliceIntToString(tankId, ","),
	}

	var result *WotAccountTanks
	err := client.doGetDataRequest(realm, "/wot/account/tanks/", reqParam, &result)
	return result, err
}
