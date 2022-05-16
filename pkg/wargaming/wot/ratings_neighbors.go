package wot

type RatingsNeighbors struct {
	// Maximum experience per battle
	XpMax struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"xp_max,omitempty"`
	// Average experience per battle
	XpAvg struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"xp_avg,omitempty"`
	// Total experience
	XpAmount struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"xp_amount,omitempty"`
	// Victories/Battles ratio
	WinsRatio struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"wins_ratio,omitempty"`
	// Survive ratio
	SurvivedRatio struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"survived_ratio,omitempty"`
	// Vehicles spotted
	SpottedCount struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"spotted_count,omitempty"`
	// Average number of vehicles spotted per battle
	SpottedAvg struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"spotted_avg,omitempty"`
	// Hit ratio
	HitsRatio struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"hits_ratio,omitempty"`
	// Personal rating
	GlobalRating struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"global_rating,omitempty"`
	// Vehicles destroyed
	FragsCount struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"frags_count,omitempty"`
	// Average number of vehicles destroyed per battle
	FragsAvg struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"frags_avg,omitempty"`
	// Total damage caused
	DamageDealt struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"damage_dealt,omitempty"`
	// Average damage caused per battle
	DamageAvg struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"damage_avg,omitempty"`
	// Base capture points
	CapturePoints struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"capture_points,omitempty"`
	// Number of battles left to be included in ratings
	BattlesToPlay int `json:"battles_to_play,omitempty"`
	// Battles fought
	BattlesCount struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"battles_count,omitempty"`
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
} 
