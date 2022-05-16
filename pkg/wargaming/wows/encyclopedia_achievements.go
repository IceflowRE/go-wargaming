package wows

type EncyclopediaAchievements struct {
	// Battle achievements
	Battle struct {
		// Type of achievement
		Type_ string `json:"type,omitempty"`
		// Subtype of achievement
		SubType string `json:"sub_type,omitempty"`
		// Indicates if a reward is granted for achievement
		Reward bool `json:"reward,omitempty"`
		// Achievement name
		Name string `json:"name,omitempty"`
		// Achievement that can be received multiple times.
		Multiple int `json:"multiple,omitempty"`
		// Minimum tier of ship to receive this achievement
		MinShipLevel int `json:"min_ship_level,omitempty"`
		// Maximum tier of ship to receive this achievement
		MaxShipLevel int `json:"max_ship_level,omitempty"`
		// Maximum progress
		MaxProgress int `json:"max_progress,omitempty"`
		// Indicates if achievement is in progress
		IsProgress int `json:"is_progress,omitempty"`
		// Image of an unearned achievement
		ImageInactive string `json:"image_inactive,omitempty"`
		// Image link
		Image string `json:"image,omitempty"`
		// Achievement unavailable
		Hidden int `json:"hidden,omitempty"`
		// Achievement description
		Description string `json:"description,omitempty"`
		// Indicates how many times achievement can be obtained per battle
		CountPerBattle string `json:"count_per_battle,omitempty"`
		// Battle types in which players can receive achievements. Battle types according to the Battle Types method response
		BattleTypes []string `json:"battle_types,omitempty"`
		// Achievement ID
		AchievementId string `json:"achievement_id,omitempty"`
	} `json:"battle,omitempty"`
} 
