package wotb

// AccountListOptions options.
type AccountListOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Localization language. Default is "ru". Valid values:
	//
	// "en" - English
	// "ru" - Русский (by default)
	// "pl" - Polski
	// "de" - Deutsch
	// "fr" - Français
	// "es" - Español
	// "zh-cn" - 简体中文
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string
	// Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of None is applied (by default).
	Limit *int
	// Search type. Default is "startswith". Valid values:
	//
	// "startswith" - Search by initial characters of player name. Minimum length: 3 characters. Maximum length: 24 characters. (by default)
	// "exact" - Search by exact match of player name. Case insensitive. You can enter several names, separated with commas (up to 100).
	Type_ *string
}

type AccountList struct {
	// Player ID
	AccountId *int `json:"account_id,omitempty"`
	// Player name
	Nickname *string `json:"nickname,omitempty"`
}
