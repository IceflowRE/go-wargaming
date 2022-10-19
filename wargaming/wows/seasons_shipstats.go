// Auto generated file!

package wows

type SeasonsShipstatsOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
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
	// Season ID. Maximum limit: 100.
	SeasonId []int `json:"season_id,omitempty"`
	// Ship ID. Maximum limit: 100.
	ShipId []int `json:"ship_id,omitempty"`
}

type SeasonsShipstats struct {
	// User ID
	AccountId *int `json:"account_id,omitempty"`
	// Ranked Battles seasons
	Seasons *struct {
		// Player statistics in Ranked Battles in division of 2 players
		RankDiv2 *struct {
			// Statistics of aircraft used
			Aircraft *struct {
				// Warships destroyed
				Frags *int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			} `json:"aircraft,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
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
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
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
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins *int `json:"survived_wins,omitempty"`
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
		// Player statistics in Ranked Battles in division of 3 players
		RankDiv3 *struct {
			// Statistics of aircraft used
			Aircraft *struct {
				// Warships destroyed
				Frags *int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			} `json:"aircraft,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
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
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
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
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins *int `json:"survived_wins,omitempty"`
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
		// Player statistics in Ranked Battles in solo game
		RankSolo *struct {
			// Statistics of aircraft used
			Aircraft *struct {
				// Warships destroyed
				Frags *int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			} `json:"aircraft,omitempty"`
			// Battles fought
			Battles *int `json:"battles,omitempty"`
			// Damage caused
			DamageDealt *int `json:"damage_dealt,omitempty"`
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
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
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
			// Battles survived
			SurvivedBattles *int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins *int `json:"survived_wins,omitempty"`
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
	} `json:"seasons,omitempty"`
	// Ship ID
	ShipId *int `json:"ship_id,omitempty"`
}
