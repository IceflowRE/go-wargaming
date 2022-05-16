package wot

type TanksStats struct {
	// Team battles statistics
	Team struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of times an enemy was stunned by you
		StunNumber int `json:"stun_number,omitempty"`
		// Damage to enemy vehicles stunned by you
		StunAssistedDamage int `json:"stun_assisted_damage,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Maximum experience per battle
		MaxXp int `json:"max_xp,omitempty"`
		// Maximum destroyed in battle
		MaxFrags int `json:"max_frags,omitempty"`
		// Maximum damage caused in a battle
		MaxDamage int `json:"max_damage,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hit ratio
		HitsPercents int `json:"hits_percents,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles on vehicles that cause the stun effect
		BattlesOnStunningVehicles int `json:"battles_on_stunning_vehicles,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average experience per battle
		BattleAvgXp int `json:"battle_avg_xp,omitempty"`
	} `json:"team,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// General stats for player's battles in Skirmishes
	StrongholdSkirmish struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
		// Value is calculated starting from version 9.0.
		TankingFactor float32 `json:"tanking_factor,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of times an enemy was stunned by you
		StunNumber int `json:"stun_number,omitempty"`
		// Damage to enemy vehicles stunned by you
		StunAssistedDamage int `json:"stun_assisted_damage,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Penetrations received
		PiercingsReceived int `json:"piercings_received,omitempty"`
		// Penetrations
		Piercings int `json:"piercings,omitempty"`
		// Direct hits received that caused no damage
		NoDamageDirectHitsReceived int `json:"no_damage_direct_hits_received,omitempty"`
		// Maximum experience per battle
		MaxXp int `json:"max_xp,omitempty"`
		// Maximum destroyed in battle
		MaxFrags int `json:"max_frags,omitempty"`
		// Maximum damage caused in a battle
		MaxDamage int `json:"max_damage,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hit ratio
		HitsPercents int `json:"hits_percents,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Hits received as a result of splash damage
		ExplosionHitsReceived int `json:"explosion_hits_received,omitempty"`
		// Hits on enemy as a result of splash damage
		ExplosionHits int `json:"explosion_hits,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Direct hits received
		DirectHitsReceived int `json:"direct_hits_received,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles on vehicles that cause the stun effect
		BattlesOnStunningVehicles int `json:"battles_on_stunning_vehicles,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average experience per battle
		BattleAvgXp int `json:"battle_avg_xp,omitempty"`
	} `json:"stronghold_skirmish,omitempty"`
	// General stats for player's battles in Stronghold defense
	StrongholdDefense struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
		// Value is calculated starting from version 9.0.
		TankingFactor float32 `json:"tanking_factor,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of times an enemy was stunned by you
		StunNumber int `json:"stun_number,omitempty"`
		// Damage to enemy vehicles stunned by you
		StunAssistedDamage int `json:"stun_assisted_damage,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Penetrations received
		PiercingsReceived int `json:"piercings_received,omitempty"`
		// Penetrations
		Piercings int `json:"piercings,omitempty"`
		// Direct hits received that caused no damage
		NoDamageDirectHitsReceived int `json:"no_damage_direct_hits_received,omitempty"`
		// Maximum experience per battle
		MaxXp int `json:"max_xp,omitempty"`
		// Maximum destroyed in battle
		MaxFrags int `json:"max_frags,omitempty"`
		// Maximum damage caused in a battle
		MaxDamage int `json:"max_damage,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hit ratio
		HitsPercents int `json:"hits_percents,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Hits received as a result of splash damage
		ExplosionHitsReceived int `json:"explosion_hits_received,omitempty"`
		// Hits on enemy as a result of splash damage
		ExplosionHits int `json:"explosion_hits,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Direct hits received
		DirectHitsReceived int `json:"direct_hits_received,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles on vehicles that cause the stun effect
		BattlesOnStunningVehicles int `json:"battles_on_stunning_vehicles,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average experience per battle
		BattleAvgXp int `json:"battle_avg_xp,omitempty"`
	} `json:"stronghold_defense,omitempty"`
	// Battle statistics of permanent teams
	RegularTeam struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of times an enemy was stunned by you
		StunNumber int `json:"stun_number,omitempty"`
		// Damage to enemy vehicles stunned by you
		StunAssistedDamage int `json:"stun_assisted_damage,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Maximum experience per battle
		MaxXp int `json:"max_xp,omitempty"`
		// Maximum destroyed in battle
		MaxFrags int `json:"max_frags,omitempty"`
		// Maximum damage caused in a battle
		MaxDamage int `json:"max_damage,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hit ratio
		HitsPercents int `json:"hits_percents,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles on vehicles that cause the stun effect
		BattlesOnStunningVehicles int `json:"battles_on_stunning_vehicles,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average experience per battle
		BattleAvgXp int `json:"battle_avg_xp,omitempty"`
	} `json:"regular_team,omitempty"`
	// Statistics in Ranked Battles.
	// An extra field.
	RankedBattles struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of times an enemy was stunned by you
		StunNumber int `json:"stun_number,omitempty"`
		// Damage to enemy vehicles stunned by you
		StunAssistedDamage int `json:"stun_assisted_damage,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Maximum experience per battle
		MaxXp int `json:"max_xp,omitempty"`
		// Maximum destroyed in battle
		MaxFrags int `json:"max_frags,omitempty"`
		// Maximum damage caused in a battle
		MaxDamage int `json:"max_damage,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hit ratio
		HitsPercents int `json:"hits_percents,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles on vehicles that cause the stun effect
		BattlesOnStunningVehicles int `json:"battles_on_stunning_vehicles,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average experience per battle
		BattleAvgXp int `json:"battle_avg_xp,omitempty"`
	} `json:"ranked_battles,omitempty"`
	// Archive of statistics for ranked 10x10 battles.
	// An extra field.
	Ranked10X10 struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of times an enemy was stunned by you
		StunNumber int `json:"stun_number,omitempty"`
		// Damage to enemy vehicles stunned by you
		StunAssistedDamage int `json:"stun_assisted_damage,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Maximum experience per battle
		MaxXp int `json:"max_xp,omitempty"`
		// Maximum destroyed in battle
		MaxFrags int `json:"max_frags,omitempty"`
		// Maximum damage caused in a battle
		MaxDamage int `json:"max_damage,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hit ratio
		HitsPercents int `json:"hits_percents,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles on vehicles that cause the stun effect
		BattlesOnStunningVehicles int `json:"battles_on_stunning_vehicles,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average experience per battle
		BattleAvgXp int `json:"battle_avg_xp,omitempty"`
	} `json:"ranked_10x10,omitempty"`
	// Random battles statistics.
	// An extra field.
	Random struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of times an enemy was stunned by you
		StunNumber int `json:"stun_number,omitempty"`
		// Damage to enemy vehicles stunned by you
		StunAssistedDamage int `json:"stun_assisted_damage,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Maximum experience per battle
		MaxXp int `json:"max_xp,omitempty"`
		// Maximum destroyed in battle
		MaxFrags int `json:"max_frags,omitempty"`
		// Maximum damage caused in a battle
		MaxDamage int `json:"max_damage,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hit ratio
		HitsPercents int `json:"hits_percents,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles on vehicles that cause the stun effect
		BattlesOnStunningVehicles int `json:"battles_on_stunning_vehicles,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average experience per battle
		BattleAvgXp int `json:"battle_avg_xp,omitempty"`
	} `json:"random,omitempty"`
	// Maximum experience per battle
	MaxXp int `json:"max_xp,omitempty"`
	// Maximum destroyed in battle
	MaxFrags int `json:"max_frags,omitempty"`
	// Mastery Badges:
	// 
	// 0 — None
	// 1 — 3rd Class 
	// 2 — 2nd Class
	// 3 — 1st Class
	// 4 — Ace Tanker
	MarkOfMastery int `json:"mark_of_mastery,omitempty"`
	// Availability of vehicle in the Garage. This data requires a valid access_token for the specified account.
	InGarage bool `json:"in_garage,omitempty"`
	// All battle statistics on the Global Map
	Globalmap struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
		// Value is calculated starting from version 9.0.
		TankingFactor float32 `json:"tanking_factor,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of times an enemy was stunned by you
		StunNumber int `json:"stun_number,omitempty"`
		// Damage to enemy vehicles stunned by you
		StunAssistedDamage int `json:"stun_assisted_damage,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Penetrations received
		PiercingsReceived int `json:"piercings_received,omitempty"`
		// Penetrations
		Piercings int `json:"piercings,omitempty"`
		// Direct hits received that caused no damage
		NoDamageDirectHitsReceived int `json:"no_damage_direct_hits_received,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hit ratio
		HitsPercents int `json:"hits_percents,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Hits received as a result of splash damage
		ExplosionHitsReceived int `json:"explosion_hits_received,omitempty"`
		// Hits on enemy as a result of splash damage
		ExplosionHits int `json:"explosion_hits,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Direct hits received
		DirectHitsReceived int `json:"direct_hits_received,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles on vehicles that cause the stun effect
		BattlesOnStunningVehicles int `json:"battles_on_stunning_vehicles,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average experience per battle
		BattleAvgXp int `json:"battle_avg_xp,omitempty"`
		// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
		// Value is calculated starting from version 9.0.
		AvgDamageBlocked float32 `json:"avg_damage_blocked,omitempty"`
		// Average damage upon your shooting the track
		AvgDamageAssistedTrack float32 `json:"avg_damage_assisted_track,omitempty"`
		// Average damage upon your spotting
		AvgDamageAssistedRadio float32 `json:"avg_damage_assisted_radio,omitempty"`
		// Average damage caused with your assistance
		AvgDamageAssisted float32 `json:"avg_damage_assisted,omitempty"`
	} `json:"globalmap,omitempty"`
	// Details on vehicles destroyed. This data requires a valid access_token for the specified account.
	Frags map[string]string `json:"frags,omitempty"`
	// Fallout statistics.
	// An extra field.
	Fallout struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Victory Points
		WinPoints int `json:"win_points,omitempty"`
		// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
		// Value is calculated starting from version 9.0.
		TankingFactor float32 `json:"tanking_factor,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of times an enemy was stunned by you
		StunNumber int `json:"stun_number,omitempty"`
		// Damage to enemy vehicles stunned by you
		StunAssistedDamage int `json:"stun_assisted_damage,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Resources captured at resource points
		ResourceAbsorbed int `json:"resource_absorbed,omitempty"`
		// Penetrations received
		PiercingsReceived int `json:"piercings_received,omitempty"`
		// Penetrations
		Piercings int `json:"piercings,omitempty"`
		// Direct hits received that caused no damage
		NoDamageDirectHitsReceived int `json:"no_damage_direct_hits_received,omitempty"`
		// Maximum experience per battle
		MaxXp int `json:"max_xp,omitempty"`
		// Maximum Victory Points earned in Fallout
		MaxWinPoints int `json:"max_win_points,omitempty"`
		// Maximum destroyed in one battle including vehicles destroyed by avatar
		MaxFragsWithAvatar int `json:"max_frags_with_avatar,omitempty"`
		// Maximum destroyed in battle
		MaxFrags int `json:"max_frags,omitempty"`
		// Maximum damage caused in one battle including damage from avatar
		MaxDamageWithAvatar int `json:"max_damage_with_avatar,omitempty"`
		// Maximum damage caused in a battle
		MaxDamage int `json:"max_damage,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hit ratio
		HitsPercents int `json:"hits_percents,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Flags captured as solo player
		FlagCaptureSolo int `json:"flag_capture_solo,omitempty"`
		// Flags captured in platoon
		FlagCapture int `json:"flag_capture,omitempty"`
		// Hits received as a result of splash damage
		ExplosionHitsReceived int `json:"explosion_hits_received,omitempty"`
		// Hits on enemy as a result of splash damage
		ExplosionHits int `json:"explosion_hits,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Direct hits received
		DirectHitsReceived int `json:"direct_hits_received,omitempty"`
		// Deaths
		DeathCount int `json:"death_count,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles on vehicles that cause the stun effect
		BattlesOnStunningVehicles int `json:"battles_on_stunning_vehicles,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average experience per battle
		BattleAvgXp int `json:"battle_avg_xp,omitempty"`
		// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
		// Value is calculated starting from version 9.0.
		AvgDamageBlocked float32 `json:"avg_damage_blocked,omitempty"`
		// Average damage upon your shooting the track
		AvgDamageAssistedTrack float32 `json:"avg_damage_assisted_track,omitempty"`
		// Average damage upon your spotting
		AvgDamageAssistedRadio float32 `json:"avg_damage_assisted_radio,omitempty"`
		// Average damage caused with your assistance
		AvgDamageAssisted float32 `json:"avg_damage_assisted,omitempty"`
		// Destroyed by Combat Reserves
		AvatarFrags int `json:"avatar_frags,omitempty"`
		// Damage caused by Combat Reserves
		AvatarDamageDealt int `json:"avatar_damage_dealt,omitempty"`
	} `json:"fallout,omitempty"`
	// Statistics in Grand Battles.
	// An extra field.
	Epic struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of times an enemy was stunned by you
		StunNumber int `json:"stun_number,omitempty"`
		// Damage to enemy vehicles stunned by you
		StunAssistedDamage int `json:"stun_assisted_damage,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Maximum experience per battle
		MaxXp int `json:"max_xp,omitempty"`
		// Maximum destroyed in battle
		MaxFrags int `json:"max_frags,omitempty"`
		// Maximum damage caused in a battle
		MaxDamage int `json:"max_damage,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hit ratio
		HitsPercents int `json:"hits_percents,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles on vehicles that cause the stun effect
		BattlesOnStunningVehicles int `json:"battles_on_stunning_vehicles,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average experience per battle
		BattleAvgXp int `json:"battle_avg_xp,omitempty"`
	} `json:"epic,omitempty"`
	// Tank Company battles statistics
	Company struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of times an enemy was stunned by you
		StunNumber int `json:"stun_number,omitempty"`
		// Damage to enemy vehicles stunned by you
		StunAssistedDamage int `json:"stun_assisted_damage,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hit ratio
		HitsPercents int `json:"hits_percents,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles on vehicles that cause the stun effect
		BattlesOnStunningVehicles int `json:"battles_on_stunning_vehicles,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average experience per battle
		BattleAvgXp int `json:"battle_avg_xp,omitempty"`
	} `json:"company,omitempty"`
	// Clan battles statistics
	Clan struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of times an enemy was stunned by you
		StunNumber int `json:"stun_number,omitempty"`
		// Damage to enemy vehicles stunned by you
		StunAssistedDamage int `json:"stun_assisted_damage,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Maximum experience per battle
		MaxXp int `json:"max_xp,omitempty"`
		// Maximum destroyed in battle
		MaxFrags int `json:"max_frags,omitempty"`
		// Maximum damage caused in a battle
		MaxDamage int `json:"max_damage,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hit ratio
		HitsPercents int `json:"hits_percents,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles on vehicles that cause the stun effect
		BattlesOnStunningVehicles int `json:"battles_on_stunning_vehicles,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average experience per battle
		BattleAvgXp int `json:"battle_avg_xp,omitempty"`
	} `json:"clan,omitempty"`
	// Overall Statistics
	All struct {
		// Total experience
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Ratio of damage blocked by armor from AP, HEAT, and APCR shells to damage received from these types of shells.
		// Value is calculated starting from version 9.0.
		TankingFactor float32 `json:"tanking_factor,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of times an enemy was stunned by you
		StunNumber int `json:"stun_number,omitempty"`
		// Damage to enemy vehicles stunned by you
		StunAssistedDamage int `json:"stun_assisted_damage,omitempty"`
		// Enemies spotted
		Spotted int `json:"spotted,omitempty"`
		// Shots fired
		Shots int `json:"shots,omitempty"`
		// Penetrations received
		PiercingsReceived int `json:"piercings_received,omitempty"`
		// Penetrations
		Piercings int `json:"piercings,omitempty"`
		// Direct hits received that caused no damage
		NoDamageDirectHitsReceived int `json:"no_damage_direct_hits_received,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Hit ratio
		HitsPercents int `json:"hits_percents,omitempty"`
		// Hits
		Hits int `json:"hits,omitempty"`
		// Vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Hits received as a result of splash damage
		ExplosionHitsReceived int `json:"explosion_hits_received,omitempty"`
		// Hits on enemy as a result of splash damage
		ExplosionHits int `json:"explosion_hits,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Direct hits received
		DirectHitsReceived int `json:"direct_hits_received,omitempty"`
		// Damage received
		DamageReceived int `json:"damage_received,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles on vehicles that cause the stun effect
		BattlesOnStunningVehicles int `json:"battles_on_stunning_vehicles,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Average experience per battle
		BattleAvgXp int `json:"battle_avg_xp,omitempty"`
		// Average damage blocked by armor per battle. Damage blocked by armor is damage received from shells (AP, HEAT and APCR) that hit a vehicle but caused no damage.
		// Value is calculated starting from version 9.0.
		AvgDamageBlocked float32 `json:"avg_damage_blocked,omitempty"`
	} `json:"all,omitempty"`
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
} 
