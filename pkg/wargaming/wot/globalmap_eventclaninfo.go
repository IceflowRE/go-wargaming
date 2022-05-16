package wot

type GlobalmapEventclaninfo struct {
	// Clan info by events and Fronts
	Events struct {
		// Victories
		Wins int `json:"wins,omitempty"`
		// Link to Front
		Url string `json:"url,omitempty"`
		// Fame Points for completing a clan task
		TaskFamePoints int `json:"task_fame_points,omitempty"`
		// Rating changes during previous turn
		RankDelta int `json:"rank_delta,omitempty"`
		// Current rating
		Rank int `json:"rank,omitempty"`
		// Front ID
		FrontId string `json:"front_id,omitempty"`
		// Change of Fame Points since last turn calculation
		FamePointsSinceTurn int `json:"fame_points_since_turn,omitempty"`
		// Total Fame Points
		FamePoints int `json:"fame_points,omitempty"`
		// Event ID
		EventId string `json:"event_id,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Battle Fame Points
		BattleFamePoints int `json:"battle_fame_points,omitempty"`
	} `json:"events,omitempty"`
} 
