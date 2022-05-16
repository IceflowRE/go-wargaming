package wowp

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type PlanesStats struct {
	// Date when details on aircraft were updated
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Private data on the player's aircraft
	Private struct {
		// Last time when the information about the availability of aircraft in the Hangar was updated
		InGarageUpdatedAt wgnTime.UnixTime `json:"in_garage_updated_at,omitempty"`
		// Aircraft availability in the Hangar
		InGarage bool `json:"in_garage,omitempty"`
	} `json:"private,omitempty"`
	// Aircraft ID
	PlaneId int `json:"plane_id,omitempty"`
	// General statistics
	All struct {
		// Total participations in capturing sectors
		ZoneCaptures int `json:"zone_captures,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Players
		Players struct {
			// Most aerial targets destroyed when defending sectors
			MaxKilledInDefence int `json:"max_killed_in_defence,omitempty"`
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
			// Most aerial targets destroyed per sortie
			MaxKilledPerFlight int `json:"max_killed_per_flight,omitempty"`
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
			// Average damage caused per sortie
			AvgDamageDealtPerFlight float32 `json:"avg_damage_dealt_per_flight,omitempty"`
			// Number of assists
			Assisted int `json:"assisted,omitempty"`
		} `json:"ground_objects,omitempty"`
		// Number of sorties
		Flights int `json:"flights,omitempty"`
		// Sortie duration
		FlightTime int `json:"flight_time,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Air defense aircraft and bombers
		DefendersAndBombers struct {
			// Most aerial targets destroyed per battle
			MaxKilled int `json:"max_killed,omitempty"`
			// Most damage caused per battle
			MaxDamageDealt float32 `json:"max_damage_dealt,omitempty"`
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
		// Number of times that an aircraft was selected in the Hangar
		ChosenFirst int `json:"chosen_first,omitempty"`
		// Number of times that an aircraft was selected during a battle
		Chosen int `json:"chosen,omitempty"`
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
		// Average sortie duration
		AvgFlightTime float32 `json:"avg_flight_time,omitempty"`
		// Average number of combat points earned per battle
		AvgBattleScore float32 `json:"avg_battle_score,omitempty"`
		// Aerial targets
		AirTargets struct {
			// Most aerial targets destroyed per sortie
			MaxKilledPerFlight int `json:"max_killed_per_flight,omitempty"`
			// Most damage caused per battle
			MaxDamageDealt float32 `json:"max_damage_dealt,omitempty"`
			// Average number of aerial targets destroyed per sortie
			AvgKilledPerFlight float32 `json:"avg_killed_per_flight,omitempty"`
			// Average damage caused per sortie
			AvgDamageDealtPerFlight float32 `json:"avg_damage_dealt_per_flight,omitempty"`
		} `json:"air_targets,omitempty"`
	} `json:"all,omitempty"`
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
} 
