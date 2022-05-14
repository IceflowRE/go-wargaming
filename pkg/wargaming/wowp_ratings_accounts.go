package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
	"strconv"
)

type WowpRatingsAccounts struct {
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
	// Battles fought
	BattlesCount struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"battles_count,omitempty"`
	// Total damage caused to aircraft
	DamageDealt struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"damage_dealt,omitempty"`
	// Total damage caused to targets
	DamageDealtGround struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"damage_dealt_ground,omitempty"`
	// Enemy aircraft destroyed
	FragsCount struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"frags_count,omitempty"`
	// Total number of targets destroyed
	FragsCountGround struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"frags_count_ground,omitempty"`
	// Victories/Battles ratio
	WinsRatio struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"wins_ratio,omitempty"`
	// Average experience per battle
	XpAvg struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"xp_avg,omitempty"`
	// Maximum experience per battle
	XpMax struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"xp_max,omitempty"`
}

// WowpRatingsAccounts Method returns player ratings by specified IDs.
//
// https://developers.wargaming.net/reference/all/wowp/ratings/accounts
//
// account_id:
//     Player account ID. Maximum limit: 100.
// date:
//     Ratings calculation date. Up to 7 days before the current date; default value: yesterday. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
// type:
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
func (client *Client) WowpRatingsAccounts(realm Realm, accountId []int, date UnixTime, type_ string, fields []string, language string) (*WowpRatingsAccounts, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"date": strconv.FormatInt(date.Unix(), 10),
		"type": type_,
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowpRatingsAccounts
	err := client.doGetDataRequest(realm, "/wowp/ratings/accounts/", reqParam, &result)
	return result, err
}
