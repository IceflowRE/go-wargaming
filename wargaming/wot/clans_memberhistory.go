package wot

import (
	"github.com/IceflowRE/go-wargaming/wargaming/wgnTime"
)

// ClansMemberhistoryOptions options.
type ClansMemberhistoryOptions struct {
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
}

type ClansMemberhistory struct {
	// Player account ID
	AccountId *int `json:"account_id,omitempty"`
	// Clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Date when player joined clan
	JoinedAt *wgnTime.UnixTime `json:"joined_at,omitempty"`
	// Date when player left clan
	LeftAt *wgnTime.UnixTime `json:"left_at,omitempty"`
	// Last position in clan
	Role *string `json:"role,omitempty"`
}
