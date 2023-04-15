// Auto generated file!

package wotb

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type TournamentsListOptions struct {
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
	// Page limit. Number of tournaments in one page. Default is 10. Min value is 1. Maximum value: 25.
	Limit *int `json:"limit,omitempty"`
	// Result page number. Default is 1. Min value is 1.
	PageNo *int `json:"page_no,omitempty"`
	// First letters in tournament name for search. Minimum length: 2 characters. Maximum length: 50 characters.
	Search *string `json:"search,omitempty"`
	// Tournament status. Maximum limit: 100. Valid values:
	//
	// "upcoming" - Tournament is planned
	// "registration_started" - Players can apply to join the tournament
	// "registration_finished" - Registration has finished
	// "running" - The first match has started
	// "finished" - The last match among all stages has been played
	// "complete" - Tournament has been completed
	Status []string `json:"status,omitempty"`
}

type TournamentsList struct {
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
	// Registration end date and time
	RegistrationEndAt *wgnTime.UnixTime `json:"registration_end_at,omitempty"`
	// Registration start date and time
	RegistrationStartAt *wgnTime.UnixTime `json:"registration_start_at,omitempty"`
	// Tournament start date and time
	StartAt *wgnTime.UnixTime `json:"start_at,omitempty"`
	// Tournament status
	Status *string `json:"status,omitempty"`
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
