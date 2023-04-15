// Auto generated file!

package wotb

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type TournamentsInfoOptions struct {
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
}

type TournamentsInfo struct {
	// Award for participating in tournament
	Award *struct {
		// Award amount
		Amount *int `json:"amount,omitempty"`
		// Award currency: Free XP, gold or credits
		Currency *string `json:"currency,omitempty"`
	} `json:"award,omitempty"`
	// Tournament description
	Description *string `json:"description,omitempty"`
	// Tournament end date and time
	EndAt *wgnTime.UnixTime `json:"end_at,omitempty"`
	// Fee for participating in tournament
	Fee *struct {
		// Fee amount
		Amount *int `json:"amount,omitempty"`
		// Fee currency: Free XP, gold or credits
		Currency *string `json:"currency,omitempty"`
	} `json:"fee,omitempty"`
	// Tournament Logo
	Logo *struct {
		// Link to logo
		Original *string `json:"original,omitempty"`
		// Link to preview
		Preview *string `json:"preview,omitempty"`
	} `json:"logo,omitempty"`
	// Matches start date and time
	MatchesStartAt *wgnTime.UnixTime `json:"matches_start_at,omitempty"`
	// Maximum number of players per team in tournament
	MaxPlayersCount *int `json:"max_players_count,omitempty"`
	// Section contains information about links to tournament media resources
	MediaLinks *struct {
		// ID of media resource link
		Id *string `json:"id,omitempty"`
		// Image of media resource link; available only for links of "Custom" type
		Image *string `json:"image,omitempty"`
		// Type of media resource link
		Kind *string `json:"kind,omitempty"`
		// Link to media resource
		Url *string `json:"url,omitempty"`
	} `json:"media_links,omitempty"`
	// Minimum number of players per team in tournament
	MinPlayersCount *int `json:"min_players_count,omitempty"`
	// Tournament other regulations
	OtherRules *string `json:"other_rules,omitempty"`
	// Tournament award description
	PrizeDescription *string `json:"prize_description,omitempty"`
	// Registration end date and time
	RegistrationEndAt *wgnTime.UnixTime `json:"registration_end_at,omitempty"`
	// Registration start date and time
	RegistrationStartAt *wgnTime.UnixTime `json:"registration_start_at,omitempty"`
	// Tournament rules and regulations
	Rules *string `json:"rules,omitempty"`
	// Tournament start date and time
	StartAt *wgnTime.UnixTime `json:"start_at,omitempty"`
	// Tournament status
	Status *string `json:"status,omitempty"`
	// Total number of teams in tournament, both already registered and in process of forming
	Teams *struct {
		// Number of confirmed teams in tournament
		Confirmed *int `json:"confirmed,omitempty"`
		// Maximum number of teams available in tournament
		Max *int `json:"max,omitempty"`
		// Minimum number of teams in tournament
		Min *int `json:"min,omitempty"`
		// Total number of teams in tournament, both already registered and in process of forming
		Total *int `json:"total,omitempty"`
	} `json:"teams,omitempty"`
	// Tournament name
	Title *string `json:"title,omitempty"`
	// Tournament id
	TournamentId *int `json:"tournament_id,omitempty"`
	// Award for winning tournament
	WinnerAward *struct {
		// Winner Award amount
		Amount *int `json:"amount,omitempty"`
		// Winner Award currency: Free XP, gold or credits
		Currency *string `json:"currency,omitempty"`
	} `json:"winner_award,omitempty"`
}
