package wotx

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type TanksStats struct {
	// Trees knocked down
	TreesCut int `json:"trees_cut,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Maximum experience per battle
	MaxXp int `json:"max_xp,omitempty"`
	// Maximum destroyed in battle
	MaxFrags int `json:"max_frags,omitempty"`
	// Mastery Badges:
	// 
	// 0 — None
	// 1 — 3rd Class 
	// 2 — 2nd Class
	// 3 — 1st Class
	// 4 — Ace Tanker
	MarkOfMastery int `json:"mark_of_mastery,omitempty"`
	// Last battle time
	LastBattleTime wgnTime.UnixTime `json:"last_battle_time,omitempty"`
	// Time of status update. This data requires a valid access_token for the specified account.
	InGarageUpdated wgnTime.UnixTime `json:"in_garage_updated,omitempty"`
	// Availability of vehicle in the Garage. This data requires a valid access_token for the specified account.
	InGarage bool `json:"in_garage,omitempty"`
	// Details on vehicles destroyed. This data requires a valid access_token for the specified account.
	Frags map[string]string `json:"frags,omitempty"`
	// Tank Company battles statistics
	Company struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Penetrations received. Value is calculated starting from version 1.10.
		PiercingsReceived int `json:"piercings_received,omitempty"`
		// Penetrations. Value is calculated starting from version 1.10.
		Piercings int `json:"piercings,omitempty"`
		// Direct hits received that caused no damage. Value is calculated starting from version 1.10.
		NoDamageDirectHitsReceived int `json:"no_damage_direct_hits_received,omitempty"`
		// Maximum damage caused in a battle.
		MaxDamage int `json:"max_damage,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Hits received as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHitsReceived int `json:"explosion_hits_received,omitempty"`
		// Hits on enemy as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHits int `json:"explosion_hits,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Direct hits received. Value is calculated starting from version 1.10.
		DirectHitsReceived int `json:"direct_hits_received,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Average damage upon your destroying a track. Value is calculated starting from version 1.3.
		DamageAssistedTrack int `json:"damage_assisted_track,omitempty"`
		// Average damage upon your spotting. Value is calculated starting from version 1.3.
		DamageAssistedRadio int `json:"damage_assisted_radio,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
	} `json:"company,omitempty"`
	// Total time of destruction
	BattleLifeTime wgnTime.UnixTime `json:"battle_life_time,omitempty"`
	// Overall Statistics
	All struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Penetrations received. Value is calculated starting from version 1.10.
		PiercingsReceived int `json:"piercings_received,omitempty"`
		// Penetrations. Value is calculated starting from version 1.10.
		Piercings int `json:"piercings,omitempty"`
		// Direct hits received that caused no damage. Value is calculated starting from version 1.10.
		NoDamageDirectHitsReceived int `json:"no_damage_direct_hits_received,omitempty"`
		// Maximum damage caused in a battle.
		MaxDamage int `json:"max_damage,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Hits received as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHitsReceived int `json:"explosion_hits_received,omitempty"`
		// Hits on enemy as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHits int `json:"explosion_hits,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Direct hits received. Value is calculated starting from version 1.10.
		DirectHitsReceived int `json:"direct_hits_received,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Average damage upon your destroying a track. Value is calculated starting from version 1.3.
		DamageAssistedTrack int `json:"damage_assisted_track,omitempty"`
		// Average damage upon your spotting. Value is calculated starting from version 1.3.
		DamageAssistedRadio int `json:"damage_assisted_radio,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
	} `json:"all,omitempty"`
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
} 
