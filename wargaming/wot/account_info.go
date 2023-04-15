// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type AccountInfoOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
	// Extra fields that will be added to the response. Valid values:
	//
	// "private.boosters"
	// "private.garage"
	// "private.grouped_contacts"
	// "private.personal_missions"
	// "private.rented"
	// "statistics.epic"
	// "statistics.fallout"
	// "statistics.globalmap_absolute"
	// "statistics.globalmap_champion"
	// "statistics.globalmap_middle"
	// "statistics.random"
	// "statistics.ranked_10x10"
	// "statistics.ranked_15x15"
	// "statistics.ranked_battles"
	// "statistics.ranked_battles_current"
	// "statistics.ranked_battles_previous"
	// "statistics.ranked_season_1"
	// "statistics.ranked_season_2"
	// "statistics.ranked_season_3"
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
	// Clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Language selected in the game client
	ClientLanguage *string `json:"client_language,omitempty"`
	// Date when player's account was created
	CreatedAt *wgnTime.UnixTime `json:"created_at,omitempty"`
	// Personal rating
	GlobalRating *int `json:"global_rating,omitempty"`
	// Last battle time
	LastBattleTime *wgnTime.UnixTime `json:"last_battle_time,omitempty"`
	// End time of last game session
	LogoutAt *wgnTime.UnixTime `json:"logout_at,omitempty"`
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
		// Bonds
		Bonds *int `json:"bonds,omitempty"`
		// Personal Reserves.
		// An extra field.
		Boosters map[string]*struct {
			// Amount of Personal Reserves
			Count *int `json:"count,omitempty"`
			// Expiration time
			ExpirationTime *wgnTime.UnixTime `json:"expiration_time,omitempty"`
			// Status of Personal Reserves. Valid values:
			//
			// ACTIVE - Active
			// INACTIVE - Inactive
			// USED - Used
			State *string `json:"state,omitempty"`
		} `json:"boosters,omitempty"`
		// Credits
		Credits *int `json:"credits,omitempty"`
		// Free Experience
		FreeXp *int `json:"free_xp,omitempty"`
		// Vehicles in the Garage.
		// An extra field.
		Garage []int `json:"garage,omitempty"`
		// Gold
		Gold *int `json:"gold,omitempty"`
		// Contact groups.
		// An extra field.
		GroupedContacts *struct {
			// The contact was added to the blacklist. Note that the blocked contact still belongs to contact groups or to the ungrouped list of contacts.
			Blocked []int `json:"blocked,omitempty"`
			// Groups
			Groups map[string]string `json:"groups,omitempty"`
			// Ignored
			Ignored []int `json:"ignored,omitempty"`
			// Muted
			Muted []int `json:"muted,omitempty"`
			// Not grouped
			Ungrouped []int `json:"ungrouped,omitempty"`
		} `json:"grouped_contacts,omitempty"`
		// Indicates if mobile phone number was added to the account
		IsBoundToPhone *bool `json:"is_bound_to_phone,omitempty"`
		// Indicates if the account is Premium Account
		IsPremium *bool `json:"is_premium,omitempty"`
		// Personal Missions Progress. The key is a task id, the value is a status.
		// Possible statuses:
		//
		// NONE - a mission is unavailable
		// UNLOCKED - a mission is available
		// NEED_GET_MAIN_REWARD - the main reward has not been received
		// MAIN_REWARD_GOTTEN - the main reward has been received
		// NEED_GET_ADD_REWARD - additional reward has not been received
		// NEED_GET_ALL_REWARDS - no rewards have been received
		// ALL_REWARDS_GOTTEN - all rewards have been received
		// .
		//
		// An extra field.
		PersonalMissions map[string]string `json:"personal_missions,omitempty"`
		// Premium Account expiration time
		PremiumExpiresAt *wgnTime.UnixTime `json:"premium_expires_at,omitempty"`
		// Vehicle Rental.
		// An extra field.
		Rented *struct {
			// Rental compensation in credits
			CompensationCredits *int `json:"compensation_credits,omitempty"`
			// Rental compensation in gold
			CompensationGold *int `json:"compensation_gold,omitempty"`
			// Vehicle Rental expiration time
			ExpirationTime *wgnTime.UnixTime `json:"expiration_time,omitempty"`
			// Rented vehicle ID
			TankId *int `json:"tank_id,omitempty"`
		} `json:"rented,omitempty"`
		// Account restrictions
		Restrictions *struct {
			// End time of chat ban
			ChatBanTime *wgnTime.UnixTime `json:"chat_ban_time,omitempty"`
		} `json:"restrictions,omitempty"`
	} `json:"private,omitempty"`
	// Player statistics
	Statistics *struct {
		// Total statistics in Random and clan battles without the Global Map 2.0 statistics
		All *struct {
			// Average damage caused with your assistance.
			// Value is calculated starting from version 8.8.
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting.
			// Value is calculated starting from version 8.8.
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track.
			// Value is calculated starting from version 8.8.
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"all,omitempty"`
		// Clan battles statistics
		Clan *struct {
			// Average damage caused with your assistance.
			// Value is calculated starting from version 8.9.
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting.
			// Value is calculated starting from version 8.9.
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track.
			// Value is calculated starting from version 8.9.
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"clan,omitempty"`
		// Tank Company battles statistics
		Company *struct {
			// Average damage caused with your assistance.
			// Value is calculated starting from version 8.9.
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting.
			// Value is calculated starting from version 8.9.
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track.
			// Value is calculated starting from version 8.9.
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"company,omitempty"`
		// Statistics in Grand Battles.
		// An extra field.
		Epic *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"epic,omitempty"`
		// Fallout statistics.
		// An extra field.
		Fallout *struct {
			// Damage caused by Combat Reserves
			AvatarDamageDealt *int `json:"avatar_damage_dealt,omitempty"`
			// Destroyed by Combat Reserves
			AvatarFrags *int `json:"avatar_frags,omitempty"`
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Deaths
			DeathCount *int `json:"death_count,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Flags captured in platoon
			FlagCapture *int `json:"flag_capture,omitempty"`
			// Flags captured as solo player
			FlagCaptureSolo *int `json:"flag_capture_solo,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
			// Maximum damage caused in a battle
			MaxDamage *int `json:"max_damage,omitempty"`
			// Vehicle used to cause maximum damage
			MaxDamageTankId *int `json:"max_damage_tank_id,omitempty"`
			// Maximum damage caused in one battle including damage from avatar
			MaxDamageWithAvatar *int `json:"max_damage_with_avatar,omitempty"`
			// Maximum destroyed in battle
			MaxFrags *int `json:"max_frags,omitempty"`
			// Vehicle, in which maximum number of enemy vehicles was destroyed
			MaxFragsTankId *int `json:"max_frags_tank_id,omitempty"`
			// Maximum destroyed in one battle including vehicles destroyed by avatar
			MaxFragsWithAvatar *int `json:"max_frags_with_avatar,omitempty"`
			// Maximum Victory Points earned in Fallout
			MaxWinPoints *int `json:"max_win_points,omitempty"`
			// Maximum experience per battle
			MaxXp *int `json:"max_xp,omitempty"`
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Resources captured at resource points
			ResourceAbsorbed *int `json:"resource_absorbed,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victory Points
			WinPoints *int `json:"win_points,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"fallout,omitempty"`
		// Number and models of vehicles destroyed by a player. Player's private data.
		Frags map[string]int `json:"frags,omitempty"`
		// Battle statistics on the Global Map in Absolute division.
		// An extra field.
		GlobalmapAbsolute *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"globalmap_absolute,omitempty"`
		// Battle statistics on the Global Map in Champion division.
		// An extra field.
		GlobalmapChampion *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"globalmap_champion,omitempty"`
		// Battle statistics on the Global Map in Medium division.
		// An extra field.
		GlobalmapMiddle *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"globalmap_middle,omitempty"`
		// Historical battles statistics
		Historical *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"historical,omitempty"`
		// Random battles statistics.
		// An extra field.
		Random *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"random,omitempty"`
		// Archive of statistics for ranked 10x10 battles.
		// An extra field.
		Ranked10x10 *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"ranked_10x10,omitempty"`
		// Archive of statistics for ranked 15x15 battles.
		// An extra field.
		Ranked15x15 *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"ranked_15x15,omitempty"`
		// Ranked Battles statistics.
		// An extra field.
		RankedBattles *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"ranked_battles,omitempty"`
		// Current Ranked Battles statistics.
		// An extra field.
		RankedBattlesCurrent *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"ranked_battles_current,omitempty"`
		// Previous Ranked Battles statistics.
		// An extra field.
		RankedBattlesPrevious *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"ranked_battles_previous,omitempty"`
		// Statistics of ranked battles for the first season of the year. Updated with every battle this season, data is reseted at the end of the year..
		// An extra field.
		RankedSeason1 *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"ranked_season_1,omitempty"`
		// Statistics of ranked battles for the second season of the year. Updated with every battle this season, data is reseted at the end of the year..
		// An extra field.
		RankedSeason2 *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"ranked_season_2,omitempty"`
		// Statistics of ranked battles for the third season of the year. Updated with every battle this season, data is reseted at the end of the year..
		// An extra field.
		RankedSeason3 *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"ranked_season_3,omitempty"`
		// Battle statistics of permanent teams
		RegularTeam *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"regular_team,omitempty"`
		// General stats for player's battles in Stronghold defense
		StrongholdDefense *struct {
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"stronghold_defense,omitempty"`
		// General stats for player's battles in Skirmishes
		StrongholdSkirmish *struct {
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"stronghold_skirmish,omitempty"`
		// Team battles statistics
		Team *struct {
			// Average damage caused with your assistance
			AvgDamageAssisted *float32 `json:"avg_damage_assisted,omitempty"`
			// Average damage upon your spotting
			AvgDamageAssistedRadio *float32 `json:"avg_damage_assisted_radio,omitempty"`
			// Average damage upon your shooting the track
			AvgDamageAssistedTrack *float32 `json:"avg_damage_assisted_track,omitempty"`
			// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
			// Value is calculated starting from version 9.0.
			AvgDamageBlocked *float32 `json:"avg_damage_blocked,omitempty"`
			// Average experience per battle
			BattleAvgXp *int `json:"battle_avg_xp,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Number of battles on vehicles that cause the stun effect
			BattlesOnStunningVehicles *int `json:"battles_on_stunning_vehicles,omitempty"`
			// Base capture points
			CapturePoints *int `json:"capture_points,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
			// Damage received
			DamageReceived *int `json:"damage_received,omitempty"`
			// Direct hits received
			DirectHitsReceived *int `json:"direct_hits_received,omitempty"`
			// Draws
			Draws *int `json:"draws,omitempty"`
			// Base defense points
			DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
			// Hits on enemy as a result of splash damage
			ExplosionHits *int `json:"explosion_hits,omitempty"`
			// Hits received as a result of splash damage
			ExplosionHitsReceived *int `json:"explosion_hits_received,omitempty"`
			// Vehicles destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Hit ratio
			HitsPercents *int `json:"hits_percents,omitempty"`
			// Defeats
			Losses *int `json:"losses,omitempty"`
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
			// Vehicle used to gain maximum experience per battle
			MaxXpTankId *int `json:"max_xp_tank_id,omitempty"`
			// Direct hits received that caused no damage
			NoDamageDirectHitsReceived *int `json:"no_damage_direct_hits_received,omitempty"`
			// Penetrations
			Piercings *int `json:"piercings,omitempty"`
			// Penetrations received
			PiercingsReceived *int `json:"piercings_received,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
			// Enemies spotted
			Spotted *int `json:"spotted,omitempty"`
			// Damage to enemy vehicles stunned by you
			StunAssistedDamage *int `json:"stun_assisted_damage,omitempty"`
			// Number of times an enemy was stunned by you
			StunNumber *int `json:"stun_number,omitempty"`
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
			// Value is calculated starting from version 9.0.
			TankingFactor *float32 `json:"tanking_factor,omitempty"`
			// Victories
			Wins *int `json:"wins,omitempty"`
			// Total experience
			Xp *int `json:"xp,omitempty"`
		} `json:"team,omitempty"`
		// Trees knocked down
		TreesCut *int `json:"trees_cut,omitempty"`
	} `json:"statistics,omitempty"`
	// Date when player details were updated
	UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
}
