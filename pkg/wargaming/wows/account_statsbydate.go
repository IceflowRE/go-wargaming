package wows

type AccountStatsbydate struct {
	// Statistics in Random Battles
	Pvp struct {
		// Total Experience Points, including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled int `json:"planes_killed,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp int `json:"max_xp,omitempty"`
		// Number of vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Date in the format "%Y%m%d"
		Date string `json:"date,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Battle type
		BattleType string `json:"battle_type,omitempty"`
		// Player account ID
		AccountId int `json:"account_id,omitempty"`
	} `json:"pvp,omitempty"`
	// Statistics in Co-op Battles.
	// An extra field.
	Pve struct {
		// Total Experience Points, including XP earned with Premium Account
		Xp int `json:"xp,omitempty"`
		// Victories
		Wins int `json:"wins,omitempty"`
		// Battles survived
		SurvivedBattles int `json:"survived_battles,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled int `json:"planes_killed,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp int `json:"max_xp,omitempty"`
		// Number of vehicles destroyed
		Frags int `json:"frags,omitempty"`
		// Base defense points
		DroppedCapturePoints int `json:"dropped_capture_points,omitempty"`
		// Date in the format "%Y%m%d"
		Date string `json:"date,omitempty"`
		// Damage caused
		DamageDealt int `json:"damage_dealt,omitempty"`
		// Base capture points
		CapturePoints int `json:"capture_points,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
		// Battle type
		BattleType string `json:"battle_type,omitempty"`
		// Player account ID
		AccountId int `json:"account_id,omitempty"`
	} `json:"pve,omitempty"`
} 
