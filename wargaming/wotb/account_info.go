// Auto generated file!

package wotb

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type AccountInfoOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
	// Extra fields that will be added to the response. Valid values:
	//
	// "private.grouped_contacts"
	// "statistics.rating"
	Extra []string `json:"extra,omitempty"`
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
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
}

type AccountInfo struct {
	// Player account ID
	AccountId *int `json:"account_id,omitempty"`
	// Date when player's account was created
	CreatedAt *wgnTime.UnixTime `json:"created_at,omitempty"`
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
		// Overall battle life time in seconds
		BattleLifeTime *int `json:"battle_life_time,omitempty"`
		// Credits
		Credits *int `json:"credits,omitempty"`
		// Free Experience
		FreeXp *int `json:"free_xp,omitempty"`
		// Gold
		Gold *int `json:"gold,omitempty"`
		// Contact groups.
		// An extra field.
		GroupedContacts *struct {
			// Blocked
			Blocked []int `json:"blocked,omitempty"`
			// Groups
			Groups map[string]string `json:"groups,omitempty"`
			// Not grouped
			Ungrouped []int `json:"ungrouped,omitempty"`
		} `json:"grouped_contacts,omitempty"`
		// Indicates if the account is Premium Account
		IsPremium *bool `json:"is_premium,omitempty"`
		// Premium Account expiration time
		PremiumExpiresAt *wgnTime.UnixTime `json:"premium_expires_at,omitempty"`
		// Account restrictions
		Restrictions *struct {
			// End time of chat ban
			ChatBanTime *wgnTime.UnixTime `json:"chat_ban_time,omitempty"`
		} `json:"restrictions,omitempty"`
	} `json:"private,omitempty"`
	// Player statistics
	Statistics *struct {
		// Overall Statistics
		All *struct {
			// Number of battles
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
			// Vehicles destroyed (Tier >= 8)
			Frags8p *int `json:"frags8p,omitempty"`
			// Number of hits
			Hits *int `json:"hits,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
			// Maximum destroyed in battle
			MaxFrags *int `json:"max_frags,omitempty"`
			// Vehicle with maximum number of enemy vehicles destroyed
			MaxFragsTankId *int `json:"max_frags_tank_id,omitempty"`
			// Maximum experience per battle
			MaxXp *int `json:"max_xp,omitempty"`
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Vehicles spotted
			Spotted *int `json:"spotted,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			WinAndSurvived *int `json:"win_and_survived,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"all,omitempty"`
		// Clan battle statistics
		Clan *struct {
			// Number of battles
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
			// Vehicles destroyed (Tier >= 8)
			Frags8p *int `json:"frags8p,omitempty"`
			// Number of hits
			Hits *int `json:"hits,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
			// Maximum destroyed in battle
			MaxFrags *int `json:"max_frags,omitempty"`
			// Vehicle with maximum number of enemy vehicles destroyed
			MaxFragsTankId *int `json:"max_frags_tank_id,omitempty"`
			// Maximum experience per battle
			MaxXp *int `json:"max_xp,omitempty"`
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Vehicles spotted
			Spotted *int `json:"spotted,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			WinAndSurvived *int `json:"win_and_survived,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"clan,omitempty"`
		// Number and models of vehicles destroyed by a player. Player's private data.
		Frags map[string]string `json:"frags,omitempty"`
		// Rating battles statistics.
		// An extra field.
		Rating *struct {
			// Number of battles
			Battles *int `json:"battles,omitempty"`
			// Battles before end of calibration
			CalibrationBattlesLeft *int `json:"calibration_battles_left,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Number of current season for player
			CurrentSeason *int `json:"current_season,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Vehicles destroyed (Tier >= 8)
			Frags8p *int `json:"frags8p,omitempty"`
			// Number of hits
			Hits *int `json:"hits,omitempty"`
			// Flag of recalibration start
			IsRecalibration *bool `json:"is_recalibration,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
			// Matchmaking rating
			MmRating *float32 `json:"mm_rating,omitempty"`
			// Recalibration start time
			RecalibrationStartTime *wgnTime.UnixTime `json:"recalibration_start_time,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Vehicles spotted
			Spotted *int `json:"spotted,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			WinAndSurvived *int `json:"win_and_survived,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"rating,omitempty"`
	} `json:"statistics,omitempty"`
	// Date when player details were updated
	UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
}
