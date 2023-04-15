// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type RatingsNeighborsOptions struct {
	// Battle types. Default is "default". Valid values:
	//
	// "company" - Tank Company Battles
	// "random" - Random Battles
	// "team" - Team Battles
	// "default" - any battle type (by default)
	BattleType *string `json:"battle_type,omitempty"`
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
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
	// Number of returned entries. Default is 5. Min value is 1. Maximum value: 50.
	Limit *int `json:"limit,omitempty"`
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
	// Number of battles left to be included in ratings
	BattlesToPlay *int `json:"battles_to_play,omitempty"`
	// Base capture points
	CapturePoints *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"capture_points,omitempty"`
	// Average damage caused per battle
	DamageAvg *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"damage_avg,omitempty"`
	// Total damage caused
	DamageDealt *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"damage_dealt,omitempty"`
	// Average number of vehicles destroyed per battle
	FragsAvg *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"frags_avg,omitempty"`
	// Vehicles destroyed
	FragsCount *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"frags_count,omitempty"`
	// Personal rating
	GlobalRating *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"global_rating,omitempty"`
	// Hit ratio
	HitsRatio *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"hits_ratio,omitempty"`
	// Average number of vehicles spotted per battle
	SpottedAvg *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"spotted_avg,omitempty"`
	// Vehicles spotted
	SpottedCount *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"spotted_count,omitempty"`
	// Survive ratio
	SurvivedRatio *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"survived_ratio,omitempty"`
	// Victories/Battles ratio
	WinsRatio *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"wins_ratio,omitempty"`
	// Total experience
	XpAmount *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"xp_amount,omitempty"`
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
