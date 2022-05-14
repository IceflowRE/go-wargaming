package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
	"strconv"
)

type WotbTanksStats struct {
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
	// Overall battle life time in seconds
	BattleLifeTime UnixTime `json:"battle_life_time,omitempty"`
	// Time of last update of vehicle availability in the Garage. This data requires a valid access_token for the specified account.
	InGarageUpdated UnixTime `json:"in_garage_updated,omitempty"`
	// Last battle time
	LastBattleTime UnixTime `json:"last_battle_time,omitempty"`
	// Mastery Badges:
	// 
	// 0 — None
	// 1 — 3rd Class 
	// 2 — 2nd Class
	// 3 — 1st Class
	// 4 — Ace Tanker
	MarkOfMastery int `json:"mark_of_mastery,omitempty"`
	// Maximum destroyed in battle
	MaxFrags int `json:"max_frags,omitempty"`
	// Maximum experience per battle
	MaxXp int `json:"max_xp,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Details on vehicles destroyed. This data requires a valid access_token for the specified account.
	Frags map[string]string `json:"frags,omitempty"`
	// Availability of vehicle in the Garage. This data requires a valid access_token for the specified account.
	InGarage bool `json:"in_garage,omitempty"`
	// Overall Statistics
	All struct {
		// Number of battles
		Battles int `json:"battles,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Vehicles destroyed (Tier &gt;= 8)
		Frags8P int `json:"frags8p,omitempty"`
		// Number of hits
		Hits int `json:"hits,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Maximum destroyed in battle
		MaxFrags int `json:"max_frags,omitempty"`
		// Maximum experience per battle
		MaxXp int `json:"max_xp,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Vehicles spotted
		Spotted int `json:"spotted,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		WinAndSurvived int `json:"win_and_survived,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Total experience
		Xp int `json:"xp,omitempty"`
	} `json:"all,omitempty"`
}

// WotbTanksStats Method returns general statistics for each vehicle of each player.
//
// https://developers.wargaming.net/reference/all/wotb/tanks/stats
//
// account_id:
//     Player account ID
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// in_garage:
//     Filter by vehicle availability in the Garage. If the parameter is not specified, all vehicles are returned. Parameter processing requires a valid access_token for the specified account_id. Valid values:
//     
//     "1" &mdash; Return vehicles available in the Garage. 
//     "0" &mdash; Return vehicles that are no longer in the Garage.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "en" &mdash; English 
//     "ru" &mdash; Русский (by default)
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "zh-cn" &mdash; 简体中文 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
// tank_id:
//     Player's vehicle ID. Maximum limit: 100.
func (client *Client) WotbTanksStats(realm Realm, accountId int, accessToken string, fields []string, inGarage string, language string, tankId []int) (*WotbTanksStats, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
		"in_garage": inGarage,
		"language": language,
		"tank_id": utils.SliceIntToString(tankId, ","),
	}

	var result *WotbTanksStats
	err := client.doGetDataRequest(realm, "/wotb/tanks/stats/", reqParam, &result)
	return result, err
}
