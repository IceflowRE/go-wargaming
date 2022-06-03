package wows

// AccountListOptions options.
type AccountListOptions struct {
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
	// Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
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
