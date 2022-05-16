package wows

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type AccountInfo struct {
	// Date when player details were updated
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Date when stats for player and player's ships were updated
	StatsUpdatedAt wgnTime.UnixTime `json:"stats_updated_at,omitempty"`
	// Player statistics
	Statistics struct {
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"ramming,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Shots fired
				Shots int `json:"shots,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"ramming,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Shots fired
				Shots int `json:"shots,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"ramming,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Shots fired
				Shots int `json:"shots,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"ramming,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that surpressed the most installations
			MaxSuppressionsShipId int `json:"max_suppressions_ship_id,omitempty"`
			// Most installations surpressed in a battle
			MaxSuppressionsCount int `json:"max_suppressions_count,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Ship that caused the most damage to installations
			MaxDamageDealtToBuildingsShipId int `json:"max_damage_dealt_to_buildings_ship_id,omitempty"`
			// Most damage caused to installations in a battle
			MaxDamageDealtToBuildings int `json:"max_damage_dealt_to_buildings,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Shots fired
				Shots int `json:"shots,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Damage dealt to installations, starting from version 5.12
			DamageToBuildings int `json:"damage_to_buildings,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Total number of capture points
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"ramming,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that surpressed the most installations
			MaxSuppressionsShipId int `json:"max_suppressions_ship_id,omitempty"`
			// Most installations surpressed in a battle
			MaxSuppressionsCount int `json:"max_suppressions_count,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Ship that caused the most damage to installations
			MaxDamageDealtToBuildingsShipId int `json:"max_damage_dealt_to_buildings_ship_id,omitempty"`
			// Most damage caused to installations in a battle
			MaxDamageDealtToBuildings int `json:"max_damage_dealt_to_buildings,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Shots fired
				Shots int `json:"shots,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Damage dealt to installations, starting from version 5.12
			DamageToBuildings int `json:"damage_to_buildings,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Total number of capture points
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"ramming,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that surpressed the most installations
			MaxSuppressionsShipId int `json:"max_suppressions_ship_id,omitempty"`
			// Most installations surpressed in a battle
			MaxSuppressionsCount int `json:"max_suppressions_count,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Ship that caused the most damage to installations
			MaxDamageDealtToBuildingsShipId int `json:"max_damage_dealt_to_buildings_ship_id,omitempty"`
			// Most damage caused to installations in a battle
			MaxDamageDealtToBuildings int `json:"max_damage_dealt_to_buildings,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Shots fired
				Shots int `json:"shots,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Damage dealt to installations, starting from version 5.12
			DamageToBuildings int `json:"damage_to_buildings,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Total number of capture points
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"ramming,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that surpressed the most installations
			MaxSuppressionsShipId int `json:"max_suppressions_ship_id,omitempty"`
			// Most installations surpressed in a battle
			MaxSuppressionsCount int `json:"max_suppressions_count,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Ship that caused the most damage to installations
			MaxDamageDealtToBuildingsShipId int `json:"max_damage_dealt_to_buildings_ship_id,omitempty"`
			// Most damage caused to installations in a battle
			MaxDamageDealtToBuildings int `json:"max_damage_dealt_to_buildings,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Shots fired
				Shots int `json:"shots,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Damage dealt to installations, starting from version 5.12
			DamageToBuildings int `json:"damage_to_buildings,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Total number of capture points
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"ramming,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Shots fired
				Shots int `json:"shots,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Total number of capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"ramming,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Shots fired
				Shots int `json:"shots,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Total number of capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"ramming,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Shots fired
				Shots int `json:"shots,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Total number of capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"ramming,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Shots fired
				Shots int `json:"shots,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Total number of capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"aircraft,omitempty"`
		} `json:"pve,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"ramming,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Shots fired
				Shots int `json:"shots,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
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
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Total number of capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"aircraft,omitempty"`
		} `json:"club,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
	} `json:"statistics,omitempty"`
	// Player's private data
	Private struct {
		// WoWS premium expiration time
		WowsPremiumExpiresAt int `json:"wows_premium_expires_at,omitempty"`
		// Number of slots in the Port
		Slots int `json:"slots,omitempty"`
		// Premium Account expiration time
		PremiumExpiresAt int `json:"premium_expires_at,omitempty"`
		// Ships in Port.
		// An extra field.
		Port []int `json:"port,omitempty"`
		// Contact groups.
		// An extra field.
		GroupedContacts struct {
			// Not grouped
			Ungrouped []int `json:"ungrouped,omitempty"`
			// Groups
			Groups map[string]string `json:"groups,omitempty"`
			// Blocked
			Blocked []int `json:"blocked,omitempty"`
		} `json:"grouped_contacts,omitempty"`
		// Gold
		Gold int `json:"gold,omitempty"`
		// Free Experience
		FreeXp int `json:"free_xp,omitempty"`
		// Number of slots available in the Port
		EmptySlots int `json:"empty_slots,omitempty"`
		// Credits
		Credits int `json:"credits,omitempty"`
		// Overall battle life time in seconds
		BattleLifeTime int `json:"battle_life_time,omitempty"`
	} `json:"private,omitempty"`
	// Player name
	Nickname string `json:"nickname,omitempty"`
	// End time of last game session
	LogoutAt wgnTime.UnixTime `json:"logout_at,omitempty"`
	// Service Record points
	LevelingTier int `json:"leveling_tier,omitempty"`
	// Service Record level
	LevelingPoints int `json:"leveling_points,omitempty"`
	// Last battle time
	LastBattleTime wgnTime.UnixTime `json:"last_battle_time,omitempty"`
	// Player's karma
	Karma int `json:"karma,omitempty"`
	// Indicates if the game profile is hidden
	HiddenProfile bool `json:"hidden_profile,omitempty"`
	// Date when player's account was created
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// User ID
	AccountId int `json:"account_id,omitempty"`
} 
