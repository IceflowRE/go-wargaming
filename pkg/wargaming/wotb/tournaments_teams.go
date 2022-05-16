package wotb

type TournamentsTeams struct {
	// Tournament ID
	TournamentId int `json:"tournament_id,omitempty"`
	// Team name
	Title string `json:"title,omitempty"`
	// Team ID
	TeamId int `json:"team_id,omitempty"`
	// Team status
	Status string `json:"status,omitempty"`
	// Information on team players
	Players struct {
		// Technical position name
		Role string `json:"role,omitempty"`
		// Player name
		Name string `json:"name,omitempty"`
		// Link to player image
		Image string `json:"image,omitempty"`
		// Player account ID
		AccountId int `json:"account_id,omitempty"`
	} `json:"players,omitempty"`
	// Team clan ID
	ClanId int `json:"clan_id,omitempty"`
} 
