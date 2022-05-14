package wargaming

import (
	"strings"
	"strconv"
)

type WotRatingsTop struct {
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
	// Number of battles left to be included in ratings
	BattlesToPlay int `json:"battles_to_play,omitempty"`
	// Battles fought
	BattlesCount struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"battles_count,omitempty"`
	// Base capture points
	CapturePoints struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"capture_points,omitempty"`
	// Average damage caused per battle
	DamageAvg struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"damage_avg,omitempty"`
	// Total damage caused
	DamageDealt struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"damage_dealt,omitempty"`
	// Average number of vehicles destroyed per battle
	FragsAvg struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"frags_avg,omitempty"`
	// Vehicles destroyed
	FragsCount struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"frags_count,omitempty"`
	// Personal rating
	GlobalRating struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"global_rating,omitempty"`
	// Hit ratio
	HitsRatio struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"hits_ratio,omitempty"`
	// Average number of vehicles spotted per battle
	SpottedAvg struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"spotted_avg,omitempty"`
	// Vehicles spotted
	SpottedCount struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"spotted_count,omitempty"`
	// Survive ratio
	SurvivedRatio struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"survived_ratio,omitempty"`
	// Victories/Battles ratio
	WinsRatio struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"wins_ratio,omitempty"`
	// Total experience
	XpAmount struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"xp_amount,omitempty"`
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

// WotRatingsTop Deprecated: Attention! The method is deprecated.
// Method returns list of top players by specified parameter.
//
// https://developers.wargaming.net/reference/all/wot/ratings/top
//
// limit:
//     Number of returned entries. Default is 10. Min value is 1. Maximum value: 1000.
// page_no:
//     Result page number. Default is 1. Min value is 1.
// rank_field:
//     Rating category
// date:
//     Ratings calculation date. Up to 7 days before the current date; default value: yesterday. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
// type:
//     Rating period. For valid values, check the Types of ratings method.
// battle_type:
//     Battle types. Default is "default". Valid values:
//     
//     "company" &mdash; Tank Company Battles 
//     "random" &mdash; Random Battles 
//     "team" &mdash; Team Battles 
//     "default" &mdash; any battle type (by default)
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
func (client *Client) WotRatingsTop(realm Realm, limit int, pageNo int, rankField string, date UnixTime, type_ string, battleType string, fields []string, language string) (*WotRatingsTop, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"rank_field": rankField,
		"date": strconv.FormatInt(date.Unix(), 10),
		"type": type_,
		"battle_type": battleType,
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotRatingsTop
	err := client.doGetDataRequest(realm, "/wot/ratings/top/", reqParam, &result)
	return result, err
}
