package wotb

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type AccountInfo struct {
	// Date when player details were updated
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Player statistics
	Statistics struct {
		// Rating battles statistics.
		// An extra field.
		Rating struct {
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
			// Recalibration start time
			RecalibrationStartTime wgnTime.UnixTime `json:"recalibration_start_time,omitempty"`
			// Matchmaking rating
			MmRating float32 `json:"mm_rating,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Flag of recalibration start
			IsRecalibration bool `json:"is_recalibration,omitempty"`
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
			// Number of current season for player
			CurrentSeason int `json:"current_season,omitempty"`
			// Base capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Battles before end of calibration
			CalibrationBattlesLeft int `json:"calibration_battles_left,omitempty"`
			// Number of battles
			Battles int `json:"battles,omitempty"`
		} `json:"rating,omitempty"`
		// Number and models of vehicles destroyed by a player. Player's private data.
		Frags map[string]string `json:"frags,omitempty"`
		// Clan battle statistics
		Clan struct {
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId int `json:"max_xp_tank_id,omitempty"`
			// Maximum experience per battle
			MaxXp int `json:"max_xp,omitempty"`
			// Vehicle with maximum number of enemy vehicles destroyed
			MaxFragsTankId int `json:"max_frags_tank_id,omitempty"`
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
		} `json:"clan,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId int `json:"max_xp_tank_id,omitempty"`
			// Maximum experience per battle
			MaxXp int `json:"max_xp,omitempty"`
			// Vehicle with maximum number of enemy vehicles destroyed
			MaxFragsTankId int `json:"max_frags_tank_id,omitempty"`
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
	} `json:"statistics,omitempty"`
	// Player's private data
	Private struct {
		// Account restrictions
		Restrictions struct {
			// End time of chat ban
			ChatBanTime wgnTime.UnixTime `json:"chat_ban_time,omitempty"`
		} `json:"restrictions,omitempty"`
		// Premium Account expiration time
		PremiumExpiresAt wgnTime.UnixTime `json:"premium_expires_at,omitempty"`
		// Indicates if the account is Premium Account
		IsPremium bool `json:"is_premium,omitempty"`
		// Contact groups.
		// An extra field.
		GroupedContacts struct {
			// Not grouped
			Ungrouped []int `json:"ungrouped,omitempty"`
			// Groups
			Groups map[string]string `json:"groups,omitempty"`
			// Blocked
			Blocked []int `json:"blocked,omitempty"`
		} `json:"grouped_contacts,omitempty"`
		// Gold
		Gold int `json:"gold,omitempty"`
		// Free Experience
		FreeXp int `json:"free_xp,omitempty"`
		// Credits
		Credits int `json:"credits,omitempty"`
		// Overall battle life time in seconds
		BattleLifeTime int `json:"battle_life_time,omitempty"`
		// End time of account ban
		BanTime wgnTime.UnixTime `json:"ban_time,omitempty"`
		// Account ban details
		BanInfo string `json:"ban_info,omitempty"`
	} `json:"private,omitempty"`
	// Player name
	Nickname string `json:"nickname,omitempty"`
	// Last battle time
	LastBattleTime wgnTime.UnixTime `json:"last_battle_time,omitempty"`
	// Date when player's account was created
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
} 
