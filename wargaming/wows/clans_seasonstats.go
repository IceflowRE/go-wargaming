// Auto generated file!

package wows

type ClansSeasonstatsOptions struct {
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
}

type ClansSeasonstats struct {
	// User ID
	AccountId *int `json:"account_id,omitempty"`
	// Clans battles season statistics
	Seasons *struct {
		// Statistics of aircraft used
		Aircraft *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId *int `json:"max_frags_ship_id,omitempty"`
		} `json:"aircraft,omitempty"`
		// Potential damage caused by shells
		ArtAgro *int `json:"art_agro,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Total number of capture points
		CapturePoints *int `json:"capture_points,omitempty"`
		// Base capture points
		ControlCapturedPoints *int `json:"control_captured_points,omitempty"`
		// Base defense points
		ControlDroppedPoints *int `json:"control_dropped_points,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Damage caused by allies
		DamageScouting *int `json:"damage_scouting,omitempty"`
		// Draws
		Draws *int `json:"draws,omitempty"`
		// Total number of defense points
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
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId *int `json:"max_frags_ship_id,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"main_battery,omitempty"`
		// Maximum Damage caused per battle
		MaxDamageDealt *int `json:"max_damage_dealt,omitempty"`
		// Warship used to cause maximum damage
		MaxDamageDealtShipId *int `json:"max_damage_dealt_ship_id,omitempty"`
		// Most damage caused by allies to enemy ships spotted by the player
		MaxDamageScouting *int `json:"max_damage_scouting,omitempty"`
		// Maximum number of enemy warships destroyed per battle
		MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
		// ID of a ship with maximum number of enemy warships destroyed per battle
		MaxFragsShipId *int `json:"max_frags_ship_id,omitempty"`
		// Maximum number of aircraft destroyed per battle
		MaxPlanesKilled *int `json:"max_planes_killed,omitempty"`
		// Warship with maximum number of enemy aircraft destroyed
		MaxPlanesKilledShipId *int `json:"max_planes_killed_ship_id,omitempty"`
		// Ship that caused the most damage to enemy ships spotted by the player
		MaxScoutingDamageShipId *int `json:"max_scouting_damage_ship_id,omitempty"`
		// Most ships detected
		MaxShipsSpotted *int `json:"max_ships_spotted,omitempty"`
		// Ship that detected the most enemy ships
		MaxShipsSpottedShipId *int `json:"max_ships_spotted_ship_id,omitempty"`
		// Most potential damage caused by ammunition that hit or fell near the ship
		MaxTotalAgro *int `json:"max_total_agro,omitempty"`
		// Ship that caused the most potential damage
		MaxTotalAgroShipId *int `json:"max_total_agro_ship_id,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Warship used to gain maximum Experience Points per battle, including XP earned with Premium Account
		MaxXpShipId *int `json:"max_xp_ship_id,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Statistics of ramming enemy warships
		Ramming *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId *int `json:"max_frags_ship_id,omitempty"`
		} `json:"ramming,omitempty"`
		// Season ID
		SeasonId *int `json:"season_id,omitempty"`
		// Secondary armament firing statistics
		SecondBattery *struct {
			// Warships destroyed
			Frags *int `json:"frags,omitempty"`
			// Hits
			Hits *int `json:"hits,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle *int `json:"max_frags_battle,omitempty"`
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId *int `json:"max_frags_ship_id,omitempty"`
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
			// ID of a ship with maximum number of enemy warships destroyed per battle
			MaxFragsShipId *int `json:"max_frags_ship_id,omitempty"`
			// Shots fired
			Shots *int `json:"shots,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points , including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"seasons,omitempty"`
}
