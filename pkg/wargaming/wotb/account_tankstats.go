package wotb

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type AccountTankstats struct {
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
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
	// Overall battle life time in seconds
	BattleLifeTime wgnTime.UnixTime `json:"battle_life_time,omitempty"`
	// Overall Statistics
	All struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Victories in battles survived
		WinAndSurvived int `json:"win_and_survived,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Vehicles spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Maximum experience per battle
		MaxXp int `json:"max_xp,omitempty"`
		// Maximum destroyed in battle
		MaxFrags int `json:"max_frags,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Number of hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed (Tier &gt;= 8)
		Frags8P int `json:"frags8p,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles
		Battles int `json:"battles,omitempty"`
	} `json:"all,omitempty"`
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
} 
