package wotb

type TournamentsStandings struct {
	// Number of battles won by a team
	Wins int `json:"wins,omitempty"`
	// Team ID
	TeamId int `json:"team_id,omitempty"`
	// Stage ID
	StageId int `json:"stage_id,omitempty"`
	// Team's place
	Position int `json:"position,omitempty"`
	// Number of points earned by a team
	Points int `json:"points,omitempty"`
	// Number of battles lost by a team
	Losses int `json:"losses,omitempty"`
	// ID of a team's group
	GroupId int `json:"group_id,omitempty"`
	// Number of battles drawn by a team
	Draws int `json:"draws,omitempty"`
	// Number of battles played by a team
	BattlePlayed int `json:"battle_played,omitempty"`
} 
