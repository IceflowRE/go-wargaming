package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
	"strconv"
)

type WotxTanksStats struct {
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
	// Total time of destruction
	BattleLifeTime UnixTime `json:"battle_life_time,omitempty"`
	// Time of status update. This data requires a valid access_token for the specified account.
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
	// Trees knocked down
	TreesCut int `json:"trees_cut,omitempty"`
	// Details on vehicles destroyed. This data requires a valid access_token for the specified account.
	Frags map[string]string `json:"frags,omitempty"`
	// Availability of vehicle in the Garage. This data requires a valid access_token for the specified account.
	InGarage bool `json:"in_garage,omitempty"`
	// Overall Statistics
	All struct {
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Average damage upon your spotting. Value is calculated starting from version 1.3.
		DamageAssistedRadio int `json:"damage_assisted_radio,omitempty"`
		// Average damage upon your destroying a track. Value is calculated starting from version 1.3.
		DamageAssistedTrack int `json:"damage_assisted_track,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Direct hits received. Value is calculated starting from version 1.10.
		DirectHitsReceived int `json:"direct_hits_received,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Hits on enemy as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHits int `json:"explosion_hits,omitempty"`
		// Hits received as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHitsReceived int `json:"explosion_hits_received,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Maximum damage caused in a battle.
		MaxDamage int `json:"max_damage,omitempty"`
		// Direct hits received that caused no damage. Value is calculated starting from version 1.10.
		NoDamageDirectHitsReceived int `json:"no_damage_direct_hits_received,omitempty"`
		// Penetrations. Value is calculated starting from version 1.10.
		Piercings int `json:"piercings,omitempty"`
		// Penetrations received. Value is calculated starting from version 1.10.
		PiercingsReceived int `json:"piercings_received,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Total experience
		Xp int `json:"xp,omitempty"`
	} `json:"all,omitempty"`
	// Tank Company battles statistics
	Company struct {
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Average damage upon your spotting. Value is calculated starting from version 1.3.
		DamageAssistedRadio int `json:"damage_assisted_radio,omitempty"`
		// Average damage upon your destroying a track. Value is calculated starting from version 1.3.
		DamageAssistedTrack int `json:"damage_assisted_track,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Direct hits received. Value is calculated starting from version 1.10.
		DirectHitsReceived int `json:"direct_hits_received,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Hits on enemy as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHits int `json:"explosion_hits,omitempty"`
		// Hits received as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHitsReceived int `json:"explosion_hits_received,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Maximum damage caused in a battle.
		MaxDamage int `json:"max_damage,omitempty"`
		// Direct hits received that caused no damage. Value is calculated starting from version 1.10.
		NoDamageDirectHitsReceived int `json:"no_damage_direct_hits_received,omitempty"`
		// Penetrations. Value is calculated starting from version 1.10.
		Piercings int `json:"piercings,omitempty"`
		// Penetrations received. Value is calculated starting from version 1.10.
		PiercingsReceived int `json:"piercings_received,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Total experience
		Xp int `json:"xp,omitempty"`
	} `json:"company,omitempty"`
}

// WotxTanksStats Method returns overall statistics, Tank Company statistics, and clan statistics per each vehicle for each user.
//
// https://developers.wargaming.net/reference/all/wotx/tanks/stats
//
// account_id:
//     Player account ID. Min value is 0.
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
//     Localization language. Default is "en". Valid values:
//     
//     "en" &mdash; English (by default)
//     "ru" &mdash; Русский 
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "tr" &mdash; Türkçe
// tank_id:
//     Player's vehicle ID. Maximum limit: 100.
func (client *Client) WotxTanksStats(realm Realm, accountId int, accessToken string, fields []string, inGarage string, language string, tankId []int) (*WotxTanksStats, error) {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
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

	var result *WotxTanksStats
	err := client.doGetDataRequest(realm, "/wotx/tanks/stats/", reqParam, &result)
	return result, err
}
