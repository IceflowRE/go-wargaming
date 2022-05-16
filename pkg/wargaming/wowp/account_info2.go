package wowp

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type AccountInfo2 struct {
	// Date when player details were updated
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Player statistics
	Statistics struct {
		// General player statistics
		All struct {
			// Total participations in capturing sectors
			ZoneCaptures int `json:"zone_captures,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Grades for completing aircraft type-specific missions
			RanksByPlaneType map[string]string `json:"ranks_by_plane_type,omitempty"`
			// Players
			Players struct {
				// Most aerial targets destroyed per battle
				MaxKilled int `json:"max_killed,omitempty"`
				// Most damage caused per battle
				MaxDamageDealt float32 `json:"max_damage_dealt,omitempty"`
				// Number of aerial targets destroyed when defending sectors
				KilledInDefence int `json:"killed_in_defence,omitempty"`
				// Number of destroyed aerial targets
				Killed int `json:"killed,omitempty"`
				// Damage caused
				DamageDealt float32 `json:"damage_dealt,omitempty"`
				// Average number of aerial targets destroyed per battle
				AvgKilled float32 `json:"avg_killed,omitempty"`
				// Number of assists
				Assisted int `json:"assisted,omitempty"`
			} `json:"players,omitempty"`
			// Most combat points earned per battle
			MaxBattlesScore int `json:"max_battles_score,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Ground targets
			GroundObjects struct {
				// Most aerial targets destroyed per battle
				MaxKilled int `json:"max_killed,omitempty"`
				// Most damage caused per battle
				MaxDamageDealt float32 `json:"max_damage_dealt,omitempty"`
				// Number of destroyed aerial targets
				Killed int `json:"killed,omitempty"`
				// Damage caused
				DamageDealt float32 `json:"damage_dealt,omitempty"`
				// Average number of aerial targets destroyed per sortie
				AvgKilledPerFlight float32 `json:"avg_killed_per_flight,omitempty"`
				// Average number of aerial targets destroyed per battle
				AvgKilled float32 `json:"avg_killed,omitempty"`
				// Number of assists
				Assisted int `json:"assisted,omitempty"`
			} `json:"ground_objects,omitempty"`
			// Number of sorties by aircraft types
			FlightsByPlaneType map[string]string `json:"flights_by_plane_type,omitempty"`
			// Number of sorties
			Flights int `json:"flights,omitempty"`
			// Sortie duration
			FlightTime int `json:"flight_time,omitempty"`
			// Battle performance by aircraft types
			EffByPlaneType map[string]string `json:"eff_by_plane_type,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Air defense aircraft and bombers
			DefendersAndBombers struct {
				// Average number of aerial targets destroyed per battle
				AvgKilled float32 `json:"avg_killed,omitempty"`
			} `json:"defenders_and_bombers,omitempty"`
			// Air defense aircraft
			Defenders struct {
				// Most aerial targets destroyed per battle
				MaxKilled int `json:"max_killed,omitempty"`
				// Most damage caused per battle
				MaxDamageDealt float32 `json:"max_damage_dealt,omitempty"`
				// Number of destroyed aerial targets
				Killed int `json:"killed,omitempty"`
				// Damage caused
				DamageDealt float32 `json:"damage_dealt,omitempty"`
				// Number of assists
				Assisted int `json:"assisted,omitempty"`
			} `json:"defenders,omitempty"`
			// Number of times player's aircraft was destroyed
			Deaths int `json:"deaths,omitempty"`
			// Bombers
			Bombers struct {
				// Most aerial targets destroyed per battle
				MaxKilled int `json:"max_killed,omitempty"`
				// Most damage caused per battle
				MaxDamageDealt float32 `json:"max_damage_dealt,omitempty"`
				// Number of destroyed aerial targets
				Killed int `json:"killed,omitempty"`
				// Damage caused
				DamageDealt float32 `json:"damage_dealt,omitempty"`
				// Number of assists
				Assisted int `json:"assisted,omitempty"`
			} `json:"bombers,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Combat points
			BattleScore int `json:"battle_score,omitempty"`
			// Average experience earned per battle
			AvgXp float32 `json:"avg_xp,omitempty"`
			// Average number of combat points earned per battle
			AvgBattleScore float32 `json:"avg_battle_score,omitempty"`
			// Aerial targets
			AirTargets struct {
				// Most damage caused per battle
				MaxDamageDealt float32 `json:"max_damage_dealt,omitempty"`
				// Average number of aerial targets destroyed per sortie
				AvgKilledPerFlight float32 `json:"avg_killed_per_flight,omitempty"`
			} `json:"air_targets,omitempty"`
		} `json:"all,omitempty"`
	} `json:"statistics,omitempty"`
	// Player's private data
	Private struct {
		// Premium Account expiration time
		PremiumExpiresAt wgnTime.UnixTime `json:"premium_expires_at,omitempty"`
		// Indicates if the account is Premium Account
		IsPremium bool `json:"is_premium,omitempty"`
		// Gold
		Gold int `json:"gold,omitempty"`
		// Free Experience
		FreeXp int `json:"free_xp,omitempty"`
		// Credits
		Credits int `json:"credits,omitempty"`
		// Contact groups
		Contacts map[string]string `json:"contacts,omitempty"`
	} `json:"private,omitempty"`
	// Battles fought during Open Beta Test
	ObtGamesPlayed int `json:"obt_games_played,omitempty"`
	// Player name
	Nickname string `json:"nickname,omitempty"`
	// Last battle time
	LastBattleTime wgnTime.UnixTime `json:"last_battle_time,omitempty"`
	// Personal rating
	GlobalRating int `json:"global_rating,omitempty"`
	// Date when player's account was created
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Language selected in the game client
	ClientLanguage string `json:"client_language,omitempty"`
	// Battles fought during Close Beta Test
	CbtGamesPlayed int `json:"cbt_games_played,omitempty"`
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
} 
