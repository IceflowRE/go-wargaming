// Auto generated file!

package wgn

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type AccountInfoOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
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
}

type AccountInfo struct {
	// Player ID
	AccountId *int `json:"account_id,omitempty"`
	// Date when player's account was created
	CreatedAt *wgnTime.UnixTime `json:"created_at,omitempty"`
	// List of games player has played
	Games []string `json:"games,omitempty"`
	// Player name
	Nickname *string `json:"nickname,omitempty"`
	// Player's private data
	Private *struct {
		// Amount of Free Experience
		FreeXp *int `json:"free_xp,omitempty"`
		// Current gold balance
		Gold *int `json:"gold,omitempty"`
		// Premium Account expiration date
		PremiumExpiresAt *wgnTime.UnixTime `json:"premium_expires_at,omitempty"`
	} `json:"private,omitempty"`
}
