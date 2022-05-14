package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
	"strconv"
)

type WotbAccountTankstats struct {
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
	// Overall battle life time in seconds
	BattleLifeTime UnixTime `json:"battle_life_time,omitempty"`
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
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
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

// WotbAccountTankstats Method returns players' statistics on the vehicle.
//
// https://developers.wargaming.net/reference/all/wotb/account/tankstats
//
// account_id:
//     Player account ID. Maximum limit: 100.
// tank_id:
//     Player's vehicle ID
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
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
func (client *Client) WotbAccountTankstats(realm Realm, accountId []int, tankId int, fields []string, language string) (*WotbAccountTankstats, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"tank_id": strconv.Itoa(tankId),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotbAccountTankstats
	err := client.doGetDataRequest(realm, "/wotb/account/tankstats/", reqParam, &result)
	return result, err
}
