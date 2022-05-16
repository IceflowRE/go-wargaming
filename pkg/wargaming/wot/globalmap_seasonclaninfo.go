package wot

type GlobalmapSeasonclaninfo struct {
	// Clan info by seasons and vehicle Tiers
	Seasons struct {
		// Victories
		Wins int `json:"wins,omitempty"`
		// Change of Victory Points since last turn calculation
		VictoryPointsSinceTurn int `json:"victory_points_since_turn,omitempty"`
		// Victory Points
		VictoryPoints int `json:"victory_points,omitempty"`
		// Vehicle Tier
		VehicleLevel int `json:"vehicle_level,omitempty"`
		// Rating changes during previous turn
		RankDelta int `json:"rank_delta,omitempty"`
		// Current rating
		Rank int `json:"rank,omitempty"`
		// Elo rating
		Elo int `json:"elo,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
	} `json:"seasons,omitempty"`
} 
