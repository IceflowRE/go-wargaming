package wowp

type RatingsAccounts struct {
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
	// Victories/Battles ratio
	WinsRatio struct {
		// Absolute value
		Value float32 `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"wins_ratio,omitempty"`
	// Total number of targets destroyed
	FragsCountGround struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"frags_count_ground,omitempty"`
	// Enemy aircraft destroyed
	FragsCount struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"frags_count,omitempty"`
	// Total damage caused to targets
	DamageDealtGround struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"damage_dealt_ground,omitempty"`
	// Total damage caused to aircraft
	DamageDealt struct {
		// Absolute value
		Value int `json:"value,omitempty"`
		// Change of position in rating
		RankDelta int `json:"rank_delta,omitempty"`
		// Position
		Rank int `json:"rank,omitempty"`
	} `json:"damage_dealt,omitempty"`
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
