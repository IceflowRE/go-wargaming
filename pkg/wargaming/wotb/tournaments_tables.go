package wotb

type TournamentsTables struct {
	// Tournament ID
	TournamentId int `json:"tournament_id,omitempty"`
	// Team name
	Title string `json:"title,omitempty"`
	// Points earned by a team; relevant only for "Round Robin" tournament brackets
	TeamPoints int `json:"team_points,omitempty"`
	// Team ID
	TeamId int `json:"team_id,omitempty"`
	// Stage ID
	StageId int `json:"stage_id,omitempty"`
	// Number of the tour in which a team exited the tournament; relevant only for Single Elimination (SE) and Double Elimination (DE) tournament brackets
	Round int `json:"round,omitempty"`
	// Team's place
	Position int `json:"position,omitempty"`
	// Number of matches played by a team
	MatchesPlayed int `json:"matches_played,omitempty"`
	// Sequence number of a group in a stage
	GroupOrder int `json:"group_order,omitempty"`
	// ID of a team's group
	GroupId int `json:"group_id,omitempty"`
	// Name of a team's clan
	ClanLabel string `json:"clan_label,omitempty"`
	// Team clan ID
	ClanId int `json:"clan_id,omitempty"`
	// ID of a default clan emblem
	ClanEmblemPresetId int `json:"clan_emblem_preset_id,omitempty"`
} 
