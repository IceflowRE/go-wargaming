package wotx

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type AccountInfo struct {
	// Date when player details were updated
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Player statistics
	Statistics struct {
		// Trees knocked down
		TreesCut int `json:"trees_cut,omitempty"`
		// Penetrations received. Value is calculated starting from version 1.10.
		PiercingsReceived int `json:"piercings_received,omitempty"`
		// Penetrations. Value is calculated starting from version 1.10.
		Piercings int `json:"piercings,omitempty"`
		// Direct hits received that caused no damage. Value is calculated starting from version 1.10.
		NoDamageDirectHitsReceived int `json:"no_damage_direct_hits_received,omitempty"`
		// Vehicle, in which maximum experience per battle was earned
		MaxXpTankId int `json:"max_xp_tank_id,omitempty"`
		// Maximum experience per battle
		MaxXp int `json:"max_xp,omitempty"`
		// Vehicle, in which maximum number of enemy vehicles was destroyed
		MaxFragsTankId int `json:"max_frags_tank_id,omitempty"`
		// Maximum destroyed in battle
		MaxFrags int `json:"max_frags,omitempty"`
		// Vehicle used to cause maximum damage
		MaxDamageTankId int `json:"max_damage_tank_id,omitempty"`
		// Maximum damage caused in a battle
		MaxDamage int `json:"max_damage,omitempty"`
		// Number and models of vehicles destroyed by a player. Player's private data.
		Frags map[string]string `json:"frags,omitempty"`
		// Hits received as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHitsReceived int `json:"explosion_hits_received,omitempty"`
		// Hits on enemy as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHits int `json:"explosion_hits,omitempty"`
		// Direct hits received. Value is calculated starting from version 1.10.
		DirectHitsReceived int `json:"direct_hits_received,omitempty"`
		// Average damage upon your destroying a track. Value is calculated starting from version 1.3.
		DamageAssistedTrack int `json:"damage_assisted_track,omitempty"`
		// Average damage upon your spotting. Value is calculated starting from version 1.3.
		DamageAssistedRadio int `json:"damage_assisted_radio,omitempty"`
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
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
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
			// Battles fought
			Battles int `json:"battles,omitempty"`
		} `json:"company,omitempty"`
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
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
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
			// Battles fought
			Battles int `json:"battles,omitempty"`
		} `json:"all,omitempty"`
	} `json:"statistics,omitempty"`
	// Player's private data
	Private struct {
		// Number of slots in the Garage
		Slots int `json:"slots,omitempty"`
		// Account restrictions
		Restrictions struct {
			// End time of chat ban
			ChatBanTime wgnTime.UnixTime `json:"chat_ban_time,omitempty"`
		} `json:"restrictions,omitempty"`
		// Premium Account expiration time
		PremiumExpiresAt wgnTime.UnixTime `json:"premium_expires_at,omitempty"`
		// Indicates if the account is Premium Account
		IsPremium bool `json:"is_premium,omitempty"`
		// Gold
		Gold int `json:"gold,omitempty"`
		// Free Experience
		FreeXp int `json:"free_xp,omitempty"`
		// Number of slots available in the Garage
		EmptySlots int `json:"empty_slots,omitempty"`
		// Credits
		Credits int `json:"credits,omitempty"`
		// Total time of destruction
		BattleLifeTime wgnTime.UnixTime `json:"battle_life_time,omitempty"`
		// End time of account ban
		BanTime wgnTime.UnixTime `json:"ban_time,omitempty"`
		// Account ban details
		BanInfo string `json:"ban_info,omitempty"`
	} `json:"private,omitempty"`
	// Player name
	Nickname string `json:"nickname,omitempty"`
	// Last battle time
	LastBattleTime wgnTime.UnixTime `json:"last_battle_time,omitempty"`
	// Personal rating
	GlobalRating int `json:"global_rating,omitempty"`
	// Date when player's account was created
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
} 
