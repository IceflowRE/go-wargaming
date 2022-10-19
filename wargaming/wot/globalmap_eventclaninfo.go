// Auto generated file!

package wot

type GlobalmapEventclaninfoOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
}

type GlobalmapEventclaninfo struct {
	// Clan info by events and Fronts
	Events *struct {
		// Battle Fame Points
		BattleFamePoints *int `json:"battle_fame_points,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Event ID
		EventId *string `json:"event_id,omitempty"`
		// Total Fame Points
		FamePoints *int `json:"fame_points,omitempty"`
		// Change of Fame Points since last turn calculation
		FamePointsSinceTurn *int `json:"fame_points_since_turn,omitempty"`
		// Front ID
		FrontId *string `json:"front_id,omitempty"`
		// Current rating
		Rank *int `json:"rank,omitempty"`
		// Rating changes during previous turn
		RankDelta *int `json:"rank_delta,omitempty"`
		// Fame Points for completing a clan task
		TaskFamePoints *int `json:"task_fame_points,omitempty"`
		// Link to Front
		Url *string `json:"url,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
	} `json:"events,omitempty"`
}
