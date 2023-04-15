// Auto generated file!

package wowp

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type PlanesStatsOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Filter by aircraft availability in the Hangar. If the parameter is not specified, all aircraft are returned. Parameter processing requires a valid access_token for the specified account_id. Valid values:
	//
	// "1" - Return aircraft available in the Hangar.
	// "0" - Return aircraft that are no longer available in the Hangar.
	InGarage *string `json:"in_garage,omitempty"`
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
	// Aircraft ID. Maximum limit: 100.
	PlaneId []int `json:"plane_id,omitempty"`
}

type PlanesStats struct {
	// Player account ID
	AccountId *int `json:"account_id,omitempty"`
	// General statistics
	All *struct {
		// Aerial targets
		AirTargets *struct {
			// Average damage caused per sortie
			AvgDamageDealtPerFlight *float32 `json:"avg_damage_dealt_per_flight,omitempty"`
			// Average number of aerial targets destroyed per sortie
			AvgKilledPerFlight *float32 `json:"avg_killed_per_flight,omitempty"`
			// Most damage caused per battle
			MaxDamageDealt *float32 `json:"max_damage_dealt,omitempty"`
			// Most aerial targets destroyed per sortie
			MaxKilledPerFlight *int `json:"max_killed_per_flight,omitempty"`
		} `json:"air_targets,omitempty"`
		// Average number of combat points earned per battle
		AvgBattleScore *float32 `json:"avg_battle_score,omitempty"`
		// Average sortie duration
		AvgFlightTime *float32 `json:"avg_flight_time,omitempty"`
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
		// Number of times that an aircraft was selected during a battle
		Chosen *int `json:"chosen,omitempty"`
		// Number of times that an aircraft was selected in the Hangar
		ChosenFirst *int `json:"chosen_first,omitempty"`
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
			// Most damage caused per battle
			MaxDamageDealt *float32 `json:"max_damage_dealt,omitempty"`
			// Most aerial targets destroyed per battle
			MaxKilled *int `json:"max_killed,omitempty"`
		} `json:"defenders_and_bombers,omitempty"`
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Sortie duration
		FlightTime *int `json:"flight_time,omitempty"`
		// Number of sorties
		Flights *int `json:"flights,omitempty"`
		// Ground targets
		GroundObjects *struct {
			// Number of assists
			Assisted *int `json:"assisted,omitempty"`
			// Average damage caused per sortie
			AvgDamageDealtPerFlight *float32 `json:"avg_damage_dealt_per_flight,omitempty"`
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
			// Most aerial targets destroyed per sortie
			MaxKilledPerFlight *int `json:"max_killed_per_flight,omitempty"`
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
			// Most aerial targets destroyed when defending sectors
			MaxKilledInDefence *int `json:"max_killed_in_defence,omitempty"`
		} `json:"players,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total participations in capturing sectors
		ZoneCaptures *int `json:"zone_captures,omitempty"`
	} `json:"all,omitempty"`
	// Aircraft ID
	PlaneId *int `json:"plane_id,omitempty"`
	// Private data on the player's aircraft
	Private *struct {
		// Aircraft availability in the Hangar
		InGarage *bool `json:"in_garage,omitempty"`
		// Last time when the information about the availability of aircraft in the Hangar was updated
		InGarageUpdatedAt *wgnTime.UnixTime `json:"in_garage_updated_at,omitempty"`
	} `json:"private,omitempty"`
	// Date when details on aircraft were updated
	UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
}
