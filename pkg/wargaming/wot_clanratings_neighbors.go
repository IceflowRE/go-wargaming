package wargaming

import (
	"strings"
	"strconv"
)

type WotClanratingsNeighbors struct {
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Clan name
	ClanName string `json:"clan_name,omitempty"`
	// Clan tag
	ClanTag string `json:"clan_tag,omitempty"`
	// Reasons why specified rating categories were not calculated. Contains data in "key-value" format, where the key is category name and the value is reason.
	// Possible reasons:
	// 
	// inactivity - Inactivity for 28 days
	// newbies_measure - Under 10 members in the clan
	// limits - Rank conditions not met
	// blocked - Clan blocked
	// other - Technical reasons
	ExcludeReasons map[string]string `json:"exclude_reasons,omitempty"`
	// Average number of battles
	BattlesCountAvg struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"battles_count_avg,omitempty"`
	// Average number of battles per day
	BattlesCountAvgDaily struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"battles_count_avg_daily,omitempty"`
	// Indicator of clan's performance.
	Efficiency struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"efficiency,omitempty"`
	// Weighted Elo rating achieved in the Stronghold mode
	FbEloRating struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"fb_elo_rating,omitempty"`
	// Elo rating achieved by the clan on Tier X vehicles in the Stronghold mode
	FbEloRating10 struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"fb_elo_rating_10,omitempty"`
	// Elo rating achieved on Tier VI vehicles in the Stronghold mode
	FbEloRating6 struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"fb_elo_rating_6,omitempty"`
	// Elo rating achieved on Tier VIII vehicles in the Stronghold mode
	FbEloRating8 struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"fb_elo_rating_8,omitempty"`
	// Average global rating value
	GlobalRatingAvg struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"global_rating_avg,omitempty"`
	// Weighted average of global rating value
	GlobalRatingWeightedAvg struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"global_rating_weighted_avg,omitempty"`
	// Elo rating on the Global Map
	GmEloRating struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"gm_elo_rating,omitempty"`
	// Elo rating on the Global Map in Absolute division
	GmEloRating10 struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"gm_elo_rating_10,omitempty"`
	// Elo rating on the Global Map in Medium division
	GmEloRating6 struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"gm_elo_rating_6,omitempty"`
	// Elo rating on the Global Map in Champion division
	GmEloRating8 struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value int `json:"value,omitempty"`
	} `json:"gm_elo_rating_8,omitempty"`
	// Rating in Battles for Stronghold
	RatingFort struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"rating_fort,omitempty"`
	// Average number of vehicles of Tier 10 per clan member
	V10LAvg struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"v10l_avg,omitempty"`
	// Average victory rate
	WinsRatioAvg struct {
		// Position
		Rank int `json:"rank,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Absolute value
		Value float32 `json:"value,omitempty"`
	} `json:"wins_ratio_avg,omitempty"`
}

// WotClanratingsNeighbors Method returns list of adjacent positions in specified clan rating.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/neighbors
//
// clan_id:
//     Clan ID
// rank_field:
//     Rating category
// date:
//     Ratings calculation date. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
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
// limit:
//     Number of returned entries (fewer can be returned, but not more than 50). If the limit sent exceeds 50, a limit of 5 is applied (by default).
func (client *Client) WotClanratingsNeighbors(realm Realm, clanId int, rankField string, date UnixTime, fields []string, language string, limit int) (*WotClanratingsNeighbors, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": strconv.Itoa(clanId),
		"rank_field": rankField,
		"date": strconv.FormatInt(date.Unix(), 10),
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
	}

	var result *WotClanratingsNeighbors
	err := client.doGetDataRequest(realm, "/wot/clanratings/neighbors/", reqParam, &result)
	return result, err
}
