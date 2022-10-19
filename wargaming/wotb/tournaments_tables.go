// Auto generated file!

package wotb

type TournamentsTablesOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Group ID. Maximum limit: 10.
	GroupId []int `json:"group_id,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "cs" - Čeština
	// "de" - Deutsch
	// "es" - Español
	// "fr" - Français
	// "pl" - Polski
	// "tr" - Türkçe
	Language *string `json:"language,omitempty"`
	// Number of returned entries. Default is 10. Min value is 1. Maximum value: 25.
	Limit *int `json:"limit,omitempty"`
	// Result page number. Default is 1. Min value is 1.
	PageNo *int `json:"page_no,omitempty"`
}

type TournamentsTables struct {
	// ID of a default clan emblem
	ClanEmblemPresetId *int `json:"clan_emblem_preset_id,omitempty"`
	// Team clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Name of a team's clan
	ClanLabel *string `json:"clan_label,omitempty"`
	// ID of a team's group
	GroupId *int `json:"group_id,omitempty"`
	// Sequence number of a group in a stage
	GroupOrder *int `json:"group_order,omitempty"`
	// Number of matches played by a team
	MatchesPlayed *int `json:"matches_played,omitempty"`
	// Team's place
	Position *int `json:"position,omitempty"`
	// Number of the tour in which a team exited the tournament; relevant only for Single Elimination (SE) and Double Elimination (DE) tournament brackets
	Round *int `json:"round,omitempty"`
	// Stage ID
	StageId *int `json:"stage_id,omitempty"`
	// Team ID
	TeamId *int `json:"team_id,omitempty"`
	// Points earned by a team; relevant only for "Round Robin" tournament brackets
	TeamPoints *int `json:"team_points,omitempty"`
	// Team name
	Title *string `json:"title,omitempty"`
	// Tournament ID
	TournamentId *int `json:"tournament_id,omitempty"`
}
