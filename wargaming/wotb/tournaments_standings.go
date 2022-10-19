// Auto generated file!

package wotb

type TournamentsStandingsOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Allows to get all team standings starting from a specific place, including this place
	FromPosition *int `json:"from_position,omitempty"`
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
	// Team ID. Maximum limit: 10.
	TeamId []int `json:"team_id,omitempty"`
	// Allows to get all team standings up to a specific place, including this place
	ToPosition *int `json:"to_position,omitempty"`
}

type TournamentsStandings struct {
	// Number of battles played by a team
	BattlePlayed *int `json:"battle_played,omitempty"`
	// Number of battles drawn by a team
	Draws *int `json:"draws,omitempty"`
	// ID of a team's group
	GroupId *int `json:"group_id,omitempty"`
	// Number of battles lost by a team
	Losses *int `json:"losses,omitempty"`
	// Number of points earned by a team
	Points *int `json:"points,omitempty"`
	// Team's place
	Position *int `json:"position,omitempty"`
	// Stage ID
	StageId *int `json:"stage_id,omitempty"`
	// Team ID
	TeamId *int `json:"team_id,omitempty"`
	// Number of battles won by a team
	Wins *int `json:"wins,omitempty"`
}
