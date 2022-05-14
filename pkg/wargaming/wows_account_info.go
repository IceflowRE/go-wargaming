package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WowsAccountInfo struct {
	// User ID
	AccountId int `json:"account_id,omitempty"`
	// Date when player's account was created
	CreatedAt UnixTime `json:"created_at,omitempty"`
	// Indicates if the game profile is hidden
	HiddenProfile bool `json:"hidden_profile,omitempty"`
	// Player's karma
	Karma int `json:"karma,omitempty"`
	// Last battle time
	LastBattleTime UnixTime `json:"last_battle_time,omitempty"`
	// Service Record level
	LevelingPoints int `json:"leveling_points,omitempty"`
	// Service Record points
	LevelingTier int `json:"leveling_tier,omitempty"`
	// End time of last game session
	LogoutAt UnixTime `json:"logout_at,omitempty"`
	// Player name
	Nickname string `json:"nickname,omitempty"`
	// Date when stats for player and player's ships were updated
	StatsUpdatedAt UnixTime `json:"stats_updated_at,omitempty"`
	// Date when player details were updated
	UpdatedAt UnixTime `json:"updated_at,omitempty"`
	// Player's private data
	Private struct {
		// Overall battle life time in seconds
		BattleLifeTime int `json:"battle_life_time,omitempty"`
		// Credits
		Credits int `json:"credits,omitempty"`
		// Number of slots available in the Port
		EmptySlots int `json:"empty_slots,omitempty"`
		// Free Experience
		FreeXp int `json:"free_xp,omitempty"`
		// Gold
		Gold int `json:"gold,omitempty"`
		// Ships in Port.
		// An extra field.
		Port []int `json:"port,omitempty"`
		// Premium Account expiration time
		PremiumExpiresAt int `json:"premium_expires_at,omitempty"`
		// Number of slots in the Port
		Slots int `json:"slots,omitempty"`
		// WoWS premium expiration time
		WowsPremiumExpiresAt int `json:"wows_premium_expires_at,omitempty"`
		// Contact groups.
		// An extra field.
		GroupedContacts struct {
			// Blocked
			Blocked []int `json:"blocked,omitempty"`
			// Groups
			Groups map[string]string `json:"groups,omitempty"`
			// Not grouped
			Ungrouped []int `json:"ungrouped,omitempty"`
		} `json:"grouped_contacts,omitempty"`
	} `json:"private,omitempty"`
	// Player statistics
	Statistics struct {
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Miles travelled
		Distance int `json:"distance,omitempty"`
		// Statistics in Team battles.
		// An extra field.
		Club struct {
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Total number of capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Ships spotted by the player first
			ShipsSpotted int `json:"ships_spotted,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Total number of base capture points earned by the player's team
			TeamCapturePoints int `json:"team_capture_points,omitempty"`
			// Total number of base defense points earned by the player's team
			TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
			// Potential damage caused by torpedoes
			TorpedoAgro int `json:"torpedo_agro,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"club,omitempty"`
		// Player's statistics for playing the Operation mode, normal difficulty, in a Division..
		// An extra field.
		OperDiv struct {
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Victories for completing missions. The key is a number of completed missions, and the value is a number of victories.
			WinsByTasks map[string]string `json:"wins_by_tasks,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
		} `json:"oper_div,omitempty"`
		// Player's statistics for playing the Operation mode, hard difficulty, in a Division..
		// An extra field.
		OperDivHard struct {
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Victories for completing missions. The key is a number of completed missions, and the value is a number of victories.
			WinsByTasks map[string]string `json:"wins_by_tasks,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
		} `json:"oper_div_hard,omitempty"`
		// Player's statistics for playing solo the Operation mode, normal difficulty..
		// An extra field.
		OperSolo struct {
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Victories for completing missions. The key is a number of completed missions, and the value is a number of victories.
			WinsByTasks map[string]string `json:"wins_by_tasks,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
		} `json:"oper_solo,omitempty"`
		// Player statistics in all Co-op Battles.
		// An extra field.
		Pve struct {
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Total number of capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Ships spotted by the player first
			ShipsSpotted int `json:"ships_spotted,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Total number of base capture points earned by the player's team
			TeamCapturePoints int `json:"team_capture_points,omitempty"`
			// Total number of base defense points earned by the player's team
			TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
			// Potential damage caused by torpedoes
			TorpedoAgro int `json:"torpedo_agro,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"pve,omitempty"`
		// Player statistics in Co-op Battles in Division of 2 players.
		// An extra field.
		PveDiv2 struct {
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Total number of capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Ships spotted by the player first
			ShipsSpotted int `json:"ships_spotted,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Total number of base capture points earned by the player's team
			TeamCapturePoints int `json:"team_capture_points,omitempty"`
			// Total number of base defense points earned by the player's team
			TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
			// Potential damage caused by torpedoes
			TorpedoAgro int `json:"torpedo_agro,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"pve_div2,omitempty"`
		// Player statistics in Co-op Battles in Division of 3 players.
		// An extra field.
		PveDiv3 struct {
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Total number of capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Ships spotted by the player first
			ShipsSpotted int `json:"ships_spotted,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Total number of base capture points earned by the player's team
			TeamCapturePoints int `json:"team_capture_points,omitempty"`
			// Total number of base defense points earned by the player's team
			TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
			// Potential damage caused by torpedoes
			TorpedoAgro int `json:"torpedo_agro,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"pve_div3,omitempty"`
		// Player statistics in solo battles fought in Co-op mode.
		// An extra field.
		PveSolo struct {
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Total number of capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Ships spotted by the player first
			ShipsSpotted int `json:"ships_spotted,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Total number of base capture points earned by the player's team
			TeamCapturePoints int `json:"team_capture_points,omitempty"`
			// Total number of base defense points earned by the player's team
			TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
			// Potential damage caused by torpedoes
			TorpedoAgro int `json:"torpedo_agro,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"pve_solo,omitempty"`
		// Player statistics in all Random Battles
		Pvp struct {
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Number of battles, starting from version 5.10
			BattlesSince510 int `json:"battles_since_510,omitempty"`
			// Number of battles, starting from version 5.12
			BattlesSince512 int `json:"battles_since_512,omitempty"`
			// Total number of capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Damage dealt to installations, starting from version 5.12
			DamageToBuildings int `json:"damage_to_buildings,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Most damage caused to installations in a battle
			MaxDamageDealtToBuildings int `json:"max_damage_dealt_to_buildings,omitempty"`
			// Ship that caused the most damage to installations
			MaxDamageDealtToBuildingsShipId int `json:"max_damage_dealt_to_buildings_ship_id,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most installations surpressed in a battle
			MaxSuppressionsCount int `json:"max_suppressions_count,omitempty"`
			// Ship that surpressed the most installations
			MaxSuppressionsShipId int `json:"max_suppressions_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Ships spotted by the player first
			ShipsSpotted int `json:"ships_spotted,omitempty"`
			// Number of installations surpressed, starting from version 5.12
			SuppressionsCount int `json:"suppressions_count,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Total number of base capture points earned by the player's team
			TeamCapturePoints int `json:"team_capture_points,omitempty"`
			// Total number of base defense points earned by the player's team
			TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
			// Potential damage caused by torpedoes
			TorpedoAgro int `json:"torpedo_agro,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"pvp,omitempty"`
		// Player statistics in Random Battles in Division of 2 players.
		// An extra field.
		PvpDiv2 struct {
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Number of battles, starting from version 5.10
			BattlesSince510 int `json:"battles_since_510,omitempty"`
			// Number of battles, starting from version 5.12
			BattlesSince512 int `json:"battles_since_512,omitempty"`
			// Total number of capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Damage dealt to installations, starting from version 5.12
			DamageToBuildings int `json:"damage_to_buildings,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Most damage caused to installations in a battle
			MaxDamageDealtToBuildings int `json:"max_damage_dealt_to_buildings,omitempty"`
			// Ship that caused the most damage to installations
			MaxDamageDealtToBuildingsShipId int `json:"max_damage_dealt_to_buildings_ship_id,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most installations surpressed in a battle
			MaxSuppressionsCount int `json:"max_suppressions_count,omitempty"`
			// Ship that surpressed the most installations
			MaxSuppressionsShipId int `json:"max_suppressions_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Ships spotted by the player first
			ShipsSpotted int `json:"ships_spotted,omitempty"`
			// Number of installations surpressed, starting from version 5.12
			SuppressionsCount int `json:"suppressions_count,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Total number of base capture points earned by the player's team
			TeamCapturePoints int `json:"team_capture_points,omitempty"`
			// Total number of base defense points earned by the player's team
			TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
			// Potential damage caused by torpedoes
			TorpedoAgro int `json:"torpedo_agro,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"pvp_div2,omitempty"`
		// Player statistics in Random Battles in Division of 3 players.
		// An extra field.
		PvpDiv3 struct {
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Number of battles, starting from version 5.10
			BattlesSince510 int `json:"battles_since_510,omitempty"`
			// Number of battles, starting from version 5.12
			BattlesSince512 int `json:"battles_since_512,omitempty"`
			// Total number of capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Damage dealt to installations, starting from version 5.12
			DamageToBuildings int `json:"damage_to_buildings,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Most damage caused to installations in a battle
			MaxDamageDealtToBuildings int `json:"max_damage_dealt_to_buildings,omitempty"`
			// Ship that caused the most damage to installations
			MaxDamageDealtToBuildingsShipId int `json:"max_damage_dealt_to_buildings_ship_id,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most installations surpressed in a battle
			MaxSuppressionsCount int `json:"max_suppressions_count,omitempty"`
			// Ship that surpressed the most installations
			MaxSuppressionsShipId int `json:"max_suppressions_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Ships spotted by the player first
			ShipsSpotted int `json:"ships_spotted,omitempty"`
			// Number of installations surpressed, starting from version 5.12
			SuppressionsCount int `json:"suppressions_count,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Total number of base capture points earned by the player's team
			TeamCapturePoints int `json:"team_capture_points,omitempty"`
			// Total number of base defense points earned by the player's team
			TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
			// Potential damage caused by torpedoes
			TorpedoAgro int `json:"torpedo_agro,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"pvp_div3,omitempty"`
		// Player statistics in solo battles fought in Random mode.
		// An extra field.
		PvpSolo struct {
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Number of battles, starting from version 5.10
			BattlesSince510 int `json:"battles_since_510,omitempty"`
			// Number of battles, starting from version 5.12
			BattlesSince512 int `json:"battles_since_512,omitempty"`
			// Total number of capture points
			CapturePoints int `json:"capture_points,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Damage dealt to installations, starting from version 5.12
			DamageToBuildings int `json:"damage_to_buildings,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Total number of defense points
			DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Most damage caused to installations in a battle
			MaxDamageDealtToBuildings int `json:"max_damage_dealt_to_buildings,omitempty"`
			// Ship that caused the most damage to installations
			MaxDamageDealtToBuildingsShipId int `json:"max_damage_dealt_to_buildings_ship_id,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most installations surpressed in a battle
			MaxSuppressionsCount int `json:"max_suppressions_count,omitempty"`
			// Ship that surpressed the most installations
			MaxSuppressionsShipId int `json:"max_suppressions_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Ships spotted by the player first
			ShipsSpotted int `json:"ships_spotted,omitempty"`
			// Number of installations surpressed, starting from version 5.12
			SuppressionsCount int `json:"suppressions_count,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Total number of base capture points earned by the player's team
			TeamCapturePoints int `json:"team_capture_points,omitempty"`
			// Total number of base defense points earned by the player's team
			TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
			// Potential damage caused by torpedoes
			TorpedoAgro int `json:"torpedo_agro,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"pvp_solo,omitempty"`
		// Player statistics in Ranked Battles in division of 2 players.
		// An extra field.
		RankDiv2 struct {
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Ships spotted by the player first
			ShipsSpotted int `json:"ships_spotted,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Total number of base capture points earned by the player's team
			TeamCapturePoints int `json:"team_capture_points,omitempty"`
			// Total number of base defense points earned by the player's team
			TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
			// Potential damage caused by torpedoes
			TorpedoAgro int `json:"torpedo_agro,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"rank_div2,omitempty"`
		// Player statistics in Ranked Battles in division of 3 players.
		// An extra field.
		RankDiv3 struct {
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Ships spotted by the player first
			ShipsSpotted int `json:"ships_spotted,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Total number of base capture points earned by the player's team
			TeamCapturePoints int `json:"team_capture_points,omitempty"`
			// Total number of base defense points earned by the player's team
			TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
			// Potential damage caused by torpedoes
			TorpedoAgro int `json:"torpedo_agro,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"rank_div3,omitempty"`
		// Player statistics in Ranked Battles in solo game.
		// An extra field.
		RankSolo struct {
			// Potential damage caused by shells
			ArtAgro int `json:"art_agro,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Base capture points
			ControlCapturedPoints int `json:"control_captured_points,omitempty"`
			// Base defense points
			ControlDroppedPoints int `json:"control_dropped_points,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Damage caused by allies
			DamageScouting int `json:"damage_scouting,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Warship used to cause maximum damage
			MaxDamageDealtShipId int `json:"max_damage_dealt_ship_id,omitempty"`
			// Most damage caused by allies to enemy ships spotted by the player
			MaxDamageScouting int `json:"max_damage_scouting,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Warship with maximum number of enemy aircraft destroyed
			MaxPlanesKilledShipId int `json:"max_planes_killed_ship_id,omitempty"`
			// Ship that caused the most damage to enemy ships spotted by the player
			MaxScoutingDamageShipId int `json:"max_scouting_damage_ship_id,omitempty"`
			// Most ships detected
			MaxShipsSpotted int `json:"max_ships_spotted,omitempty"`
			// Ship that detected the most enemy ships
			MaxShipsSpottedShipId int `json:"max_ships_spotted_ship_id,omitempty"`
			// Most potential damage caused by ammunition that hit or fell near the ship
			MaxTotalAgro int `json:"max_total_agro,omitempty"`
			// Ship that caused the most potential damage
			MaxTotalAgroShipId int `json:"max_total_agro_ship_id,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
			MaxXpShipId int `json:"max_xp_ship_id,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Ships spotted by the player first
			ShipsSpotted int `json:"ships_spotted,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Total number of base capture points earned by the player's team
			TeamCapturePoints int `json:"team_capture_points,omitempty"`
			// Total number of base defense points earned by the player's team
			TeamDroppedCapturePoints int `json:"team_dropped_capture_points,omitempty"`
			// Potential damage caused by torpedoes
			TorpedoAgro int `json:"torpedo_agro,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// ID of a ship with maximum number of enemy warships destroyed per battle
				MaxFragsShipId int `json:"max_frags_ship_id,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"rank_solo,omitempty"`
	} `json:"statistics,omitempty"`
}

// WowsAccountInfo Method returns player details. Players may hide their game profiles, use field hidden_profile for determination.
//
// https://developers.wargaming.net/reference/all/wows/account/info
//
// account_id:
//     Player account ID. Maximum limit: 100. Min value is 1.
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "private.grouped_contacts" 
//     "private.port" 
//     "statistics.club" 
//     "statistics.oper_div" 
//     "statistics.oper_div_hard" 
//     "statistics.oper_solo" 
//     "statistics.pve" 
//     "statistics.pve_div2" 
//     "statistics.pve_div3" 
//     "statistics.pve_solo" 
//     "statistics.pvp_div2" 
//     "statistics.pvp_div3" 
//     "statistics.pvp_solo" 
//     "statistics.rank_div2" 
//     "statistics.rank_div3" 
//     "statistics.rank_solo"
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "cs" &mdash; Čeština 
//     "de" &mdash; Deutsch 
//     "en" &mdash; English 
//     "es" &mdash; Español 
//     "fr" &mdash; Français 
//     "ja" &mdash; 日本語 
//     "pl" &mdash; Polski 
//     "ru" &mdash; Русский (by default)
//     "th" &mdash; ไทย 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "zh-cn" &mdash; 中文 
//     "pt-br" &mdash; Português do Brasil 
//     "es-mx" &mdash; Español (México)
func (client *Client) WowsAccountInfo(realm Realm, accountId []int, accessToken string, extra []string, fields []string, language string) (*WowsAccountInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"access_token": accessToken,
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowsAccountInfo
	err := client.doGetDataRequest(realm, "/wows/account/info/", reqParam, &result)
	return result, err
}
