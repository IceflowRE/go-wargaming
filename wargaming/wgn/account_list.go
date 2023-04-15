// Auto generated file!

package wgn

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type AccountListOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Name of the game to player search. If the parameter is not specified, search will be executed across known games. Maximum limit: 10. Valid values:
	//
	// "wotb" - World of Tanks Blitz
	// "wot" - World of Tanks
	// "wows" - World of Warships
	// "wowp" - World of Warplanes
	Game []string `json:"game,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "ru" - Русский
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
	Language *string `json:"language,omitempty"`
	// Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of None is applied (by default).
	Limit *int `json:"limit,omitempty"`
	// Search type. Default is "startswith". Valid values:
	//
	// "startswith" - Search by initial characters of player name. Minimum length: 3 characters. Maximum length: 24 characters. (by default)
	// "exact" - Search by exact match of player name. Case insensitive. You can enter several names, separated with commas (up to 100).
	Type *string `json:"type,omitempty"`
}

type AccountList struct {
	// Player ID
	AccountId *int `json:"account_id,omitempty"`
	// Date when player's account was created
	CreatedAt *wgnTime.UnixTime `json:"created_at,omitempty"`
	// List of games player has played
	Games []string `json:"games,omitempty"`
	// Player name
	Nickname *string `json:"nickname,omitempty"`
}
