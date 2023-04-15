// Auto generated file!

package wotb

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type TanksStatsOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Filter by vehicle availability in the Garage. If the parameter is not specified, all vehicles are returned. Parameter processing requires a valid access_token for the specified account_id. Valid values:
	//
	// "1" - Return vehicles available in the Garage.
	// "0" - Return vehicles that are no longer in the Garage.
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
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
	// Player's vehicle ID. Maximum limit: 100.
	TankId []int `json:"tank_id,omitempty"`
}

type TanksStats struct {
	// Player account ID
	AccountId *int `json:"account_id,omitempty"`
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
		// Maximum experience per battle
		MaxXp *int `json:"max_xp,omitempty"`
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
	// Overall battle life time in seconds
	BattleLifeTime *wgnTime.UnixTime `json:"battle_life_time,omitempty"`
	// Details on vehicles destroyed. This data requires a valid access_token for the specified account.
	Frags map[string]string `json:"frags,omitempty"`
	// Availability of vehicle in the Garage. This data requires a valid access_token for the specified account.
	InGarage *bool `json:"in_garage,omitempty"`
	// Time of last update of vehicle availability in the Garage. This data requires a valid access_token for the specified account.
	InGarageUpdated *wgnTime.UnixTime `json:"in_garage_updated,omitempty"`
	// Last battle time
	LastBattleTime *wgnTime.UnixTime `json:"last_battle_time,omitempty"`
	// Mastery Badges:
	//
	// 0 - None
	// 1 - 3rd Class
	// 2 - 2nd Class
	// 3 - 1st Class
	// 4 - Ace Tanker
	MarkOfMastery *int `json:"mark_of_mastery,omitempty"`
	// Maximum destroyed in battle
	MaxFrags *int `json:"max_frags,omitempty"`
	// Maximum experience per battle
	MaxXp *int `json:"max_xp,omitempty"`
	// Vehicle ID
	TankId *int `json:"tank_id,omitempty"`
}
