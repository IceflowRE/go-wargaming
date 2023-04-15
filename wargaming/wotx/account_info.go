// Auto generated file!

package wotx

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type AccountInfoOptions struct {
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
	// "tr" - Türkçe
	Language *string `json:"language,omitempty"`
}

type AccountInfo struct {
	// Player account ID
	AccountId *int `json:"account_id,omitempty"`
	// Date when player's account was created
	CreatedAt *wgnTime.UnixTime `json:"created_at,omitempty"`
	// Personal rating
	GlobalRating *int `json:"global_rating,omitempty"`
	// Last battle time
	LastBattleTime *wgnTime.UnixTime `json:"last_battle_time,omitempty"`
	// Player name
	Nickname *string `json:"nickname,omitempty"`
	// Player's private data
	Private *struct {
		// Account ban details
		BanInfo *string `json:"ban_info,omitempty"`
		// End time of account ban
		BanTime *wgnTime.UnixTime `json:"ban_time,omitempty"`
		// Total time of destruction
		BattleLifeTime *wgnTime.UnixTime `json:"battle_life_time,omitempty"`
		// Credits
		Credits *int `json:"credits,omitempty"`
		// Number of slots available in the Garage
		EmptySlots *int `json:"empty_slots,omitempty"`
		// Free Experience
		FreeXp *int `json:"free_xp,omitempty"`
		// Gold
		Gold *int `json:"gold,omitempty"`
		// Indicates if the account is Premium Account
		IsPremium *bool `json:"is_premium,omitempty"`
		// Premium Account expiration time
		PremiumExpiresAt *wgnTime.UnixTime `json:"premium_expires_at,omitempty"`
		// Account restrictions
		Restrictions *struct {
			// End time of chat ban
			ChatBanTime *wgnTime.UnixTime `json:"chat_ban_time,omitempty"`
		} `json:"restrictions,omitempty"`
		// Number of slots in the Garage
		Slots *int `json:"slots,omitempty"`
	} `json:"private,omitempty"`
	// Player statistics
	Statistics *struct {
		// Overall Statistics
		All *struct {
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"all,omitempty"`
		// Tank Company battles statistics
		Company *struct {
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"company,omitempty"`
		// Average damage upon your spotting. Value is calculated starting from version 1.3.
		DamageAssistedRadio *int `json:"damage_assisted_radio,omitempty"`
		// Average damage upon your destroying a track. Value is calculated starting from version 1.3.
		DamageAssistedTrack *int `json:"damage_assisted_track,omitempty"`
		// Average damage upon your shooting the wheel.
		DamageAssistedWheel *int `json:"damage_assisted_wheel,omitempty"`
		// Direct hits received. Value is calculated starting from version 1.10.
		DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
		// Hits on enemy as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHits *int `json:"explosion_hits,omitempty"`
		// Hits received as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
		// Number and models of vehicles destroyed by a player. Player's private data.
		Frags map[string]string `json:"frags,omitempty"`
		// Maximum damage caused in a battle
		MaxDamage *int `json:"max_damage,omitempty"`
		// Vehicle used to cause maximum damage
		MaxDamageTankId *int `json:"max_damage_tank_id,omitempty"`
		// Maximum destroyed in battle
		MaxFrags *int `json:"max_frags,omitempty"`
		// Vehicle, in which maximum number of enemy vehicles was destroyed
		MaxFragsTankId *int `json:"max_frags_tank_id,omitempty"`
		// Maximum experience per battle
		MaxXp *int `json:"max_xp,omitempty"`
		// Vehicle, in which maximum experience per battle was earned
		MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
		// Direct hits received that caused no damage. Value is calculated starting from version 1.10.
		NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
		// Penetrations. Value is calculated starting from version 1.10.
		Piercings *int `json:"piercings,omitempty"`
		// Penetrations received. Value is calculated starting from version 1.10.
		PiercingsReceived *int `json:"piercings_received,omitempty"`
		// Trees knocked down
		TreesCut *int `json:"trees_cut,omitempty"`
	} `json:"statistics,omitempty"`
	// Date when player details were updated
	UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
}
