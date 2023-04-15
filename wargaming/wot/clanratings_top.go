// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type ClanratingsTopOptions struct {
	// Ratings calculation date. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
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
	// Number of returned entries (fewer can be returned, but not more than 1000). If the limit sent exceeds 1000, a limit of 10 is applied (by default).
	Limit *int `json:"limit,omitempty"`
	// Page number. Default is 1. Min value is 1.
	PageNo *int `json:"page_no,omitempty"`
}

type ClanratingsTop struct {
	// Average number of battles
	BattlesCountAvg *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"battles_count_avg,omitempty"`
	// Average number of battles per day
	BattlesCountAvgDaily *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"battles_count_avg_daily,omitempty"`
	// Clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Clan name
	ClanName *string `json:"clan_name,omitempty"`
	// Clan tag
	ClanTag *string `json:"clan_tag,omitempty"`
	// Indicator of clan's performance.
	Efficiency *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"efficiency,omitempty"`
	// Reasons why specified rating categories were not calculated. Contains data in "key-value" format, where the key is category name and the value is reason.
	// Possible reasons:
	//
	// inactivity - Inactivity for 28 days
	// newbies_measure - Under 10 members in the clan
	// limits - Rank conditions not met
	// blocked - Clan blocked
	// other - Technical reasons
	ExcludeReasons map[string]string `json:"exclude_reasons,omitempty"`
	// Weighted Elo rating achieved in the Stronghold mode
	FbEloRating *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"fb_elo_rating,omitempty"`
	// Elo rating achieved by the clan on Tier X vehicles in the Stronghold mode
	FbEloRating10 *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"fb_elo_rating_10,omitempty"`
	// Elo rating achieved on Tier VI vehicles in the Stronghold mode
	FbEloRating6 *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"fb_elo_rating_6,omitempty"`
	// Elo rating achieved on Tier VIII vehicles in the Stronghold mode
	FbEloRating8 *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"fb_elo_rating_8,omitempty"`
	// Average global rating value
	GlobalRatingAvg *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"global_rating_avg,omitempty"`
	// Weighted average of global rating value
	GlobalRatingWeightedAvg *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"global_rating_weighted_avg,omitempty"`
	// Elo rating on the Global Map
	GmEloRating *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"gm_elo_rating,omitempty"`
	// Elo rating on the Global Map in Absolute division
	GmEloRating10 *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"gm_elo_rating_10,omitempty"`
	// Elo rating on the Global Map in Medium division
	GmEloRating6 *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"gm_elo_rating_6,omitempty"`
	// Elo rating on the Global Map in Champion division
	GmEloRating8 *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *int `json:"value,omitempty"`
	} `json:"gm_elo_rating_8,omitempty"`
	// Rating in Battles for Stronghold
	RatingFort *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"rating_fort,omitempty"`
	// Average number of vehicles of Tier 10 per clan member
	V10lAvg *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"v10l_avg,omitempty"`
	// Average victory rate
	WinsRatioAvg *struct {
		// Position
		Rank *int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta *int `json:"rank_delta,omitempty"`
		// Absolute value
		Value *float32 `json:"value,omitempty"`
	} `json:"wins_ratio_avg,omitempty"`
}
