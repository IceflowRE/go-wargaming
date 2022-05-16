package wows

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type ShipsStats struct {
	// Date when details on ships were updated
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Ship ID
	ShipId int `json:"ship_id,omitempty"`
	// Player statistics in Ranked Battles in solo game.
	// An extra field.
	RankSolo struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro int `json:"torpedo_agro,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints int `json:"team_capture_points,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted int `json:"ships_spotted,omitempty"`
		// Secondary armament firing statistics
		SecondBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"second_battery,omitempty"`
		// Statistics of ramming enemy warships
		Ramming struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"ramming,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled int `json:"planes_killed,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp int `json:"max_xp,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro int `json:"max_total_agro,omitempty"`
		// Most ships detected
		MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle int `json:"max_frags_battle,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
		// Main battery firing statistics
		MainBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"main_battery,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Warships destroyed
		Frags int `json:"frags,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage caused by allies
		DamageScouting int `json:"damage_scouting,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Potential damage caused by shells
		ArtAgro int `json:"art_agro,omitempty"`
		// Statistics of aircraft used
		Aircraft struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"aircraft,omitempty"`
	} `json:"rank_solo,omitempty"`
	// Player statistics in Ranked Battles in division of 3 players.
	// An extra field.
	RankDiv3 struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro int `json:"torpedo_agro,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints int `json:"team_capture_points,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted int `json:"ships_spotted,omitempty"`
		// Secondary armament firing statistics
		SecondBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"second_battery,omitempty"`
		// Statistics of ramming enemy warships
		Ramming struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"ramming,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled int `json:"planes_killed,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp int `json:"max_xp,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro int `json:"max_total_agro,omitempty"`
		// Most ships detected
		MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle int `json:"max_frags_battle,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
		// Main battery firing statistics
		MainBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"main_battery,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Warships destroyed
		Frags int `json:"frags,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage caused by allies
		DamageScouting int `json:"damage_scouting,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Potential damage caused by shells
		ArtAgro int `json:"art_agro,omitempty"`
		// Statistics of aircraft used
		Aircraft struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"aircraft,omitempty"`
	} `json:"rank_div3,omitempty"`
	// Player statistics in Ranked Battles in division of 2 players.
	// An extra field.
	RankDiv2 struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro int `json:"torpedo_agro,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints int `json:"team_capture_points,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted int `json:"ships_spotted,omitempty"`
		// Secondary armament firing statistics
		SecondBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"second_battery,omitempty"`
		// Statistics of ramming enemy warships
		Ramming struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"ramming,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled int `json:"planes_killed,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp int `json:"max_xp,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro int `json:"max_total_agro,omitempty"`
		// Most ships detected
		MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle int `json:"max_frags_battle,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
		// Main battery firing statistics
		MainBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"main_battery,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Warships destroyed
		Frags int `json:"frags,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage caused by allies
		DamageScouting int `json:"damage_scouting,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Potential damage caused by shells
		ArtAgro int `json:"art_agro,omitempty"`
		// Statistics of aircraft used
		Aircraft struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"aircraft,omitempty"`
	} `json:"rank_div2,omitempty"`
	// Player statistics in solo battles fought in Random mode.
	// An extra field.
	PvpSolo struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro int `json:"torpedo_agro,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints int `json:"team_capture_points,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of installations surpressed, starting from version 5.12
		SuppressionsCount int `json:"suppressions_count,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted int `json:"ships_spotted,omitempty"`
		// Secondary armament firing statistics
		SecondBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"second_battery,omitempty"`
		// Statistics of ramming enemy warships
		Ramming struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"ramming,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled int `json:"planes_killed,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp int `json:"max_xp,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro int `json:"max_total_agro,omitempty"`
		// Most installations surpressed in a battle
		MaxSuppressionsCount int `json:"max_suppressions_count,omitempty"`
		// Most ships detected
		MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle int `json:"max_frags_battle,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
		// Most damage caused to installations in a battle
		MaxDamageDealtToBuildings int `json:"max_damage_dealt_to_buildings,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
		// Main battery firing statistics
		MainBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"main_battery,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Warships destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage dealt to installations, starting from version 5.12
		DamageToBuildings int `json:"damage_to_buildings,omitempty"`
		// Damage caused by allies
		DamageScouting int `json:"damage_scouting,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles, starting from version 5.12
		BattlesSince512 int `json:"battles_since_512,omitempty"`
		// Number of battles, starting from version 5.10
		BattlesSince510 int `json:"battles_since_510,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Potential damage caused by shells
		ArtAgro int `json:"art_agro,omitempty"`
		// Statistics of aircraft used
		Aircraft struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"aircraft,omitempty"`
	} `json:"pvp_solo,omitempty"`
	// Player statistics in Random Battles in Division of 3 players.
	// An extra field.
	PvpDiv3 struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro int `json:"torpedo_agro,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints int `json:"team_capture_points,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of installations surpressed, starting from version 5.12
		SuppressionsCount int `json:"suppressions_count,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted int `json:"ships_spotted,omitempty"`
		// Secondary armament firing statistics
		SecondBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"second_battery,omitempty"`
		// Statistics of ramming enemy warships
		Ramming struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"ramming,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled int `json:"planes_killed,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp int `json:"max_xp,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro int `json:"max_total_agro,omitempty"`
		// Most installations surpressed in a battle
		MaxSuppressionsCount int `json:"max_suppressions_count,omitempty"`
		// Most ships detected
		MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle int `json:"max_frags_battle,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
		// Most damage caused to installations in a battle
		MaxDamageDealtToBuildings int `json:"max_damage_dealt_to_buildings,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
		// Main battery firing statistics
		MainBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"main_battery,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Warships destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage dealt to installations, starting from version 5.12
		DamageToBuildings int `json:"damage_to_buildings,omitempty"`
		// Damage caused by allies
		DamageScouting int `json:"damage_scouting,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles, starting from version 5.12
		BattlesSince512 int `json:"battles_since_512,omitempty"`
		// Number of battles, starting from version 5.10
		BattlesSince510 int `json:"battles_since_510,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Potential damage caused by shells
		ArtAgro int `json:"art_agro,omitempty"`
		// Statistics of aircraft used
		Aircraft struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"aircraft,omitempty"`
	} `json:"pvp_div3,omitempty"`
	// Player statistics in Random Battles in Division of 2 players.
	// An extra field.
	PvpDiv2 struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro int `json:"torpedo_agro,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints int `json:"team_capture_points,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of installations surpressed, starting from version 5.12
		SuppressionsCount int `json:"suppressions_count,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted int `json:"ships_spotted,omitempty"`
		// Secondary armament firing statistics
		SecondBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"second_battery,omitempty"`
		// Statistics of ramming enemy warships
		Ramming struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"ramming,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled int `json:"planes_killed,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp int `json:"max_xp,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro int `json:"max_total_agro,omitempty"`
		// Most installations surpressed in a battle
		MaxSuppressionsCount int `json:"max_suppressions_count,omitempty"`
		// Most ships detected
		MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle int `json:"max_frags_battle,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
		// Most damage caused to installations in a battle
		MaxDamageDealtToBuildings int `json:"max_damage_dealt_to_buildings,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
		// Main battery firing statistics
		MainBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"main_battery,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Warships destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage dealt to installations, starting from version 5.12
		DamageToBuildings int `json:"damage_to_buildings,omitempty"`
		// Damage caused by allies
		DamageScouting int `json:"damage_scouting,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles, starting from version 5.12
		BattlesSince512 int `json:"battles_since_512,omitempty"`
		// Number of battles, starting from version 5.10
		BattlesSince510 int `json:"battles_since_510,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Potential damage caused by shells
		ArtAgro int `json:"art_agro,omitempty"`
		// Statistics of aircraft used
		Aircraft struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"aircraft,omitempty"`
	} `json:"pvp_div2,omitempty"`
	// Player statistics in all Random Battles
	Pvp struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro int `json:"torpedo_agro,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints int `json:"team_capture_points,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Number of installations surpressed, starting from version 5.12
		SuppressionsCount int `json:"suppressions_count,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted int `json:"ships_spotted,omitempty"`
		// Secondary armament firing statistics
		SecondBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"second_battery,omitempty"`
		// Statistics of ramming enemy warships
		Ramming struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"ramming,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled int `json:"planes_killed,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp int `json:"max_xp,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro int `json:"max_total_agro,omitempty"`
		// Most installations surpressed in a battle
		MaxSuppressionsCount int `json:"max_suppressions_count,omitempty"`
		// Most ships detected
		MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle int `json:"max_frags_battle,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
		// Most damage caused to installations in a battle
		MaxDamageDealtToBuildings int `json:"max_damage_dealt_to_buildings,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
		// Main battery firing statistics
		MainBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"main_battery,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Warships destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage dealt to installations, starting from version 5.12
		DamageToBuildings int `json:"damage_to_buildings,omitempty"`
		// Damage caused by allies
		DamageScouting int `json:"damage_scouting,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Number of battles, starting from version 5.12
		BattlesSince512 int `json:"battles_since_512,omitempty"`
		// Number of battles, starting from version 5.10
		BattlesSince510 int `json:"battles_since_510,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Potential damage caused by shells
		ArtAgro int `json:"art_agro,omitempty"`
		// Statistics of aircraft used
		Aircraft struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"aircraft,omitempty"`
	} `json:"pvp,omitempty"`
	// Player statistics in solo battles fought in Co-op mode.
	// An extra field.
	PveSolo struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro int `json:"torpedo_agro,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints int `json:"team_capture_points,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted int `json:"ships_spotted,omitempty"`
		// Secondary armament firing statistics
		SecondBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"second_battery,omitempty"`
		// Statistics of ramming enemy warships
		Ramming struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"ramming,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled int `json:"planes_killed,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp int `json:"max_xp,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro int `json:"max_total_agro,omitempty"`
		// Most ships detected
		MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle int `json:"max_frags_battle,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
		// Main battery firing statistics
		MainBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"main_battery,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Warships destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage caused by allies
		DamageScouting int `json:"damage_scouting,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Potential damage caused by shells
		ArtAgro int `json:"art_agro,omitempty"`
		// Statistics of aircraft used
		Aircraft struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"aircraft,omitempty"`
	} `json:"pve_solo,omitempty"`
	// Player statistics in Co-op Battles in Division of 3 players.
	// An extra field.
	PveDiv3 struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro int `json:"torpedo_agro,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints int `json:"team_capture_points,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted int `json:"ships_spotted,omitempty"`
		// Secondary armament firing statistics
		SecondBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"second_battery,omitempty"`
		// Statistics of ramming enemy warships
		Ramming struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"ramming,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled int `json:"planes_killed,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp int `json:"max_xp,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro int `json:"max_total_agro,omitempty"`
		// Most ships detected
		MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle int `json:"max_frags_battle,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
		// Main battery firing statistics
		MainBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"main_battery,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Warships destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage caused by allies
		DamageScouting int `json:"damage_scouting,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Potential damage caused by shells
		ArtAgro int `json:"art_agro,omitempty"`
		// Statistics of aircraft used
		Aircraft struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"aircraft,omitempty"`
	} `json:"pve_div3,omitempty"`
	// Player statistics in Co-op Battles in Division of 2 players.
	// An extra field.
	PveDiv2 struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro int `json:"torpedo_agro,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints int `json:"team_capture_points,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted int `json:"ships_spotted,omitempty"`
		// Secondary armament firing statistics
		SecondBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"second_battery,omitempty"`
		// Statistics of ramming enemy warships
		Ramming struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"ramming,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled int `json:"planes_killed,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp int `json:"max_xp,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro int `json:"max_total_agro,omitempty"`
		// Most ships detected
		MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle int `json:"max_frags_battle,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
		// Main battery firing statistics
		MainBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"main_battery,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Warships destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage caused by allies
		DamageScouting int `json:"damage_scouting,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Potential damage caused by shells
		ArtAgro int `json:"art_agro,omitempty"`
		// Statistics of aircraft used
		Aircraft struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"aircraft,omitempty"`
	} `json:"pve_div2,omitempty"`
	// Player statistics in all Co-op Battles.
	// An extra field.
	Pve struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro int `json:"torpedo_agro,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints int `json:"team_capture_points,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted int `json:"ships_spotted,omitempty"`
		// Secondary armament firing statistics
		SecondBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"second_battery,omitempty"`
		// Statistics of ramming enemy warships
		Ramming struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"ramming,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled int `json:"planes_killed,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp int `json:"max_xp,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro int `json:"max_total_agro,omitempty"`
		// Most ships detected
		MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle int `json:"max_frags_battle,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
		// Main battery firing statistics
		MainBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"main_battery,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Warships destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage caused by allies
		DamageScouting int `json:"damage_scouting,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Potential damage caused by shells
		ArtAgro int `json:"art_agro,omitempty"`
		// Statistics of aircraft used
		Aircraft struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"aircraft,omitempty"`
	} `json:"pve,omitempty"`
	// Private data on the player's ships
	Private struct {
		// Availability of ships in the Port
		InGarage bool `json:"in_garage,omitempty"`
		// Overall battle life time in seconds
		BattleLifeTime int `json:"battle_life_time,omitempty"`
	} `json:"private,omitempty"`
	// Player's statistics for playing solo the Operation mode, normal difficulty..
	// An extra field.
	OperSolo struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories for completing missions. The key is a number of completed missions, and the value is a number of victories.
		WinsByTasks map[string]string `json:"wins_by_tasks,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
	} `json:"oper_solo,omitempty"`
	// Player's statistics for playing the Operation mode, hard difficulty, in a Division..
	// An extra field.
	OperDivHard struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories for completing missions. The key is a number of completed missions, and the value is a number of victories.
		WinsByTasks map[string]string `json:"wins_by_tasks,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
	} `json:"oper_div_hard,omitempty"`
	// Player's statistics for playing the Operation mode, normal difficulty, in a Division..
	// An extra field.
	OperDiv struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories for completing missions. The key is a number of completed missions, and the value is a number of victories.
		WinsByTasks map[string]string `json:"wins_by_tasks,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
	} `json:"oper_div,omitempty"`
	// Last battle time
	LastBattleTime wgnTime.UnixTime `json:"last_battle_time,omitempty"`
	// Miles travelled
	Distance int `json:"distance,omitempty"`
	// Statistics in Team battles.
	// An extra field.
	Club struct {
		// Total Experience Points , including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro int `json:"torpedo_agro,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints int `json:"team_capture_points,omitempty"`
		// Victories in battles survived
		SurvivedWins int `json:"survived_wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted int `json:"ships_spotted,omitempty"`
		// Secondary armament firing statistics
		SecondBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"second_battery,omitempty"`
		// Statistics of ramming enemy warships
		Ramming struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"ramming,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled int `json:"planes_killed,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp int `json:"max_xp,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro int `json:"max_total_agro,omitempty"`
		// Most ships detected
		MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle int `json:"max_frags_battle,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
		// Main battery firing statistics
		MainBattery struct {
			// Shots fired
			Shots int `json:"shots,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Hits
			Hits int `json:"hits,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"main_battery,omitempty"`
		// Defeats
		Losses int `json:"losses,omitempty"`
		// Warships destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Draws
		Draws int `json:"draws,omitempty"`
		// Damage caused by allies
		DamageScouting int `json:"damage_scouting,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Potential damage caused by shells
		ArtAgro int `json:"art_agro,omitempty"`
		// Statistics of aircraft used
		Aircraft struct {
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
		} `json:"aircraft,omitempty"`
	} `json:"club,omitempty"`
	// Battles fought
	Battles int `json:"battles,omitempty"`
	// User ID
	AccountId int `json:"account_id,omitempty"`
} 
