// Auto generated file!

package wot

type GlobalmapSeasonclaninfoOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
}

type GlobalmapSeasonclaninfo struct {
	// Clan info by seasons and vehicle Tiers
	Seasons *struct {
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Elo rating
		Elo *int `json:"elo,omitempty"`
		// Current rating
		Rank *int `json:"rank,omitempty"`
		// Rating changes during previous turn
		RankDelta *int `json:"rank_delta,omitempty"`
		// Vehicle Tier
		VehicleLevel *int `json:"vehicle_level,omitempty"`
		// Victory Points
		VictoryPoints *int `json:"victory_points,omitempty"`
		// Change of Victory Points since last turn calculation
		VictoryPointsSinceTurn *int `json:"victory_points_since_turn,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
	} `json:"seasons,omitempty"`
}
