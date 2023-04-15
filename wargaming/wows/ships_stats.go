// Auto generated file!

package wows

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type ShipsStatsOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
	// Extra fields that will be added to the response. Valid values:
	//
	// "club"
	// "oper_div"
	// "oper_div_hard"
	// "oper_solo"
	// "pve"
	// "pve_div2"
	// "pve_div3"
	// "pve_solo"
	// "pvp_div2"
	// "pvp_div3"
	// "pvp_solo"
	// "rank_div2"
	// "rank_div3"
	// "rank_solo"
	Extra []string `json:"extra,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Filter by ship availability in the Port. If the parameter is not specified, all ships are returned. Parameter processing requires a valid access_token for the specified account_id. Valid values:
	//
	// "1" - Return ships available in the Port.
	// "0" - Return ships that are no longer in the Port.
	InGarage *string `json:"in_garage,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "cs" - Čeština
	// "de" - Deutsch
	// "en" - English (by default)
	// "es" - Español
	// "fr" - Français
	// "ja" - 日本語
	// "pl" - Polski
	// "ru" - Русский
	// "th" - ไทย
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "zh-cn" - 中文
	// "pt-br" - Português do Brasil
	// "es-mx" - Español (México)
	Language *string `json:"language,omitempty"`
	// Player's ship ID. Maximum limit: 100.
	ShipId []int `json:"ship_id,omitempty"`
}

type ShipsStats struct {
	// User ID
	AccountId *int `json:"account_id,omitempty"`
	// Battles fought
	Battles *int `json:"battles,omitempty"`
	// Statistics in Team battles.
	// An extra field.
	Club *struct {
		// Statistics of aircraft used
		Aircraft *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"aircraft,omitempty"`
		// Potential damage caused by shells
		ArtAgro *int `json:"art_agro,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Base capture points
		CapturePoints *int `json:"capture_points,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage caused by allies
		DamageScouting *int `json:"damage_scouting,omitempty"`
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
		// Warships destroyed
		Frags *int `json:"frags,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Main battery firing statistics
		MainBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"main_battery,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt *int `json:"max_damage_dealt,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting *int `json:"max_damage_scouting,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
		// Most ships detected
		MaxShipsSpotted *int `json:"max_ships_spotted,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro *int `json:"max_total_agro,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Statistics of ramming enemy warships
		Ramming *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"ramming,omitempty"`
		// Secondary armament firing statistics
		SecondBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"second_battery,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted *int `json:"ships_spotted,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints *int `json:"team_capture_points,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints *int `json:"team_dropped_capture_points,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro *int `json:"torpedo_agro,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"club,omitempty"`
	// Miles travelled
	Distance *int `json:"distance,omitempty"`
	// Last battle time
	LastBattleTime *wgnTime.UnixTime `json:"last_battle_time,omitempty"`
	// Player's statistics for playing the Operation mode, normal difficulty, in a Division..
	// An extra field.
	OperDiv *struct {
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Victories for completing missions. The key is a number of completed missions, and the value is a number of victories.
		WinsByTasks map[string]string `json:"wins_by_tasks,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"oper_div,omitempty"`
	// Player's statistics for playing the Operation mode, hard difficulty, in a Division..
	// An extra field.
	OperDivHard *struct {
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Victories for completing missions. The key is a number of completed missions, and the value is a number of victories.
		WinsByTasks map[string]string `json:"wins_by_tasks,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"oper_div_hard,omitempty"`
	// Player's statistics for playing solo the Operation mode, normal difficulty..
	// An extra field.
	OperSolo *struct {
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Victories for completing missions. The key is a number of completed missions, and the value is a number of victories.
		WinsByTasks map[string]string `json:"wins_by_tasks,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"oper_solo,omitempty"`
	// Private data on the player's ships
	Private *struct {
		// Overall battle life time in seconds
		BattleLifeTime *int `json:"battle_life_time,omitempty"`
		// Availability of ships in the Port
		InGarage *bool `json:"in_garage,omitempty"`
	} `json:"private,omitempty"`
	// Player statistics in all Co-op Battles.
	// An extra field.
	Pve *struct {
		// Statistics of aircraft used
		Aircraft *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"aircraft,omitempty"`
		// Potential damage caused by shells
		ArtAgro *int `json:"art_agro,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Base capture points
		CapturePoints *int `json:"capture_points,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage caused by allies
		DamageScouting *int `json:"damage_scouting,omitempty"`
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
		// Warships destroyed
		Frags *int `json:"frags,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Main battery firing statistics
		MainBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"main_battery,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt *int `json:"max_damage_dealt,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting *int `json:"max_damage_scouting,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
		// Most ships detected
		MaxShipsSpotted *int `json:"max_ships_spotted,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro *int `json:"max_total_agro,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Statistics of ramming enemy warships
		Ramming *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"ramming,omitempty"`
		// Secondary armament firing statistics
		SecondBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"second_battery,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted *int `json:"ships_spotted,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints *int `json:"team_capture_points,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints *int `json:"team_dropped_capture_points,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro *int `json:"torpedo_agro,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"pve,omitempty"`
	// Player statistics in Co-op Battles in Division of 2 players.
	// An extra field.
	PveDiv2 *struct {
		// Statistics of aircraft used
		Aircraft *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"aircraft,omitempty"`
		// Potential damage caused by shells
		ArtAgro *int `json:"art_agro,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Base capture points
		CapturePoints *int `json:"capture_points,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage caused by allies
		DamageScouting *int `json:"damage_scouting,omitempty"`
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
		// Warships destroyed
		Frags *int `json:"frags,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Main battery firing statistics
		MainBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"main_battery,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt *int `json:"max_damage_dealt,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting *int `json:"max_damage_scouting,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
		// Most ships detected
		MaxShipsSpotted *int `json:"max_ships_spotted,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro *int `json:"max_total_agro,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Statistics of ramming enemy warships
		Ramming *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"ramming,omitempty"`
		// Secondary armament firing statistics
		SecondBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"second_battery,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted *int `json:"ships_spotted,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints *int `json:"team_capture_points,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints *int `json:"team_dropped_capture_points,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro *int `json:"torpedo_agro,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"pve_div2,omitempty"`
	// Player statistics in Co-op Battles in Division of 3 players.
	// An extra field.
	PveDiv3 *struct {
		// Statistics of aircraft used
		Aircraft *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"aircraft,omitempty"`
		// Potential damage caused by shells
		ArtAgro *int `json:"art_agro,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Base capture points
		CapturePoints *int `json:"capture_points,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage caused by allies
		DamageScouting *int `json:"damage_scouting,omitempty"`
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
		// Warships destroyed
		Frags *int `json:"frags,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Main battery firing statistics
		MainBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"main_battery,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt *int `json:"max_damage_dealt,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting *int `json:"max_damage_scouting,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
		// Most ships detected
		MaxShipsSpotted *int `json:"max_ships_spotted,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro *int `json:"max_total_agro,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Statistics of ramming enemy warships
		Ramming *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"ramming,omitempty"`
		// Secondary armament firing statistics
		SecondBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"second_battery,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted *int `json:"ships_spotted,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints *int `json:"team_capture_points,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints *int `json:"team_dropped_capture_points,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro *int `json:"torpedo_agro,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"pve_div3,omitempty"`
	// Player statistics in solo battles fought in Co-op mode.
	// An extra field.
	PveSolo *struct {
		// Statistics of aircraft used
		Aircraft *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"aircraft,omitempty"`
		// Potential damage caused by shells
		ArtAgro *int `json:"art_agro,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Base capture points
		CapturePoints *int `json:"capture_points,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage caused by allies
		DamageScouting *int `json:"damage_scouting,omitempty"`
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
		// Warships destroyed
		Frags *int `json:"frags,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Main battery firing statistics
		MainBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"main_battery,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt *int `json:"max_damage_dealt,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting *int `json:"max_damage_scouting,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
		// Most ships detected
		MaxShipsSpotted *int `json:"max_ships_spotted,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro *int `json:"max_total_agro,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Statistics of ramming enemy warships
		Ramming *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"ramming,omitempty"`
		// Secondary armament firing statistics
		SecondBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"second_battery,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted *int `json:"ships_spotted,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints *int `json:"team_capture_points,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints *int `json:"team_dropped_capture_points,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro *int `json:"torpedo_agro,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"pve_solo,omitempty"`
	// Player statistics in all Random Battles
	Pvp *struct {
		// Statistics of aircraft used
		Aircraft *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"aircraft,omitempty"`
		// Potential damage caused by shells
		ArtAgro *int `json:"art_agro,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Number of battles, starting from version 5.10
		BattlesSince510 *int `json:"battles_since_510,omitempty"`
		// Number of battles, starting from version 5.12
		BattlesSince512 *int `json:"battles_since_512,omitempty"`
		// Base capture points
		CapturePoints *int `json:"capture_points,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage caused by allies
		DamageScouting *int `json:"damage_scouting,omitempty"`
		// Damage dealt to installations, starting from version 5.12
		DamageToBuildings *int `json:"damage_to_buildings,omitempty"`
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
		// Warships destroyed
		Frags *int `json:"frags,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Main battery firing statistics
		MainBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"main_battery,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt *int `json:"max_damage_dealt,omitempty"`
		// Most damage caused to installations in a battle
		MaxDamageDealtToBuildings *int `json:"max_damage_dealt_to_buildings,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting *int `json:"max_damage_scouting,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
		// Most ships detected
		MaxShipsSpotted *int `json:"max_ships_spotted,omitempty"`
		// Most installations surpressed in a battle
		MaxSuppressionsCount *int `json:"max_suppressions_count,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro *int `json:"max_total_agro,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Statistics of ramming enemy warships
		Ramming *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"ramming,omitempty"`
		// Secondary armament firing statistics
		SecondBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"second_battery,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted *int `json:"ships_spotted,omitempty"`
		// Number of installations surpressed, starting from version 5.12
		SuppressionsCount *int `json:"suppressions_count,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints *int `json:"team_capture_points,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints *int `json:"team_dropped_capture_points,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro *int `json:"torpedo_agro,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"pvp,omitempty"`
	// Player statistics in Random Battles in Division of 2 players.
	// An extra field.
	PvpDiv2 *struct {
		// Statistics of aircraft used
		Aircraft *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"aircraft,omitempty"`
		// Potential damage caused by shells
		ArtAgro *int `json:"art_agro,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Number of battles, starting from version 5.10
		BattlesSince510 *int `json:"battles_since_510,omitempty"`
		// Number of battles, starting from version 5.12
		BattlesSince512 *int `json:"battles_since_512,omitempty"`
		// Base capture points
		CapturePoints *int `json:"capture_points,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage caused by allies
		DamageScouting *int `json:"damage_scouting,omitempty"`
		// Damage dealt to installations, starting from version 5.12
		DamageToBuildings *int `json:"damage_to_buildings,omitempty"`
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
		// Warships destroyed
		Frags *int `json:"frags,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Main battery firing statistics
		MainBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"main_battery,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt *int `json:"max_damage_dealt,omitempty"`
		// Most damage caused to installations in a battle
		MaxDamageDealtToBuildings *int `json:"max_damage_dealt_to_buildings,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting *int `json:"max_damage_scouting,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
		// Most ships detected
		MaxShipsSpotted *int `json:"max_ships_spotted,omitempty"`
		// Most installations surpressed in a battle
		MaxSuppressionsCount *int `json:"max_suppressions_count,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro *int `json:"max_total_agro,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Statistics of ramming enemy warships
		Ramming *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"ramming,omitempty"`
		// Secondary armament firing statistics
		SecondBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"second_battery,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted *int `json:"ships_spotted,omitempty"`
		// Number of installations surpressed, starting from version 5.12
		SuppressionsCount *int `json:"suppressions_count,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints *int `json:"team_capture_points,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints *int `json:"team_dropped_capture_points,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro *int `json:"torpedo_agro,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"pvp_div2,omitempty"`
	// Player statistics in Random Battles in Division of 3 players.
	// An extra field.
	PvpDiv3 *struct {
		// Statistics of aircraft used
		Aircraft *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"aircraft,omitempty"`
		// Potential damage caused by shells
		ArtAgro *int `json:"art_agro,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Number of battles, starting from version 5.10
		BattlesSince510 *int `json:"battles_since_510,omitempty"`
		// Number of battles, starting from version 5.12
		BattlesSince512 *int `json:"battles_since_512,omitempty"`
		// Base capture points
		CapturePoints *int `json:"capture_points,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage caused by allies
		DamageScouting *int `json:"damage_scouting,omitempty"`
		// Damage dealt to installations, starting from version 5.12
		DamageToBuildings *int `json:"damage_to_buildings,omitempty"`
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
		// Warships destroyed
		Frags *int `json:"frags,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Main battery firing statistics
		MainBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"main_battery,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt *int `json:"max_damage_dealt,omitempty"`
		// Most damage caused to installations in a battle
		MaxDamageDealtToBuildings *int `json:"max_damage_dealt_to_buildings,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting *int `json:"max_damage_scouting,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
		// Most ships detected
		MaxShipsSpotted *int `json:"max_ships_spotted,omitempty"`
		// Most installations surpressed in a battle
		MaxSuppressionsCount *int `json:"max_suppressions_count,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro *int `json:"max_total_agro,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Statistics of ramming enemy warships
		Ramming *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"ramming,omitempty"`
		// Secondary armament firing statistics
		SecondBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"second_battery,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted *int `json:"ships_spotted,omitempty"`
		// Number of installations surpressed, starting from version 5.12
		SuppressionsCount *int `json:"suppressions_count,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints *int `json:"team_capture_points,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints *int `json:"team_dropped_capture_points,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro *int `json:"torpedo_agro,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"pvp_div3,omitempty"`
	// Player statistics in solo battles fought in Random mode.
	// An extra field.
	PvpSolo *struct {
		// Statistics of aircraft used
		Aircraft *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"aircraft,omitempty"`
		// Potential damage caused by shells
		ArtAgro *int `json:"art_agro,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Number of battles, starting from version 5.10
		BattlesSince510 *int `json:"battles_since_510,omitempty"`
		// Number of battles, starting from version 5.12
		BattlesSince512 *int `json:"battles_since_512,omitempty"`
		// Base capture points
		CapturePoints *int `json:"capture_points,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage caused by allies
		DamageScouting *int `json:"damage_scouting,omitempty"`
		// Damage dealt to installations, starting from version 5.12
		DamageToBuildings *int `json:"damage_to_buildings,omitempty"`
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
		// Warships destroyed
		Frags *int `json:"frags,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Main battery firing statistics
		MainBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"main_battery,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt *int `json:"max_damage_dealt,omitempty"`
		// Most damage caused to installations in a battle
		MaxDamageDealtToBuildings *int `json:"max_damage_dealt_to_buildings,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting *int `json:"max_damage_scouting,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
		// Most ships detected
		MaxShipsSpotted *int `json:"max_ships_spotted,omitempty"`
		// Most installations surpressed in a battle
		MaxSuppressionsCount *int `json:"max_suppressions_count,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro *int `json:"max_total_agro,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Statistics of ramming enemy warships
		Ramming *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"ramming,omitempty"`
		// Secondary armament firing statistics
		SecondBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"second_battery,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted *int `json:"ships_spotted,omitempty"`
		// Number of installations surpressed, starting from version 5.12
		SuppressionsCount *int `json:"suppressions_count,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints *int `json:"team_capture_points,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints *int `json:"team_dropped_capture_points,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro *int `json:"torpedo_agro,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"pvp_solo,omitempty"`
	// Player statistics in Ranked Battles in division of 2 players.
	// An extra field.
	RankDiv2 *struct {
		// Statistics of aircraft used
		Aircraft *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"aircraft,omitempty"`
		// Potential damage caused by shells
		ArtAgro *int `json:"art_agro,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage caused by allies
		DamageScouting *int `json:"damage_scouting,omitempty"`
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Warships destroyed
		Frags *int `json:"frags,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Main battery firing statistics
		MainBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"main_battery,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt *int `json:"max_damage_dealt,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting *int `json:"max_damage_scouting,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
		// Most ships detected
		MaxShipsSpotted *int `json:"max_ships_spotted,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro *int `json:"max_total_agro,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Statistics of ramming enemy warships
		Ramming *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"ramming,omitempty"`
		// Secondary armament firing statistics
		SecondBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"second_battery,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted *int `json:"ships_spotted,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints *int `json:"team_capture_points,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints *int `json:"team_dropped_capture_points,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro *int `json:"torpedo_agro,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"rank_div2,omitempty"`
	// Player statistics in Ranked Battles in division of 3 players.
	// An extra field.
	RankDiv3 *struct {
		// Statistics of aircraft used
		Aircraft *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"aircraft,omitempty"`
		// Potential damage caused by shells
		ArtAgro *int `json:"art_agro,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage caused by allies
		DamageScouting *int `json:"damage_scouting,omitempty"`
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Warships destroyed
		Frags *int `json:"frags,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Main battery firing statistics
		MainBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"main_battery,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt *int `json:"max_damage_dealt,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting *int `json:"max_damage_scouting,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
		// Most ships detected
		MaxShipsSpotted *int `json:"max_ships_spotted,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro *int `json:"max_total_agro,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Statistics of ramming enemy warships
		Ramming *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"ramming,omitempty"`
		// Secondary armament firing statistics
		SecondBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"second_battery,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted *int `json:"ships_spotted,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints *int `json:"team_capture_points,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints *int `json:"team_dropped_capture_points,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro *int `json:"torpedo_agro,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"rank_div3,omitempty"`
	// Player statistics in Ranked Battles in solo game.
	// An extra field.
	RankSolo *struct {
		// Statistics of aircraft used
		Aircraft *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"aircraft,omitempty"`
		// Potential damage caused by shells
		ArtAgro *int `json:"art_agro,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage caused by allies
		DamageScouting *int `json:"damage_scouting,omitempty"`
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Warships destroyed
		Frags *int `json:"frags,omitempty"`
		// Defeats
		Losses *int `json:"losses,omitempty"`
		// Main battery firing statistics
		MainBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"main_battery,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt *int `json:"max_damage_dealt,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting *int `json:"max_damage_scouting,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
		// Most ships detected
		MaxShipsSpotted *int `json:"max_ships_spotted,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro *int `json:"max_total_agro,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Statistics of ramming enemy warships
		Ramming *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		} `json:"ramming,omitempty"`
		// Secondary armament firing statistics
		SecondBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"second_battery,omitempty"`
		// Ships spotted by the player first
		ShipsSpotted *int `json:"ships_spotted,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories in battles survived
		SurvivedWins *int `json:"survived_wins,omitempty"`
		// Total number of base capture points earned by the player's team
		TeamCapturePoints *int `json:"team_capture_points,omitempty"`
		// Total number of base defense points earned by the player's team
		TeamDroppedCapturePoints *int `json:"team_dropped_capture_points,omitempty"`
		// Potential damage caused by torpedoes
		TorpedoAgro *int `json:"torpedo_agro,omitempty"`
		// Statistics of attacking targets with torpedoes
		Torpedoes *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"rank_solo,omitempty"`
	// Ship ID
	ShipId *int `json:"ship_id,omitempty"`
	// Date when details on ships were updated
	UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
}
