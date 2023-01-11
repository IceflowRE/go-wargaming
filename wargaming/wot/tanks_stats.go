// Auto generated file!

package wot

type TanksStatsOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
	// Extra fields that will be added to the response. Valid values:
	//
	// "epic"
	// "fallout"
	// "random"
	// "ranked_10x10"
	// "ranked_battles"
	Extra []string `json:"extra,omitempty"`
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
	} `json:"all,omitempty"`
	// Clan battles statistics
	Clan *struct {
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
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
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
		// Maximum destroyed in battle
		MaxFrags *int `json:"max_frags,omitempty"`
		// Maximum experience per battle
		MaxXp *int `json:"max_xp,omitempty"`
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
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total experience
		Xp *int `json:"xp,omitempty"`
	} `json:"clan,omitempty"`
	// Tank Company battles statistics
	Company *struct {
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
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
		// Vehicles destroyed
		Frags *int `json:"frags,omitempty"`
		// Hits
		Hits *int `json:"hits,omitempty"`
		// Hit ratio
		HitsPercents *int `json:"hits_percents,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
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
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total experience
		Xp *int `json:"xp,omitempty"`
	} `json:"company,omitempty"`
	// Statistics in Grand Battles.
	// An extra field.
	Epic *struct {
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
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
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
		// Maximum destroyed in battle
		MaxFrags *int `json:"max_frags,omitempty"`
		// Maximum experience per battle
		MaxXp *int `json:"max_xp,omitempty"`
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
		// Maximum damage caused in one battle including damage from avatar
		MaxDamageWithAvatar *int `json:"max_damage_with_avatar,omitempty"`
		// Maximum destroyed in battle
		MaxFrags *int `json:"max_frags,omitempty"`
		// Maximum destroyed in one battle including vehicles destroyed by avatar
		MaxFragsWithAvatar *int `json:"max_frags_with_avatar,omitempty"`
		// Maximum Victory Points earned in Fallout
		MaxWinPoints *int `json:"max_win_points,omitempty"`
		// Maximum experience per battle
		MaxXp *int `json:"max_xp,omitempty"`
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
	// Details on vehicles destroyed. This data requires a valid access_token for the specified account.
	Frags map[string]string `json:"frags,omitempty"`
	// All battle statistics on the Global Map
	Globalmap *struct {
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
	} `json:"globalmap,omitempty"`
	// Availability of vehicle in the Garage. This data requires a valid access_token for the specified account.
	InGarage *bool `json:"in_garage,omitempty"`
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
	// Random battles statistics.
	// An extra field.
	Random *struct {
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
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
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
		// Maximum destroyed in battle
		MaxFrags *int `json:"max_frags,omitempty"`
		// Maximum experience per battle
		MaxXp *int `json:"max_xp,omitempty"`
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
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total experience
		Xp *int `json:"xp,omitempty"`
	} `json:"random,omitempty"`
	// Archive of statistics for ranked 10x10 battles.
	// An extra field.
	Ranked10x10 *struct {
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
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
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
		// Maximum destroyed in battle
		MaxFrags *int `json:"max_frags,omitempty"`
		// Maximum experience per battle
		MaxXp *int `json:"max_xp,omitempty"`
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
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total experience
		Xp *int `json:"xp,omitempty"`
	} `json:"ranked_10x10,omitempty"`
	// Statistics in Ranked Battles.
	// An extra field.
	RankedBattles *struct {
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
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
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
		// Maximum destroyed in battle
		MaxFrags *int `json:"max_frags,omitempty"`
		// Maximum experience per battle
		MaxXp *int `json:"max_xp,omitempty"`
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
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total experience
		Xp *int `json:"xp,omitempty"`
	} `json:"ranked_battles,omitempty"`
	// Battle statistics of permanent teams
	RegularTeam *struct {
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
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
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
		// Maximum destroyed in battle
		MaxFrags *int `json:"max_frags,omitempty"`
		// Maximum experience per battle
		MaxXp *int `json:"max_xp,omitempty"`
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
		// Maximum destroyed in battle
		MaxFrags *int `json:"max_frags,omitempty"`
		// Maximum experience per battle
		MaxXp *int `json:"max_xp,omitempty"`
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
		// Maximum destroyed in battle
		MaxFrags *int `json:"max_frags,omitempty"`
		// Maximum experience per battle
		MaxXp *int `json:"max_xp,omitempty"`
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
	// Vehicle ID
	TankId *int `json:"tank_id,omitempty"`
	// Team battles statistics
	Team *struct {
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
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
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
		// Maximum destroyed in battle
		MaxFrags *int `json:"max_frags,omitempty"`
		// Maximum experience per battle
		MaxXp *int `json:"max_xp,omitempty"`
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
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total experience
		Xp *int `json:"xp,omitempty"`
	} `json:"team,omitempty"`
}
