package wgn

// ServersInfoOptions options.
type ServersInfoOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Game ID. Maximum limit: 100. Valid values:
	//
	// "wotb" - World of Tanks Blitz
	// "wot" - World of Tanks
	// "wows" - World of Warships
	Game []string
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
}

type ServersInfo struct {
	// Number of online players
	PlayersOnline *int `json:"players_online,omitempty"`
	// Server ID
	Server *string `json:"server,omitempty"`
}
