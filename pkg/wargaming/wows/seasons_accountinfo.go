package wows

type SeasonsAccountinfo struct {
	// Ranked Battles seasons
	Seasons struct {
		// Player statistics in Ranked Battles in solo game
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
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
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
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
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
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"aircraft,omitempty"`
		} `json:"rank_solo,omitempty"`
		// Rank Battles information
		RankInfo struct {
			// Rank required to participate in a season
			StartRank int `json:"start_rank,omitempty"`
			// Number of stars
			Stars int `json:"stars,omitempty"`
			// Current season status
			Stage int `json:"stage,omitempty"`
			// Player's current rank
			Rank int `json:"rank,omitempty"`
			// Player's best rank in current season
			MaxRank int `json:"max_rank,omitempty"`
		} `json:"rank_info,omitempty"`
		// Player statistics in Ranked Battles in division of 3 players
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
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
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
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
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
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"aircraft,omitempty"`
		} `json:"rank_div3,omitempty"`
		// Player statistics in Ranked Battles in division of 2 players
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
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
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
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
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
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
			} `json:"aircraft,omitempty"`
		} `json:"rank_div2,omitempty"`
	} `json:"seasons,omitempty"`
	// User ID
	AccountId int `json:"account_id,omitempty"`
} 
