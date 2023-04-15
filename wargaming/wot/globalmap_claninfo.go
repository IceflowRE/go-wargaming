// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type GlobalmapClaninfoOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
}

type GlobalmapClaninfo struct {
	// Clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Clan name
	Name *string `json:"name,omitempty"`
	// Restricted clan information on the Global Map
	Private *struct {
		// Influence points spent per day
		DailyWage *int `json:"daily_wage,omitempty"`
		// Influence points
		Influence *int `json:"influence,omitempty"`
	} `json:"private,omitempty"`
	// Clan rating on the Global Map
	Ratings *struct {
		// Clan Elo rating in Absolute division
		Elo10 *int `json:"elo_10,omitempty"`
		// Clan Elo rating in Medium division
		Elo6 *int `json:"elo_6,omitempty"`
		// Clan Elo rating in Champion division
		Elo8 *int `json:"elo_8,omitempty"`
		// Ratings update time
		UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
	} `json:"ratings,omitempty"`
	// Clan statistics on the Global Map
	Statistics *struct {
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Battles fought in Absolute division
		Battles10Level *int `json:"battles_10_level,omitempty"`
		// Battles fought in Medium division
		Battles6Level *int `json:"battles_6_level,omitempty"`
		// Battles fought in Champion division
		Battles8Level *int `json:"battles_8_level,omitempty"`
		// Total number by provinces captured by clan
		Captures *int `json:"captures,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Current number of captured provinces
		ProvincesCount *int `json:"provinces_count,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Victories in Absolute division
		Wins10Level *int `json:"wins_10_level,omitempty"`
		// Victories in Medium division
		Wins6Level *int `json:"wins_6_level,omitempty"`
		// Victories in Champion division
		Wins8Level *int `json:"wins_8_level,omitempty"`
	} `json:"statistics,omitempty"`
	// Clan tag
	Tag *string `json:"tag,omitempty"`
}
