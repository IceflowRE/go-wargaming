package wot

type ClanratingsTop struct {
	// Average victory rate
	WinsRatioAvg struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"wins_ratio_avg,omitempty"`
	// Average number of vehicles of Tier 10 per clan member
	V10LAvg struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"v10l_avg,omitempty"`
	// Rating in Battles for Stronghold
	RatingFort struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"rating_fort,omitempty"`
	// Elo rating on the Global Map in Champion division
	GmEloRating8 struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"gm_elo_rating_8,omitempty"`
	// Elo rating on the Global Map in Medium division
	GmEloRating6 struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"gm_elo_rating_6,omitempty"`
	// Elo rating on the Global Map in Absolute division
	GmEloRating10 struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"gm_elo_rating_10,omitempty"`
	// Elo rating on the Global Map
	GmEloRating struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"gm_elo_rating,omitempty"`
	// Weighted average of global rating value
	GlobalRatingWeightedAvg struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"global_rating_weighted_avg,omitempty"`
	// Average global rating value
	GlobalRatingAvg struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"global_rating_avg,omitempty"`
	// Elo rating achieved on Tier VIII vehicles in the Stronghold mode
	FbEloRating8 struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"fb_elo_rating_8,omitempty"`
	// Elo rating achieved on Tier VI vehicles in the Stronghold mode
	FbEloRating6 struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"fb_elo_rating_6,omitempty"`
	// Elo rating achieved by the clan on Tier X vehicles in the Stronghold mode
	FbEloRating10 struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"fb_elo_rating_10,omitempty"`
	// Weighted Elo rating achieved in the Stronghold mode
	FbEloRating struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"fb_elo_rating,omitempty"`
	// Reasons why specified rating categories were not calculated. Contains data in "key-value" format, where the key is category name and the value is reason.
	// Possible reasons:
	// 
	// inactivity - Inactivity for 28 days
	// newbies_measure - Under 10 members in the clan
	// limits - Rank conditions not met
	// blocked - Clan blocked
	// other - Technical reasons
	ExcludeReasons map[string]string `json:"exclude_reasons,omitempty"`
	// Indicator of clan's performance.
	Efficiency struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"efficiency,omitempty"`
	// Clan tag
	ClanTag string `json:"clan_tag,omitempty"`
	// Clan name
	ClanName string `json:"clan_name,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Average number of battles per day
	BattlesCountAvgDaily struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"battles_count_avg_daily,omitempty"`
	// Average number of battles
	BattlesCountAvg struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"battles_count_avg,omitempty"`
} 
