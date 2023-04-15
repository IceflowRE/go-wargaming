// Auto generated file!

package wowp

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type RatingsTopOptions struct {
	// Ratings calculation date. Up to 7 days before the current date; default value: yesterday. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
	Date *wgnTime.UnixTime `json:"date,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "ru" - Русский
	// "pl" - Polski
	// "de" - Deutsch
	// "fr" - Français
	// "es" - Español
	// "zh-cn" - 简体中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
	// Maximum number of top players on the list. Default is 10. Min value is 1. Maximum value: 1000.
	Limit *int `json:"limit,omitempty"`
	// Result page number. Default is 1. Min value is 1.
	PageNo *int `json:"page_no,omitempty"`
}

type RatingsTop struct {
	// Player account ID
	AccountId *int `json:"account_id,omitempty"`
	// Battles fought
	BattlesCount *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"battles_count,omitempty"`
	// Total damage caused to aircraft
	DamageDealt *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"damage_dealt,omitempty"`
	// Total damage caused to targets
	DamageDealtGround *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"damage_dealt_ground,omitempty"`
	// Enemy aircraft destroyed
	FragsCount *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"frags_count,omitempty"`
	// Total number of targets destroyed
	FragsCountGround *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"frags_count_ground,omitempty"`
	// Victories/Battles ratio
	WinsRatio *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"wins_ratio,omitempty"`
	// Average experience per battle
	XpAvg *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"xp_avg,omitempty"`
	// Maximum experience per battle
	XpMax *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"xp_max,omitempty"`
}
