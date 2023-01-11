// Auto generated file!

package wows

type AccountStatsbydateOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
	// List of dates to return statistics slices for. Format: YYYYMMDD. Max. dates range - 28 days from the current date. Statistics slice for yesterday will be returned by default. Maximum limit: 10.
	Dates []string `json:"dates,omitempty"`
	// Extra fields that will be added to the response. Valid values:
	//
	// "pve"
	Extra []string `json:"extra,omitempty"`
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

type AccountStatsbydate struct {
	// Statistics in Co-op Battles.
	// An extra field.
	Pve *struct {
		// Player account ID
		AccountId *int `json:"account_id,omitempty"`
		// Battle type
		BattleType *string `json:"battle_type,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Base capture points
		CapturePoints *int `json:"capture_points,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Date in the format "%Y%m%d"
		Date *string `json:"date,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
		// Number of vehicles destroyed
		Frags *int `json:"frags,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points, including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"pve,omitempty"`
	// Statistics in Random Battles
	Pvp *struct {
		// Player account ID
		AccountId *int `json:"account_id,omitempty"`
		// Battle type
		BattleType *string `json:"battle_type,omitempty"`
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Base capture points
		CapturePoints *int `json:"capture_points,omitempty"`
		// Damage caused
		DamageDealt *int `json:"damage_dealt,omitempty"`
		// Date in the format "%Y%m%d"
		Date *string `json:"date,omitempty"`
		// Base defense points
		DroppedCapturePoints *int `json:"dropped_capture_points,omitempty"`
		// Number of vehicles destroyed
		Frags *int `json:"frags,omitempty"`
		// Maximum Experience Points per battle, including XP earned with Premium Account
		MaxXp *int `json:"max_xp,omitempty"`
		// Enemy aircraft destroyed
		PlanesKilled *int `json:"planes_killed,omitempty"`
		// Battles survived
		SurvivedBattles *int `json:"survived_battles,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
		// Total Experience Points, including XP earned with Premium Account
		Xp *int `json:"xp,omitempty"`
	} `json:"pvp,omitempty"`
}
