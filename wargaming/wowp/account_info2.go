// Auto generated file!

package wowp

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type AccountInfo2Options struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "ru" - Русский
	// "pl" - Polski
	// "de" - Deutsch
	// "fr" - Français
	// "es" - Español
	// "zh-cn" - 简体中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
}

type AccountInfo2 struct {
	// Player account ID
	AccountId *int `json:"account_id,omitempty"`
	// Battles fought during Close Beta Test
	CbtGamesPlayed *int `json:"cbt_games_played,omitempty"`
	// Language selected in the game client
	ClientLanguage *string `json:"client_language,omitempty"`
	// Date when player's account was created
	CreatedAt *wgnTime.UnixTime `json:"created_at,omitempty"`
	// Personal rating
	GlobalRating *int `json:"global_rating,omitempty"`
	// Last battle time
	LastBattleTime *wgnTime.UnixTime `json:"last_battle_time,omitempty"`
	// Player name
	Nickname *string `json:"nickname,omitempty"`
	// Battles fought during Open Beta Test
	ObtGamesPlayed *int `json:"obt_games_played,omitempty"`
	// Player's private data
	Private *struct {
		// Contact groups
		Contacts map[string]string `json:"contacts,omitempty"`
		// Credits
		Credits *int `json:"credits,omitempty"`
		// Free Experience
		FreeXp *int `json:"free_xp,omitempty"`
		// Gold
		Gold *int `json:"gold,omitempty"`
		// Indicates if the account is Premium Account
		IsPremium *bool `json:"is_premium,omitempty"`
		// Premium Account expiration time
		PremiumExpiresAt *wgnTime.UnixTime `json:"premium_expires_at,omitempty"`
	} `json:"private,omitempty"`
	// Player statistics
	Statistics *struct {
		// General player statistics
		All *struct {
			// Aerial targets
			AirTargets *struct {
				// Average number of aerial targets destroyed per sortie
				AvgKilledPerFlight *float32 `json:"avg_killed_per_flight,omitempty"`
				// Most damage caused per battle
				MaxDamageDealt *float32 `json:"max_damage_dealt,omitempty"`
			} `json:"air_targets,omitempty"`
			// Average number of combat points earned per battle
			AvgBattleScore *float32 `json:"avg_battle_score,omitempty"`
			// Average experience earned per battle
			AvgXp *float32 `json:"avg_xp,omitempty"`
			// Combat points
			BattleScore *int `json:"battle_score,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Bombers
			Bombers *struct {
				// Number of assists
				Assisted *int `json:"assisted,omitempty"`
				// Damage caused
				DamageDealt *float32 `json:"damage_dealt,omitempty"`
				// Number of destroyed aerial targets
				Killed *int `json:"killed,omitempty"`
				// Most damage caused per battle
				MaxDamageDealt *float32 `json:"max_damage_dealt,omitempty"`
				// Most aerial targets destroyed per battle
				MaxKilled *int `json:"max_killed,omitempty"`
			} `json:"bombers,omitempty"`
			// Number of times player's aircraft was destroyed
			Deaths *int `json:"deaths,omitempty"`
			// Air defense aircraft
			Defenders *struct {
				// Number of assists
				Assisted *int `json:"assisted,omitempty"`
				// Damage caused
				DamageDealt *float32 `json:"damage_dealt,omitempty"`
				// Number of destroyed aerial targets
				Killed *int `json:"killed,omitempty"`
				// Most damage caused per battle
				MaxDamageDealt *float32 `json:"max_damage_dealt,omitempty"`
				// Most aerial targets destroyed per battle
				MaxKilled *int `json:"max_killed,omitempty"`
			} `json:"defenders,omitempty"`
			// Air defense aircraft and bombers
			DefendersAndBombers *struct {
				// Average number of aerial targets destroyed per battle
				AvgKilled *float32 `json:"avg_killed,omitempty"`
			} `json:"defenders_and_bombers,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Battle performance by aircraft types
			EffByPlaneType map[string]string `json:"eff_by_plane_type,omitempty"`
			// Sortie duration
			FlightTime *int `json:"flight_time,omitempty"`
			// Number of sorties
			Flights *int `json:"flights,omitempty"`
			// Number of sorties by aircraft types
			FlightsByPlaneType map[string]string `json:"flights_by_plane_type,omitempty"`
			// Ground targets
			GroundObjects *struct {
				// Number of assists
				Assisted *int `json:"assisted,omitempty"`
				// Average number of aerial targets destroyed per battle
				AvgKilled *float32 `json:"avg_killed,omitempty"`
				// Average number of aerial targets destroyed per sortie
				AvgKilledPerFlight *float32 `json:"avg_killed_per_flight,omitempty"`
				// Damage caused
				DamageDealt *float32 `json:"damage_dealt,omitempty"`
				// Number of destroyed aerial targets
				Killed *int `json:"killed,omitempty"`
				// Most damage caused per battle
				MaxDamageDealt *float32 `json:"max_damage_dealt,omitempty"`
				// Most aerial targets destroyed per battle
				MaxKilled *int `json:"max_killed,omitempty"`
			} `json:"ground_objects,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
			// Most combat points earned per battle
			MaxBattlesScore *int `json:"max_battles_score,omitempty"`
			// Players
			Players *struct {
				// Number of assists
				Assisted *int `json:"assisted,omitempty"`
				// Average number of aerial targets destroyed per battle
				AvgKilled *float32 `json:"avg_killed,omitempty"`
				// Damage caused
				DamageDealt *float32 `json:"damage_dealt,omitempty"`
				// Number of destroyed aerial targets
				Killed *int `json:"killed,omitempty"`
				// Number of aerial targets destroyed when defending sectors
				KilledInDefence *int `json:"killed_in_defence,omitempty"`
				// Most damage caused per battle
				MaxDamageDealt *float32 `json:"max_damage_dealt,omitempty"`
				// Most aerial targets destroyed per battle
				MaxKilled *int `json:"max_killed,omitempty"`
			} `json:"players,omitempty"`
			// Grades for completing aircraft type-specific missions
			RanksByPlaneType map[string]string `json:"ranks_by_plane_type,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total participations in capturing sectors
			ZoneCaptures *int `json:"zone_captures,omitempty"`
		} `json:"all,omitempty"`
	} `json:"statistics,omitempty"`
	// Date when player details were updated
	UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
}
