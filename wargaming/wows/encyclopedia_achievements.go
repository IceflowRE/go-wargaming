// Auto generated file!

package wows

type EncyclopediaAchievementsOptions struct {
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

type EncyclopediaAchievements struct {
	// Battle achievements
	Battle *struct {
		// Achievement ID
		AchievementId *string `json:"achievement_id,omitempty"`
		// Battle types in which players can receive achievements. Battle types according to the Battle Types method response
		BattleTypes []string `json:"battle_types,omitempty"`
		// Indicates how many times achievement can be obtained per battle
		CountPerBattle *string `json:"count_per_battle,omitempty"`
		// Achievement description
		Description *string `json:"description,omitempty"`
		// Achievement unavailable
		Hidden *int `json:"hidden,omitempty"`
		// Image link
		Image *string `json:"image,omitempty"`
		// Image of an unearned achievement
		ImageInactive *string `json:"image_inactive,omitempty"`
		// Indicates if achievement is in progress
		IsProgress *int `json:"is_progress,omitempty"`
		// Maximum progress
		MaxProgress *int `json:"max_progress,omitempty"`
		// Maximum tier of ship to receive this achievement
		MaxShipLevel *int `json:"max_ship_level,omitempty"`
		// Minimum tier of ship to receive this achievement
		MinShipLevel *int `json:"min_ship_level,omitempty"`
		// Achievement that can be received multiple times.
		Multiple *int `json:"multiple,omitempty"`
		// Achievement name
		Name *string `json:"name,omitempty"`
		// Indicates if a reward is granted for achievement
		Reward *bool `json:"reward,omitempty"`
		// Subtype of achievement
		SubType *string `json:"sub_type,omitempty"`
		// Type of achievement
		Type *string `json:"type,omitempty"`
	} `json:"battle,omitempty"`
}
