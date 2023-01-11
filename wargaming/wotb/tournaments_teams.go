// Auto generated file!

package wotb

type TournamentsTeamsOptions struct {
	// ID of the account that belongs to the team. Maximum limit: 100.
	AccountId []int `json:"account_id,omitempty"`
	// ID of the clan that owns the team. Maximum limit: 100.
	ClanId []int `json:"clan_id,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
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
	// Number of returned entries. Default is 10. Min value is 1. Maximum value: 100.
	Limit *int `json:"limit,omitempty"`
	// Result page number. Default is 1. Min value is 1.
	PageNo *int `json:"page_no,omitempty"`
	// First letters in team name for search. Minimum length: 2 characters. Maximum length: 50 characters.
	Search *string `json:"search,omitempty"`
	// Team status. Maximum limit: 100. Valid values:
	//
	// "forming" - team roster is not yet confirmed
	// "confirmed" - team roster is confirmed
	// "disqualified" - team is disqualified
	Status []string `json:"status,omitempty"`
	// Team ID. Maximum limit: 25.
	TeamId []int `json:"team_id,omitempty"`
}

type TournamentsTeams struct {
	// Team clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Information on team players
	Players []*struct {
		// Player account ID
		AccountId *int `json:"account_id,omitempty"`
		// Link to player image
		Image *string `json:"image,omitempty"`
		// Player name
		Name *string `json:"name,omitempty"`
		// Technical position name
		Role *string `json:"role,omitempty"`
	} `json:"players,omitempty"`
	// Team status
	Status *string `json:"status,omitempty"`
	// Team ID
	TeamId *int `json:"team_id,omitempty"`
	// Team name
	Title *string `json:"title,omitempty"`
	// Tournament ID
	TournamentId *int `json:"tournament_id,omitempty"`
}
