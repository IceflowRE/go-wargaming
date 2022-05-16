package wot

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type GlobalmapClaninfo struct {
	// Clan tag
	Tag string `json:"tag,omitempty"`
	// Clan statistics on the Global Map
	Statistics struct {
		// Victories in Champion division
		Wins8Level int `json:"wins_8_level,omitempty"`
		// Victories in Medium division
		Wins6Level int `json:"wins_6_level,omitempty"`
		// Victories in Absolute division
		Wins10Level int `json:"wins_10_level,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Current number of captured provinces
		ProvincesCount int `json:"provinces_count,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Total number by provinces captured by clan
		Captures int `json:"captures,omitempty"`
		// Battles fought in Champion division
		Battles8Level int `json:"battles_8_level,omitempty"`
		// Battles fought in Medium division
		Battles6Level int `json:"battles_6_level,omitempty"`
		// Battles fought in Absolute division
		Battles10Level int `json:"battles_10_level,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
	} `json:"statistics,omitempty"`
	// Clan rating on the Global Map
	Ratings struct {
		// Ratings update time
		UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
		// Clan Elo rating in Champion division
		Elo8 int `json:"elo_8,omitempty"`
		// Clan Elo rating in Medium division
		Elo6 int `json:"elo_6,omitempty"`
		// Clan Elo rating in Absolute division
		Elo10 int `json:"elo_10,omitempty"`
	} `json:"ratings,omitempty"`
	// Restricted clan information on the Global Map
	Private struct {
		// Influence points
		Influence int `json:"influence,omitempty"`
		// Influence points spent per day
		DailyWage int `json:"daily_wage,omitempty"`
	} `json:"private,omitempty"`
	// Clan name
	Name string `json:"name,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
} 
