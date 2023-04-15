// Auto generated file!

package wotx

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
	// "tr" - Türkçe
	Language *string `json:"language,omitempty"`
	// Player's vehicle ID. Maximum limit: 100.
	TankId []int `json:"tank_id,omitempty"`
}

type TanksStats struct {
	// Player account ID
	AccountId *int `json:"account_id,omitempty"`
	// Overall Statistics
	All *struct {
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Base capture points
		CapturePoints *int `json:"capture_points,omitempty"`
		// Average damage upon your spotting. Value is calculated starting from version 1.3.
		DamageAssistedRadio *int `json:"damage_assisted_radio,omitempty"`
		// Average damage upon your destroying a track. Value is calculated starting from version 1.3.
		DamageAssistedTrack *int `json:"damage_assisted_track,omitempty"`
		// Average damage upon your shooting the wheel.
		DamageAssistedWheel *int `json:"damage_assisted_wheel,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage received
		DamageReceived *int `json:"damage_received,omitempty"`
		// Direct hits received. Value is calculated starting from version 1.10.
		DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
		// Hits on enemy as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHits *int `json:"explosion_hits,omitempty"`
		// Hits received as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
		// Vehicles destroyed
		Frags *int `json:"frags,omitempty"`
		// Hits
		Hits *int `json:"hits,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Maximum damage caused in a battle.
		MaxDamage *int `json:"max_damage,omitempty"`
		// Direct hits received that caused no damage. Value is calculated starting from version 1.10.
		NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
		// Penetrations. Value is calculated starting from version 1.10.
		Piercings *int `json:"piercings,omitempty"`
		// Penetrations received. Value is calculated starting from version 1.10.
		PiercingsReceived *int `json:"piercings_received,omitempty"`
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
	// Total time of destruction
	BattleLifeTime *wgnTime.UnixTime `json:"battle_life_time,omitempty"`
	// Tank Company battles statistics
	Company *struct {
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Base capture points
		CapturePoints *int `json:"capture_points,omitempty"`
		// Average damage upon your spotting. Value is calculated starting from version 1.3.
		DamageAssistedRadio *int `json:"damage_assisted_radio,omitempty"`
		// Average damage upon your destroying a track. Value is calculated starting from version 1.3.
		DamageAssistedTrack *int `json:"damage_assisted_track,omitempty"`
		// Average damage upon your shooting the wheel.
		DamageAssistedWheel *int `json:"damage_assisted_wheel,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage received
		DamageReceived *int `json:"damage_received,omitempty"`
		// Direct hits received. Value is calculated starting from version 1.10.
		DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
		// Hits on enemy as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHits *int `json:"explosion_hits,omitempty"`
		// Hits received as a result of splash damage. Value is calculated starting from version 1.10.
		ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
		// Vehicles destroyed
		Frags *int `json:"frags,omitempty"`
		// Hits
		Hits *int `json:"hits,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Maximum damage caused in a battle.
		MaxDamage *int `json:"max_damage,omitempty"`
		// Direct hits received that caused no damage. Value is calculated starting from version 1.10.
		NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
		// Penetrations. Value is calculated starting from version 1.10.
		Piercings *int `json:"piercings,omitempty"`
		// Penetrations received. Value is calculated starting from version 1.10.
		PiercingsReceived *int `json:"piercings_received,omitempty"`
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
	// Details on vehicles destroyed. This data requires a valid access_token for the specified account.
	Frags map[string]string `json:"frags,omitempty"`
	// Availability of vehicle in the Garage. This data requires a valid access_token for the specified account.
	InGarage *bool `json:"in_garage,omitempty"`
	// Time of status update. This data requires a valid access_token for the specified account.
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
	// Number of battles for markOfMastery achievement
	MarkOfMasteryBattlesCount *int `json:"mark_of_mastery_battles_count,omitempty"`
	// Marks on gun
	MarksOnGun *int `json:"marks_on_gun,omitempty"`
	// Number of battles for marksOnGun achievement
	MarksOnGunBattlesCount *int `json:"marks_on_gun_battles_count,omitempty"`
	// Maximum destroyed in battle
	MaxFrags *int `json:"max_frags,omitempty"`
	// Maximum experience per battle
	MaxXp *int `json:"max_xp,omitempty"`
	// Vehicle ID
	TankId *int `json:"tank_id,omitempty"`
	// Trees knocked down
	TreesCut *int `json:"trees_cut,omitempty"`
}
