package wotb

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type TournamentsList struct {
	// Award for winning tournament
	WinnerAward struct {
		// Winner Award currency: Free XP, gold or credits
		Currency string `json:"currency,omitempty"`
		// Winner Award amount
		Amount int `json:"amount,omitempty"`
	} `json:"winner_award,omitempty"`
	// Tournament id
	TournamentId int `json:"tournament_id,omitempty"`
	// Tournament name
	Title string `json:"title,omitempty"`
	// Tournament status
	Status string `json:"status,omitempty"`
	// Tournament start date and time
	StartAt wgnTime.UnixTime `json:"start_at,omitempty"`
	// Registration start date and time
	RegistrationStartAt wgnTime.UnixTime `json:"registration_start_at,omitempty"`
	// Registration end date and time
	RegistrationEndAt wgnTime.UnixTime `json:"registration_end_at,omitempty"`
	// Matches start date and time
	MatchesStartAt wgnTime.UnixTime `json:"matches_start_at,omitempty"`
	// Tournament Logo
	Logo struct {
		// Link to preview
		Preview string `json:"preview,omitempty"`
		// Link to logo
		Original string `json:"original,omitempty"`
	} `json:"logo,omitempty"`
	// Fee for participating in tournament
	Fee struct {
		// Fee currency: Free XP, gold or credits
		Currency string `json:"currency,omitempty"`
		// Fee amount
		Amount int `json:"amount,omitempty"`
	} `json:"fee,omitempty"`
	// Tournament end date and time
	EndAt wgnTime.UnixTime `json:"end_at,omitempty"`
	// Tournament description
	Description string `json:"description,omitempty"`
	// Award for participating in tournament
	Award struct {
		// Award currency: Free XP, gold or credits
		Currency string `json:"currency,omitempty"`
		// Award amount
		Amount int `json:"amount,omitempty"`
	} `json:"award,omitempty"`
} 
