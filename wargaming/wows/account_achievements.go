package wows


// AccountAchievementsOptions options.
type AccountAchievementsOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Localization language. Default is "ru". Valid values:
	// 
	// "cs" - Čeština 
	// "de" - Deutsch 
	// "en" - English 
	// "es" - Español 
	// "fr" - Français 
	// "ja" - 日本語 
	// "pl" - Polski 
	// "ru" - Русский (by default)
	// "th" - ไทย 
	// "zh-tw" - 繁體中文 
	// "tr" - Türkçe 
	// "zh-cn" - 中文 
	// "pt-br" - Português do Brasil 
	// "es-mx" - Español (México)
	Language *string
}

type AccountAchievements struct {
	// Battle achievements earned
	Battle map[string]string `json:"battle,omitempty"`
	// Achievement progress
	Progress map[string]string `json:"progress,omitempty"`
}
