package wotb

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type TournamentsInfo struct {
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
	// Total number of teams in tournament, both already registered and in process of forming
	Teams struct {
		// Total number of teams in tournament, both already registered and in process of forming
		Total int `json:"total,omitempty"`
		// Minimum number of teams in tournament
		Min int `json:"min,omitempty"`
		// Maximum number of teams available in tournament
		Max int `json:"max,omitempty"`
		// Number of confirmed teams in tournament
		Confirmed int `json:"confirmed,omitempty"`
	} `json:"teams,omitempty"`
	// Tournament status
	Status string `json:"status,omitempty"`
	// Tournament start date and time
	StartAt wgnTime.UnixTime `json:"start_at,omitempty"`
	// Tournament rules and regulations
	Rules string `json:"rules,omitempty"`
	// Registration start date and time
	RegistrationStartAt wgnTime.UnixTime `json:"registration_start_at,omitempty"`
	// Registration end date and time
	RegistrationEndAt wgnTime.UnixTime `json:"registration_end_at,omitempty"`
	// Tournament award description
	PrizeDescription string `json:"prize_description,omitempty"`
	// Tournament other regulations
	OtherRules string `json:"other_rules,omitempty"`
	// Minimum number of players per team in tournament
	MinPlayersCount int `json:"min_players_count,omitempty"`
	// Section contains information about links to tournament media resources
	MediaLinks struct {
		// Link to media resource
		Url string `json:"url,omitempty"`
		// Type of media resource link
		Kind string `json:"kind,omitempty"`
		// Image of media resource link; available only for links of "Custom" type
		Image string `json:"image,omitempty"`
		// ID of media resource link
		Id string `json:"id,omitempty"`
	} `json:"media_links,omitempty"`
	// Maximum number of players per team in tournament
	MaxPlayersCount int `json:"max_players_count,omitempty"`
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
