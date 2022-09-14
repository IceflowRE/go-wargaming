// Auto generated file!

package wowp

import (
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgnTime"
)

// RatingsNeighborsOptions options.
type RatingsNeighborsOptions struct {
	// Ratings calculation date. Up to 7 days before the current date; default value: yesterday. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
	Date *wgnTime.UnixTime
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Localization language. Default is "ru". Valid values:
	//
	// "en" - English
	// "ru" - Русский (by default)
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
	Language *string
	// Maximum number of adjacent positions in rating. Default is 5. Min value is 1. Maximum value: 50.
	Limit *int
}

type RatingsNeighbors struct {
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
